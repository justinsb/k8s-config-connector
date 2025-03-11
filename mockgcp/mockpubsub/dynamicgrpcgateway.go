package mockpubsub

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/klog/v2"
)

type grpcMux struct {
	conn     *grpc.ClientConn
	services []*grpcService
}

func NewGRPCMux(conn *grpc.ClientConn) (*grpcMux, error) {
	return &grpcMux{conn: conn}, nil
}

// func AddService[T any](mux *grpcMux) {

// 	var protoMessage proto.Message
// 	protoMessageType := reflect.TypeOf(&protoMessage).Elem()

// 	var t T
// 	goType := reflect.TypeOf(&t).Elem()

// 	klog.Infof("kind of t is %v", goType.Kind())

// 	var protoTypes []reflect.Type
// 	n := goType.NumMethod()
// 	for i := 0; i < n; i++ {
// 		method := goType.Method(i)
// 		klog.Infof("method: %v", method)

// 		for j := 0; j < method.Type.NumOut(); j++ {
// 			out := method.Type.Out(j)
// 			if out.AssignableTo(protoMessageType) {
// 				protoTypes = append(protoTypes, out)
// 			}
// 		}
// 	}

// 	if len(protoTypes) == 0 {
// 		klog.Fatalf("found no protobuf types in %T", t)
// 	}

// 	// Pick one protobuf and use it to find the FileDescriptor, from there we can find the services
// 	msg := reflect.New(protoTypes[0]).Elem().Interface()
// 	klog.Infof("msg is %T", msg)
// 	md := msg.(proto.Message).ProtoReflect().Descriptor()
// 	fd := md.ParentFile()

// 	// Find the matching ServiceDescriptor in the FileDescriptor
// 	var matchingServices []protoreflect.ServiceDescriptor
// 	services := fd.Services()
// 	for i := range services.Len() {
// 		service := services.Get(i)
// 		// TODO: Is there a better way to match this?
// 		if string(service.Name())+"Server" == goType.Name() {
// 			matchingServices = append(matchingServices, service)
// 		}
// 	}

// 	if len(matchingServices) == 0 {
// 		klog.Fatalf("cannot match service for %v", goType.Name())
// 	}
// 	if len(matchingServices) > 1 {
// 		klog.Fatalf("found multiple matching service for %v", goType.Name())
// 	}
// 	if err := mux.AddService(goType, matchingServices[0]); err != nil {
// 		klog.Fatalf("failed to add service: %v", err)
// 	}
// }

// func (m *grpcMux) AddService(goType reflect.Type, service protoreflect.ServiceDescriptor) error {
// 	s, err := newGRPCService(goType, service)
// 	if err != nil {
// 		return err
// 	}

// 	m.services = append(m.services, s)

//		return nil
//	}

func (m *grpcMux) AddService(client any) {

	var protoMessage proto.Message
	protoMessageType := reflect.TypeOf(&protoMessage).Elem()

	goType := reflect.TypeOf(client)

	klog.Infof("kind of t is %v", goType.Kind())

	goTypeMethodNames := make(map[string]bool)

	var discoveredProtobufTypes []reflect.Type
	n := goType.NumMethod()
	for i := 0; i < n; i++ {
		method := goType.Method(i)
		klog.Infof("method: %v", method)
		goTypeMethodNames[method.Name] = true

		for j := 0; j < method.Type.NumOut(); j++ {
			out := method.Type.Out(j)
			if out.AssignableTo(protoMessageType) {
				discoveredProtobufTypes = append(discoveredProtobufTypes, out)
			}
		}
	}

	if len(discoveredProtobufTypes) == 0 {
		klog.Fatalf("found no protobuf types in %T", client)
	}

	// Use the protobuf types to find the FileDescriptor, from there we can find the services
	var matchingServices []protoreflect.ServiceDescriptor
	for _, protoType := range discoveredProtobufTypes {
		msg := reflect.New(protoType).Elem().Interface()
		// klog.Infof("msg is %T", msg)
		md := msg.(proto.Message).ProtoReflect().Descriptor()
		fd := md.ParentFile()

		services := fd.Services()
		for i := range services.Len() {
			service := services.Get(i)

			isMatch := true
			methods := service.Methods()
			for j := range methods.Len() {
				method := methods.Get(j)
				if !goTypeMethodNames[string(method.Name())] {
					isMatch = false
					break
				}
			}

			// TODO: Is there a better way to match this?
			if isMatch {
				matchingServices = append(matchingServices, service)
			}
		}
		if len(matchingServices) > 0 {
			break
		}
	}

	if len(matchingServices) == 0 {
		klog.Fatalf("cannot match service for %v", goType.Name())
	}
	if len(matchingServices) > 1 {
		klog.Fatalf("found multiple matching service for %v", goType.Name())
	}

	s, err := newGRPCService(client, matchingServices[0])
	if err != nil {
		klog.Fatalf("adding grpc service: %v", err)
	}

	m.services = append(m.services, s)
}

type grpcService struct {
	grpcClient any
	service    protoreflect.ServiceDescriptor

	httpDefaultHost string

	methods []*grpcMethod
}

func newGRPCService(grpcClient any, service protoreflect.ServiceDescriptor) (*grpcService, error) {
	obj := &grpcService{
		grpcClient: grpcClient,
		service:    service,
	}

	var errs []error
	service.Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		klog.Infof("option %v %v", fd, v)
		switch fd.Kind() {
		case protoreflect.MessageKind:
			switch fd.Message().FullName() {
			default:
				errs = append(errs, fmt.Errorf("unhandled annotation message %q", fd.Message().FullName()))
			}

		case protoreflect.StringKind:
			switch fd.JSONName() {
			case "[google.api.default_host]":
				obj.httpDefaultHost = v.String()
			case "[google.api.oauth_scopes]":
				// ignore for now
				// obj.oauthScopes = v.String()
			default:
				errs = append(errs, fmt.Errorf("unhandled annotation string %q", fd.JSONName()))
			}

		default:
			errs = append(errs, fmt.Errorf("unhandled option kind in %v", fd))
		}

		return true
	})

	if len(errs) != 0 {
		return nil, errors.Join(errs...)
	}

	goType := reflect.TypeOf(grpcClient)

	serviceMethods := service.Methods()
	for j := 0; j < serviceMethods.Len(); j++ {
		serviceMethod := serviceMethods.Get(j)
		goMethodType, ok := goType.MethodByName(string(serviceMethod.Name()))
		if !ok {
			return nil, fmt.Errorf("unable to find go method for %v", serviceMethod)
		}
		clientMethod := reflect.ValueOf(grpcClient).MethodByName(string(serviceMethod.Name()))
		if clientMethod.IsZero() {
			return nil, fmt.Errorf("unable to find client method for %v", serviceMethod)
		}

		if err := obj.addGRPCMethod(clientMethod, goMethodType, serviceMethod); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

func (s *grpcService) addGRPCMethod(goMethod reflect.Value, goMethodType reflect.Method, method protoreflect.MethodDescriptor) error {
	var httpRule *annotations.HttpRule
	var errs []error
	// klog.Infof("method: %v", method)
	method.Options().ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		// klog.Infof("option %v %v", fd, v)
		switch fd.Kind() {
		case protoreflect.MessageKind:
			switch fd.Message().FullName() {
			case "google.api.HttpRule":
				httpRule = proto.GetExtension(method.Options(), annotations.E_Http).(*annotations.HttpRule)
				// klog.Infof("httpRule: %+v", httpRule)
			default:
				errs = append(errs, fmt.Errorf("unhandled annotation message %q", fd.Message().FullName()))
			}

		case protoreflect.StringKind:
			switch fd.JSONName() {
			case "[google.api.method_signature]":
				// ignore for now
			default:
				errs = append(errs, fmt.Errorf("unhandled annotation string %q", fd.JSONName()))
			}

		default:
			errs = append(errs, fmt.Errorf("unhandled option kind in %v", fd))
		}

		return true
	})

	if len(errs) != 0 {
		return errors.Join(errs...)
	}

	if httpRule == nil {
		klog.Warningf("grpc method did not have http rule: %+v", method)
		return nil
	}

	addMethod := func(httpRule *annotations.HttpRule, httpMethod string, httpPath string) {
		m := &grpcMethod{
			method:       method,
			goMethod:     goMethod,
			goMethodType: goMethodType,
			httpMethod:   httpMethod,
			httpPath:     httpPath,
			httpRule:     httpRule,
		}
		s.methods = append(s.methods, m)

	}

	processRule := func(rule *annotations.HttpRule) {
		if rule.GetGet() != "" {
			addMethod(rule, http.MethodGet, rule.GetGet())
		}
		if rule.GetDelete() != "" {
			addMethod(rule, http.MethodDelete, rule.GetDelete())
		}
		if rule.GetPut() != "" {
			addMethod(rule, http.MethodPut, rule.GetPut())
		}
		if rule.GetPost() != "" {
			addMethod(rule, http.MethodPost, rule.GetPost())
		}
		if rule.GetPatch() != "" {
			addMethod(rule, http.MethodPatch, rule.GetPatch())
		}
		if custom := rule.GetCustom(); custom != nil {
			addMethod(rule, custom.GetKind(), custom.GetPath())
		}
	}
	processRule(httpRule)
	for _, additionalBinding := range httpRule.GetAdditionalBindings() {
		processRule(additionalBinding)
	}

	for _, method := range s.methods {
		var matchers []matcher
		httpPath := method.httpPath
		httpPath = strings.TrimPrefix(httpPath, "/")
		for httpPath != "" {
			if httpPath[0] == '{' {
				closeQuote := strings.Index(httpPath, "}")
				if closeQuote == -1 {
					return fmt.Errorf("invalid httpPath %q", method.httpPath)
				}
				token := httpPath[1:closeQuote]
				matchers = append(matchers, newMatchWildcard(token))
				httpPath = httpPath[closeQuote+1:]
			} else {
				nextSlash := strings.Index(httpPath, "/")
				if nextSlash == -1 {
					nextSlash = len(httpPath)
				}
				token := httpPath[:nextSlash]
				matchers = append(matchers, newMatchLiteral(token))
				httpPath = httpPath[nextSlash:]
			}
			httpPath = strings.TrimPrefix(httpPath, "/")
		}
		method.matchers = matchers
	}

	return nil
}

func (m *grpcMethod) Name() string {
	return string(m.method.Name())
}

type grpcMethod struct {
	method       protoreflect.MethodDescriptor
	goMethod     reflect.Value
	goMethodType reflect.Method

	httpRule *annotations.HttpRule

	httpMethod string
	httpPath   string

	matchers []matcher
}

type matcher interface {
	Match(tokens []string, matches map[string]string) (bool, []string)
}

type matchLiteral struct {
	literal string
}

func newMatchLiteral(literal string) *matchLiteral {
	return &matchLiteral{literal: literal}
}

func (m *matchLiteral) Match(tokens []string, matches map[string]string) (bool, []string) {
	if len(tokens) == 0 {
		return false, nil
	}
	if tokens[0] == m.literal {
		return true, tokens[1:]
	}
	return false, nil
}

type matchWildcard struct {
	key           string
	patternTokens []string
}

func newMatchWildcard(wildcard string) *matchWildcard {
	eqTokens := strings.Split(wildcard, "=")
	if len(eqTokens) == 2 {
		// e.g. {name=projects/*/locations/*/foo/*}
		key := eqTokens[0]
		pattern := eqTokens[1]

		patternTokens := strings.Split(pattern, "/")

		return &matchWildcard{
			key:           key,
			patternTokens: patternTokens,
		}
	} else {
		klog.Fatalf("unhandled wildcard token: %q", wildcard)
		return nil
	}
}

func (m *matchWildcard) Match(tokens []string, matches map[string]string) (bool, []string) {
	if len(tokens) < len(m.patternTokens) {
		return false, nil
	}
	for i, patternToken := range m.patternTokens {
		if patternToken == "*" {
			continue
		}
		if tokens[i] != patternToken {
			return false, nil
		}
	}

	matches[m.key] = strings.Join(tokens[:len(m.patternTokens)], "/")
	return true, tokens[len(m.patternTokens):]
}

func (m *grpcMethod) Match(tokens []string) (map[string]string, bool) {
	values := make(map[string]string)
	for _, matcher := range m.matchers {
		ok, newTokens := matcher.Match(tokens, values)
		if !ok {
			return nil, false
		}
		tokens = newTokens
	}
	if len(tokens) != 0 {
		return nil, false
	}
	return values, true
}

func (m *grpcMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	klog.Infof("host: %v", host)
	url := r.URL.Path
	klog.Infof("url: %v", url)
	tokens := strings.Split(strings.TrimPrefix(url, "/"), "/")

	for _, service := range m.services {
		klog.Infof("tokens is %v", tokens)
		for _, method := range service.methods {
			if method.httpMethod != r.Method {
				continue
			}

			matches, ok := method.Match(tokens)
			if !ok {
				for i, matcher := range method.matchers {
					klog.Infof("matcher is %d %T %v", i, matcher, matcher)
				}
				klog.Infof("did not match method %v %+v: %+v %v", method.Name(), method.httpPath, method.matchers, tokens)
				continue
			}

			klog.Infof("matched method %v %+v: %v", method.Name(), method.httpPath, matches)
			m.serveHTTPMethod(w, r, method, matches)
			return
		}
	}

	klog.Warningf("http request not matched; %v %v", r.Method, r.URL)
	http.Error(w, "not found", http.StatusNotFound)
}

func (m *grpcMux) serveHTTPMethod(w http.ResponseWriter, r *http.Request, method *grpcMethod, values map[string]string) {
	call := &httpMethodCall{
		r: r,
		w: w,
	}

	host := r.Host
	klog.Infof("host: %v", host)
	url := r.URL.Path
	klog.Infof("url: %v", url)

	var body []byte
	if r.Body != nil {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			klog.Errorf("failed to read body: %v", err)
			http.Error(w, "invalid body", http.StatusBadRequest)
			return
		}
		body = b
		klog.Infof("body: %v", string(body))
	} else {
		klog.Infof("no request body")
	}

	klog.Infof("values: %v", values)
	inputMessage := method.method.Input()
	klog.Infof("inputMessage is %v", inputMessage)

	ctx := r.Context()

	var inArgs []reflect.Value
	for i := range method.goMethodType.Type.NumIn() {
		if i == 0 {
			// Skip receiver
			continue
		}
		inType := method.goMethodType.Type.In(i)
		klog.Infof("inType is %v", inType)

		prefix := ""
		if inType.Kind() == reflect.Ptr {
			inType = inType.Elem()
			prefix += "*"
		}
		if inType.Kind() == reflect.Slice {
			inType = inType.Elem()
			prefix += "[]"
		}

		inTypeName := prefix + inType.PkgPath() + "." + inType.Name()
		klog.Infof("inTypeName %q", inTypeName)
		if inTypeName == "context.Context" {
			inArgs = append(inArgs, reflect.ValueOf(ctx))
		} else if inTypeName == "[]google.golang.org/grpc.CallOption" {
			var callOptions []grpc.CallOption
			if method.goMethodType.Type.IsVariadic() {
				for _, callOption := range callOptions {
					inArgs = append(inArgs, reflect.ValueOf(callOption))
				}
			} else {
				inArgs = append(inArgs, reflect.ValueOf(callOptions))
			}
		} else {
			inArg := reflect.New(inType)
			klog.Infof("inArg is %+v", inArg)

			protoMessage := inArg.Interface().(proto.Message)

			if len(body) != 0 {
				if err := protojson.Unmarshal(body, protoMessage); err != nil {
					klog.Errorf("failed to unmarshal body: %v", err)
					http.Error(w, "invalid body", http.StatusBadRequest)
					return
				}
			}

			for k, v := range values {
				fd := protoMessage.ProtoReflect().Descriptor().Fields().ByJSONName(k)
				if fd == nil {
					klog.Fatalf("value field %q not found", k)
				}
				protoVal := protoreflect.ValueOf(v)
				protoMessage.ProtoReflect().Set(fd, protoVal)
			}
			klog.Infof("inArg is %+v", prototext.Format(protoMessage))
			inArgs = append(inArgs, inArg)
		}
	}

	klog.Infof("calling method by reflection")
	out := method.goMethod.Call(inArgs)
	klog.Infof("out is %v", out)
	// outputMessage := method.method.Output()

	if len(out) != 2 {
		klog.Fatalf("output format not handled, expected two output parameters")
	}

	if !out[1].IsNil() {
		err, ok := out[1].Interface().(error)
		if !ok {
			klog.Fatalf("expected second parameter to be error, was %T", out[1])
		}
		call.SendErrorResponse(err)
		return
	}

	response, ok := out[0].Interface().(proto.Message)
	if !ok {
		klog.Fatalf("expected first parameter to be proto.Message, was %T", out[0])
	}
	call.SendResponse(response)
}

type httpMethodCall struct {
	r *http.Request
	w http.ResponseWriter
}

func (c *httpMethodCall) SendErrorResponse(err error) {
	klog.Warningf("sending error response for %T %+v", err, err)

	statusErr, ok := status.FromError(err)
	if ok {
		httpCode := http.StatusInternalServerError
		switch statusErr.Code() {
		case codes.NotFound:
			httpCode = http.StatusNotFound
		}

		body, err := protojson.Marshal(statusErr.Proto())
		if err != nil {
			klog.Errorf("failed to marshal error: %v", err)
			http.Error(c.w, "internal error", http.StatusInternalServerError)
			return
		}

		c.w.Header().Set("Content-Type", "application/json")
		c.w.WriteHeader(httpCode)
		if _, err := c.w.Write(body); err != nil {
			klog.Errorf("failed to write error: %v", err)
		}
		klog.Infof("sent response %v with body %v", httpCode, string(body))
		return
	}
	klog.Warningf("stub-handling error %v", err)
	http.Error(c.w, err.Error(), http.StatusInternalServerError)
}

func (c *httpMethodCall) SendResponse(response proto.Message) {
	httpCode := http.StatusOK

	body, err := protojson.Marshal(response)
	if err != nil {
		klog.Errorf("failed to marshal response: %v", err)
		http.Error(c.w, "internal error", http.StatusInternalServerError)
		return
	}

	c.w.Header().Set("Content-Type", "application/json")
	c.w.WriteHeader(httpCode)
	if _, err := c.w.Write(body); err != nil {
		klog.Errorf("failed to write error: %v", err)
	}
	klog.Infof("sent response %v with body %v", httpCode, string(body))
}

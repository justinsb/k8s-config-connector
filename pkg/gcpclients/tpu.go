package gcpclients

import (
	"context"
	"math"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	tpuv2 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/tpu/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
)

type TPUV2Client struct {
	tpuv2.TpuClient
	lroClient *lroauto.OperationsClient
}

func NewTPUV2Client(ctx context.Context, opts ...option.ClientOption) (*TPUV2Client, error) {
	clientOpts := defaultGRPCClientOptions()
	// if newClientHook != nil {
	// 	hookOpts, err := newClientHook(ctx, clientHookParams{})
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	clientOpts = append(clientOpts, hookOpts...)
	// }

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	// connPool, err := gtransport.Dial(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	// client := Client{CallOptions: defaultCallOptions()}

	// c := &gRPCClient{
	// 	connPool:         connPool,
	// 	client:           apikeyspb.NewApiKeysClient(connPool),
	// 	CallOptions:      &client.CallOptions,
	// 	operationsClient: longrunningpb.NewOperationsClient(connPool),
	// }
	// c.setGoogleClientInfo()

	// client.internalClient = c

	lroClient, err := lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}

	return &TPUV2Client{
		TpuClient: tpuv2.NewTpuClient(connPool),
		lroClient: lroClient,
	}, nil
}

func (c *TPUV2Client) WaitForLRO(ctx context.Context, op *longrunningpb.Operation) error {
	lro := longrunning.InternalNewOperation(c.lroClient, op)
	if err := lro.Wait(ctx, nil); err != nil {
		return err
	}
	return nil
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("tpu.googleapis.com:443"),
		internaloption.WithDefaultEndpointTemplate("tpu.UNIVERSE_DOMAIN:443"),
		internaloption.WithDefaultMTLSEndpoint("tpu.mtls.googleapis.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://tpu.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

// DefaultAuthScopes reports the default set of authentication scopes to use with this package.
func DefaultAuthScopes() []string {
	return []string{
		"https://www.googleapis.com/auth/cloud-platform",
	}
}

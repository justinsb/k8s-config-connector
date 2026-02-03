package changecookies

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
)

func New(desired, actual proto.Message) (*Cookie, error) {
	desiredHash, err := hashProto(desired)
	if err != nil {
		return nil, fmt.Errorf("calculating desired-state hash: %w", err)
	}
	actualHash, err := hashProto(actual)
	if err != nil {
		return nil, fmt.Errorf("calculating actual-state hash: %w", err)
	}

	if objectHash != nil {
		objectHash[desiredHash] = desired
		objectHash[actualHash] = actual
	}
	return &Cookie{DesiredStateHash: desiredHash, ActualStateHash: actualHash}, nil
}

var objectHash = make(map[string]proto.Message)

func FullObject(hash string) proto.Message {
	return objectHash[hash]
}

// Cookie is used for stateful reconciliation.
// It is stored in the status of the KCC resource.
type Cookie struct {
	DesiredStateHash string
	ActualStateHash  string
}

// ParseCookie parses a cookie string into a Cookie struct.
func ParseCookie(cookieStr string) (*Cookie, error) {
	parts := strings.Split(cookieStr, "/")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid cookie format: %q", cookieStr)
	}
	if parts[0] != "c" {
		// C is for cookie, and I suggest V2 should be C2
		return nil, fmt.Errorf("invalid cookie version: %q", cookieStr)
	}
	return &Cookie{DesiredStateHash: parts[0], ActualStateHash: parts[1]}, nil
}

// ComposeCookie creates a cookie string from the desired and actual hashes, suitable for use in an annotation.
func (c *Cookie) String() string {
	return "c/" + c.DesiredStateHash + "/" + c.ActualStateHash
}

// func (c *Cookie) Equal(lastModifiedCookie *string) bool {
// 	if lastModifiedCookie == nil {
// 		return false
// 	}
// 	other := *lastModifiedCookie
// 	if other == "" {
// 		return false
// 	}
// 	return c.String() == other
// }

// cookieEncoding is the encoding we use for the hashes in the cookie.
var cookieEncoding = base64.RawURLEncoding

// hashProto calculates a hash of a proto message.
// We use this to detect changes to the GCP resource.
func hashProto(obj proto.Message) (string, error) {
	// We use a deterministic proto marshaler.
	j, err := (proto.MarshalOptions{Deterministic: true}).Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("cannot marshal proto: %w", err)
	}

	// We use sha1 for hashing. We don't need a cryptographically secure hash, we just need to detect changes. SHA1 is widely used for this purpose (e.g. in git) and has good performance.
	hasher := sha1.New()
	if _, err := hasher.Write(j); err != nil {
		return "", fmt.Errorf("cannot write to hasher: %w", err)
	}
	return cookieEncoding.EncodeToString(hasher.Sum(nil)), nil
}

// // NormalizeProto clears fields that are not significant for comparison.
// // It modifies the passed-in proto.
// func normalizeProto(pb proto.Message) {
// 	// TODO: Should we also clear fields like `uid` and `name`?
// 	// The problem is that they are not consistently marked as output-only.

// 	// We clear fields that are known to be volatile
// 	volatileFieldNames := []string{"etag", "update_time", "updated_time"}

// 	// We also clear fields that are marked as output-only
// 	// We do this by walking the fields and checking for `field_behavior: OUTPUT_ONLY`
// 	// We build a field mask and then use that to clear the fields.
// 	var paths []string
// 	pb.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
// 		options := fd.Options()
// 		if options != nil {
// 			if proto.HasExtension(options, annotations.E_FieldBehavior) {
// 				fieldBehavior := proto.GetExtension(options, annotations.E_FieldBehavior).([]annotations.FieldBehavior)
// 				for _, b := range fieldBehavior {
// 					if b == annotations.FieldBehavior_OUTPUT_ONLY {
// 						paths = append(paths, string(fd.Name()))
// 					}
// 				}
// 			}
// 		}
// 		for _, fieldName := range volatileFieldNames {
// 			if string(fd.Name()) == fieldName {
// 				paths = append(paths, string(fd.Name()))
// 			}
// 		}
// 		return true
// 	})

// 	if len(paths) > 0 {
// 		fm, err := fieldmaskpb.New(pb, paths...)
// 		if err != nil {
// 			// This should not happen
// 			panic(fmt.Sprintf("error creating fieldmask: %v", err))
// 		}
// 		fm.Normalize()
// 		if !fm.IsValid(pb) {
// 			// This should not happen
// 			panic(fmt.Sprintf("invalid fieldmask %v for %T", fm, pb))
// 		}
// 		clearFields(pb.ProtoReflect(), fm.GetPaths())
// 	}
// }

// // clearFields clears the given fields from the message.
// func clearFields(m protoreflect.Message, paths []string) {
// 	for _, path := range paths {
// 		// Note: We don't support nested fields yet
// 		fd := m.Descriptor().Fields().ByName(protoreflect.Name(path))
// 		if fd == nil {
// 			// This should not happen
// 			panic(fmt.Sprintf("field %q not found in %v", path, m.Descriptor().FullName()))
// 		}
// 		m.Clear(fd)
// 	}
// }

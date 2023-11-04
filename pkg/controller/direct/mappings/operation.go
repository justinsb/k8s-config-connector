package mappings

import (
	"context"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Operation interface {
	Get(dest any) error
}

type mapOp struct {
	Operation  Operation
	mapping    *Mapping
	rootIn     *ReflectValue
	rootOut    *ReflectValue
	krmToCloud bool
}

// TODO: Move this out
type KCCOperation struct {
	Ctx       context.Context
	Client    client.Client
	Namespace string
}

func (o *KCCOperation) Get(dest any) error {
	switch dest := dest.(type) {
	case **KCCOperation:
		*dest = o
		return nil
	default:
		return fmt.Errorf("unexpected dest type %T", dest)
	}
}

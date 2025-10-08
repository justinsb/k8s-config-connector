// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/identity"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// DatabaseIDURL is the format for the externalRef of a FirestoreDatabase.
	DatabaseIDURL = parent.ProjectURLTemplate + "/databases/{{database}}"
	ServiceDomain = "firestore.googleapis.com"
)

var _ identity.Identity = &FirestoreDatabaseIdentity{}

// FirestoreDatabaseIdentity represents the identity of a Firestore Database.
// +k8s:deepcopy-gen=false
type FirestoreDatabaseIdentity struct {
	Parent   *parent.ProjectParent
	Database string
}

func (i *FirestoreDatabaseIdentity) String() string {
	return i.Parent.String() + "/databases/" + i.Database
}

var FirestoreDatabaseIdentityURL = URLTemplate{
	Parent:  &parent.ProjectParent{},
	Service: ServiceDomain,
	Key:     "databases",
}

func (i *FirestoreDatabaseIdentity) FromExternal(ref string) error {
	parentID, value, err := FirestoreDatabaseIdentityURL.Parse(ref)
	if err != nil {
		return err
	}
	i.Parent = parentID.(*parent.ProjectParent)
	i.Database = value
	return nil
}

type URLTemplate struct {
	Parent  identity.Identity
	Service string
	Key     string
}

func (u *URLTemplate) Template() string {
	s := ""
	if u.Parent != nil {
		s += u.Parent.Template() + "/"
	}
}

func (u *URLTemplate) Parse(ref string) (identity.Identity, string, error) {
	if strings.HasPrefix(ref, "//") {
		if u.Service == "" {
			return nil, "", fmt.Errorf("unexpected service-qualified reference %q", ref)
		}
		if !strings.HasPrefix(ref, "//"+u.Service+"/") {
			return nil, "", fmt.Errorf("expected service-qualified reference to start with //%s/: %q", u.Service, ref)
		}
		ref = strings.TrimPrefix(ref, "//"+u.Service+"/")
	}

	tokens := strings.Split(ref, "/")
	if len(tokens) < 2 {
		return nil, "", fmt.Errorf("expected at least two tokens in reference %q", ref)
	}
	n := len(tokens)
	if tokens[n-2] != u.Key {
		return nil, "", fmt.Errorf("expected %q in reference %q", u.Key, ref)
	}
	value := tokens[n-1]

	if u.Parent == nil {
		if len(tokens) != 2 {
			return nil, "", fmt.Errorf("found extra tokens in reference %q, expected %q", ref, u.Key+"/{value}")
		}
		return nil, value, nil
	}

	parent := reflect.New(reflect.TypeOf(u.Parent).Elem()).Interface().(identity.Identity)
	if err := parent.FromExternal(strings.Join(tokens[:n-2], "/")); err != nil {
		return nil, "", fmt.Errorf("parsing ref %q: %w", ref, err)
	}
	return parent, value, nil
}

var _ identity.Resource = &FirestoreDatabase{}

func (obj *FirestoreDatabase) GetIdentity(ctx context.Context, reader client.Reader) (identity.Identity, error) {
	newIdentity := &FirestoreDatabaseIdentity{}

	// Resolve Parent
	if err := obj.Spec.ProjectRef.Normalize(ctx, reader, obj.GetNamespace()); err != nil {
		return nil, fmt.Errorf("resolving spec.parentRef: %w", err)
	}
	newIdentity.Parent = &parent.ProjectParent{}
	if err := newIdentity.Parent.FromExternal(obj.Spec.ProjectRef.External); err != nil {
		return nil, fmt.Errorf("parsing projectRef.external=%q: %w", obj.Spec.ProjectRef.External, err)
	}
	// Get desired ID
	newIdentity.Database = common.ValueOf(obj.Spec.ResourceID)
	if newIdentity.Database == "" {
		newIdentity.Database = obj.GetName()
	}
	if newIdentity.Database == "" {
		return nil, fmt.Errorf("cannot resolve resource ID")
	}
	// Validate against the ID stored in status.externalRef
	externalRef := common.ValueOf(obj.Status.ExternalRef)
	if externalRef != "" {
		statusIdentity := &FirestoreDatabaseIdentity{}
		if err := statusIdentity.FromExternal(externalRef); err != nil {
			return nil, fmt.Errorf("cannot parse existing externalRef=%q: %w", externalRef, err)
		}
		if statusIdentity.String() != newIdentity.String() {
			return nil, fmt.Errorf("existing externalRef=%q does not match the identity resolved from spec: %q", externalRef, newIdentity.String())
		}
	}
	return newIdentity, nil
}

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

package structuredreporting

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// DebugLogListener is a Listener that logs structured reporting events
type DebugLogListener struct {
}

var _ Listener = &DebugLogListener{}

// OnError is called when a controller calls ReportError
func (l *DebugLogListener) OnError(ctx context.Context, err error, args ...any) {
	log := log.FromContext(ctx)
	log.Info("structuredreporting OnError",
		"error", err)
}

// OnDiff is called when a controller calls ReportDiffs
func (l *DebugLogListener) OnDiff(ctx context.Context, diffs *Diff) {
	log := log.FromContext(ctx)

	jsonFormat := func(v any) string {
		if protoMsg, ok := v.(protoreflect.Message); ok {
			j, err := protojson.Marshal(protoMsg.Interface())
			if err == nil {
				return string(j)
			}
		}
		if protoMsg, ok := v.(proto.Message); ok {
			j, err := protojson.Marshal(protoMsg)
			if err == nil {
				return string(j)
			}
		}
		j, err := json.Marshal(v)
		if err != nil {
			// Fallback
			return fmt.Sprintf("%+v", v)
		}
		return string(j)
	}
	var diffFields []string
	for _, fieldID := range diffs.Fields {
		klog.Infof("OnDiff field %q changed from %T to %T", fieldID.ID, fieldID.Old, fieldID.New)
		klog.Infof("OnDiff field %q changed from %v to %v", fieldID.ID, jsonFormat(fieldID.Old), jsonFormat(fieldID.New))
		diffFields = append(diffFields, fmt.Sprintf("%v[%v => %v]", fieldID.ID, jsonFormat(fieldID.Old), jsonFormat(fieldID.New)))
	}
	log.Info("structuredreporting OnDiff",
		"diff.fields", strings.Join(diffFields, ", "),
		"diff.isNewObject", diffs.IsNewObject,
	)
}

// OnReconcileStart is called when a controller calls ReportReconcileStart
func (l *DebugLogListener) OnReconcileStart(ctx context.Context, u *unstructured.Unstructured, t k8s.ReconcilerType) {
	log := log.FromContext(ctx)
	log.Info("structuredreporting OnReconcileStart",
		"object.kind", u.GroupVersionKind().Kind,
		"object.name", u.GetName())
}

// OnReconcileEnd is called when a controller calls ReportReconcileEnd
func (l *DebugLogListener) OnReconcileEnd(ctx context.Context, u *unstructured.Unstructured, result reconcile.Result, err error, t k8s.ReconcilerType) {
	log := log.FromContext(ctx)
	log.Info("structuredreporting OnReconcileEnd",
		"object.kind", u.GroupVersionKind().Kind,
		"object.name", u.GetName(),
		"result", result,
		"error", err)
}

// LogFieldUpdates is a Listener that logs updated fields during reconciliation,
// only when objects are being updated.
type LogFieldUpdates struct {
}

var _ Listener = &LogFieldUpdates{}

// OnError is called when a controller calls ReportError
func (l *LogFieldUpdates) OnError(ctx context.Context, err error, args ...any) {
}

// OnDiff is called when a controller calls ReportDiffs
func (l *LogFieldUpdates) OnDiff(ctx context.Context, diffs *Diff) {
	log := log.FromContext(ctx)
	if !diffs.IsNewObject {
		log.Info("detected changes to fields; triggering update",
			"changedFields", diffs.FieldIDs(),
		)
	}
}

// OnReconcileStart is called when a controller calls ReportReconcileStart
func (l *LogFieldUpdates) OnReconcileStart(ctx context.Context, u *unstructured.Unstructured, t k8s.ReconcilerType) {
}

// OnReconcileEnd is called when a controller calls ReportReconcileEnd
func (l *LogFieldUpdates) OnReconcileEnd(ctx context.Context, u *unstructured.Unstructured, result reconcile.Result, err error, t k8s.ReconcilerType) {
}

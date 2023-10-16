// Copyright 2022 Google LLC
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

package resourceoverrides

import (
	"context"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/resourceoverrides/operations"
	"k8s.io/klog/v2"
)

func GetComputeForwardingRuleResourceOverrides() ResourceOverrides {
	ro := ResourceOverrides{
		Kind: "ComputeForwardingRule",
	}
	ro.Overrides = append(ro.Overrides, noLabelsOnCreate())
	return ro
}

func noLabelsOnCreate() ResourceOverride {
	o := ResourceOverride{}

	o.PreTerraformApply = func(ctx context.Context, op *operations.PreTerraformApply) error {
		if op.LiveState.Empty() {
			klog.Infof("PreTerraformApply: config=%+v", op.TerraformConfig.Config)
			delete(op.TerraformConfig.Config, "labels")
			klog.Infof("PreTerraformApply: after labels=%+v", op.TerraformConfig.Config)
		}

		return nil
	}

	return o
}

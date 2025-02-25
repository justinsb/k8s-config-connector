// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +tool:fuzz-gen
// proto.message: google.cloud.billing.budgets.v1.Budget
// api.group: billingbudgets.cnrm.cloud.google.com

package billingbudgets

import (
	pb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
	moneypb "google.golang.org/genproto/googleapis/type/money"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(BillingBudgetsBudgetFuzzer())
}

func BillingBudgetsBudgetsBudgetFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Budget{},
		BillingBudgetsBudgetSpec_FromProto, BillingBudgetsBudgetSpec_ToProto,
		BillingBudgetsBudgetObservedState_FromProto, BillingBudgetsBudgetObservedState_ToProto,
	)

	f.AddTransform(func(obj *pb.Budget) {
		// Skip the unspecified usagePeriod (usagePeriod: {})
		if obj.BudgetFilter != nil {
			if usagePeriod := obj.BudgetFilter.UsagePeriod; usagePeriod != nil {
				switch usagePeriod := usagePeriod.(type) {
				case *pb.Filter_CalendarPeriod:
					if usagePeriod.CalendarPeriod == pb.CalendarPeriod_CALENDAR_PERIOD_UNSPECIFIED {
						obj.BudgetFilter.UsagePeriod = nil
					}
				}
			}
		}

		// Make sure money values are valid
		if obj.Amount != nil {
			switch budgetAmount := obj.Amount.BudgetAmount.(type) {
			case *pb.BudgetAmount_SpecifiedAmount:
				if budgetAmount.SpecifiedAmount != nil {
					fixMoney(budgetAmount.SpecifiedAmount)
				}
			}
		}
	})

	f.SpecFields.Insert(".display_name")
	f.SpecFields.Insert(".budget_filter")
	f.SpecFields.Insert(".amount")
	f.SpecFields.Insert(".threshold_rules")
	f.SpecFields.Insert(".notifications_rule")

	// Special (identity) fields
	f.UnimplementedFields.Insert(".name")

	// Highly mutable fields we dont' want to support
	f.UnimplementedFields.Insert(".etag")

	// Not yet implemented, but maybe in future
	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".budget_filter.labels")

	return f
}

func fixMoney(money *moneypb.Money) {
	// Nanos: The value must be between -999,999,999 and +999,999,999 inclusive.
	if money.Nanos > 999999999 {
		money.Nanos = 999999999
	}
	if money.Nanos < -999999999 {
		money.Nanos = -999999999
	}

	units := money.Units
	// If `units` is positive, `nanos` must be positive or zero.
	if units > 0 && money.Nanos < 0 {
		money.Nanos = -money.Nanos
	}
	// If `units` is negative, `nanos` must be negative or zero.
	if units < 0 && money.Nanos > 0 {
		money.Nanos = -money.Nanos
	}
}

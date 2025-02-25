// Copyright 2024 Google LLC
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

package billingbudgets

import (
	"fmt"
	"strconv"
	"strings"

	pb "cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	datepb "google.golang.org/genproto/googleapis/type/date"
	moneypb "google.golang.org/genproto/googleapis/type/money"
)

func CustomPeriod_StartDate_FromProto(mapCtx *direct.MapContext, in *datepb.Date) *string {
	return date_FromProto(mapCtx, in)
}

func CustomPeriod_StartDate_ToProto(mapCtx *direct.MapContext, in *string) *datepb.Date {
	return date_ToProto(mapCtx, in)
}

func CustomPeriod_EndDate_FromProto(mapCtx *direct.MapContext, in *datepb.Date) *string {
	return date_FromProto(mapCtx, in)
}

func CustomPeriod_EndDate_ToProto(mapCtx *direct.MapContext, in *string) *datepb.Date {
	return date_ToProto(mapCtx, in)
}

func BudgetAmount_SpecifiedAmount_FromProto(mapCtx *direct.MapContext, in *moneypb.Money) *string {
	return money_FromProto(mapCtx, in)
}

func BudgetAmount_SpecifiedAmount_ToProto(mapCtx *direct.MapContext, in *string) *moneypb.Money {
	return money_ToProto(mapCtx, in)
}

func Filter_CalendarPeriod_ToProto(mapCtx *direct.MapContext, in *string) *pb.Filter_CalendarPeriod {
	if in == nil {
		return nil
	}
	v := direct.Enum_ToProto[pb.CalendarPeriod](mapCtx, in)
	return &pb.Filter_CalendarPeriod{CalendarPeriod: v}
}

func date_FromProto(mapCtx *direct.MapContext, in *datepb.Date) *string {
	if in == nil {
		return nil
	}
	s := fmt.Sprintf("%d-%d-%d", in.Year, in.Month, in.Day)
	return &s
}

func date_ToProto(mapCtx *direct.MapContext, in *string) *datepb.Date {
	if in == nil {
		return nil
	}
	s := *in
	tokens := strings.Split(s, "-")
	if len(tokens) != 3 {
		mapCtx.Errorf("invalid date %q", s)
		return nil
	}

	year, err := strconv.ParseInt(tokens[0], 10, 32)
	if err != nil {
		mapCtx.Errorf("invalid date %q", s)
		return nil
	}
	month, err := strconv.ParseInt(tokens[1], 10, 32)
	if err != nil {
		mapCtx.Errorf("invalid date %q", s)
		return nil
	}
	day, err := strconv.ParseInt(tokens[2], 10, 32)
	if err != nil {
		mapCtx.Errorf("invalid date %q", s)
		return nil
	}
	out := &datepb.Date{
		Year:  int32(year),
		Month: int32(month),
		Day:   int32(day),
	}
	return out
}

func money_FromProto(mapCtx *direct.MapContext, in *moneypb.Money) *string {
	if in == nil {
		return nil
	}
	s := strconv.FormatInt(in.Units, 10)
	if in.Nanos != 0 {
		decimal := fmt.Sprintf("%9d", in.Nanos)
		decimal = strings.TrimRight(decimal, "0")
		s += "." + decimal
	}

	if in.CurrencyCode != "" {
		s += " " + in.CurrencyCode
	}
	return &s
}

func money_ToProto(mapCtx *direct.MapContext, in *string) *moneypb.Money {
	if in == nil {
		return nil
	}
	tokens := strings.Fields(*in)
	if len(tokens) > 2 || len(tokens) < 1 {
		mapCtx.Errorf("invalid Money value %q (too many fields)", *in)
		return nil
	}

	out := &moneypb.Money{}
	if len(tokens) == 2 {
		out.CurrencyCode = tokens[1]
	}

	numberTokens := strings.Split(tokens[0], ".")
	if len(tokens) > 2 {
		mapCtx.Errorf("invalid Money value %q (multiple decimal points)", *in)
		return nil
	}

	units, err := strconv.ParseInt(numberTokens[0], 10, 64)
	if err != nil {
		mapCtx.Errorf("invalid Money value %q (invalid units)", *in)
		return nil
	}
	out.Units = units

	if len(numberTokens) == 2 {
		s := numberTokens[1]
		s = s + strings.Repeat("0", len(s)-9)

		nanos, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			mapCtx.Errorf("invalid Money value %q (invalid nanos)", *in)
			return nil
		}
		out.Nanos = int32(nanos)
	}

	return out
}

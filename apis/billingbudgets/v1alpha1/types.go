package v1alpha1

// Override because of Date
// +kcc:proto=google.cloud.billing.budgets.v1.CustomPeriod
type CustomPeriod struct {
	// Required. The start date must be after January 1, 2017.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.CustomPeriod.start_date
	StartDate *string `json:"startDate,omitempty"`

	// Optional. The end date of the time period. Budgets with elapsed end date
	//  won't be processed. If unset, specifies to track all usage incurred since
	//  the start_date.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.CustomPeriod.end_date
	EndDate *string `json:"endDate,omitempty"`
}

// Override because of Money
// +kcc:proto=google.cloud.billing.budgets.v1.BudgetAmount
type BudgetAmount struct {
	// A specified amount to use as the budget.
	//  `currency_code` is optional. If specified when creating a budget, it must
	//  match the currency of the billing account. If specified when updating a
	//  budget, it must match the currency_code of the existing budget.
	//  The `currency_code` is provided on output.
	// +kcc:proto:field=google.cloud.billing.budgets.v1.BudgetAmount.specified_amount
	SpecifiedAmount *string `json:"specifiedAmount,omitempty"`

	// Use the last period's actual spend as the budget for the present period.
	//  LastPeriodAmount can only be set when the budget's time period is a
	//  [Filter.calendar_period][google.cloud.billing.budgets.v1.Filter.calendar_period].
	//  It cannot be set in combination with
	//  [Filter.custom_period][google.cloud.billing.budgets.v1.Filter.custom_period].
	// +kcc:proto:field=google.cloud.billing.budgets.v1.BudgetAmount.last_period_amount
	LastPeriodAmount *LastPeriodAmount `json:"lastPeriodAmount,omitempty"`
}

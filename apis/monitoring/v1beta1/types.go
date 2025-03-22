package v1beta1

// +kcc:proto=google.monitoring.v3.AlertPolicy.AlertStrategy
type AlertPolicy_AlertStrategy struct {
	// Required for log-based alerting policies, i.e. policies with a `LogMatch`
	//  condition.
	//
	//  This limit is not implemented for alerting policies that do not have
	//  a LogMatch condition.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.AlertStrategy.notification_rate_limit
	NotificationRateLimit *AlertPolicy_AlertStrategy_NotificationRateLimit `json:"notificationRateLimit,omitempty"`

	/* NOTYET - backcompat with terraform
	// For log-based alert policies, the notification prompts is always
	//  [OPENED]. For non log-based alert policies, the notification prompts can
	//  be [OPENED] or [OPENED, CLOSED].
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.AlertStrategy.notification_prompts
	NotificationPrompts []string `json:"notificationPrompts,omitempty"`
	*/

	// If an alerting policy that was active has no data for this long, any open
	//  incidents will close
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.AlertStrategy.auto_close
	AutoClose *string `json:"autoClose,omitempty"`

	// Control how notifications will be sent out, on a per-channel basis.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.AlertStrategy.notification_channel_strategy
	NotificationChannelStrategy []AlertPolicy_AlertStrategy_NotificationChannelStrategy `json:"notificationChannelStrategy,omitempty"`
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition
type AlertPolicy_Condition_PrometheusQueryLanguageCondition struct {
	// Required. The PromQL expression to evaluate. Every evaluation cycle
	//  this expression is evaluated at the current time, and all resultant
	//  time series become pending/firing alerts. This field must not be empty.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.query
	Query *string `json:"query,omitempty"`

	// Optional. Alerts are considered firing once their PromQL expression was
	//  evaluated to be "true" for this long.
	//  Alerts whose PromQL expression was not evaluated to be "true" for
	//  long enough are considered pending.
	//  Must be a non-negative duration or missing.
	//  This field is optional. Its default value is zero.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.duration
	Duration *string `json:"duration,omitempty"`

	// Optional. How often this rule should be evaluated.
	//  Must be a positive multiple of 30 seconds or missing.
	//  This field is optional. Its default value is 30 seconds.
	//  If this PrometheusQueryLanguageCondition was generated from a
	//  Prometheus alerting rule, then this value should be taken from the
	//  enclosing rule group.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.evaluation_interval
	EvaluationInterval *string `json:"evaluationInterval,omitempty"`

	// Optional. Labels to add to or overwrite in the PromQL query result.
	//  Label names [must be
	//  valid](https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels).
	//  Label values can be [templatized by using
	//  variables](https://cloud.google.com/monitoring/alerts/doc-variables#doc-vars).
	//  The only available variable names are the names of the labels in the
	//  PromQL result, including "__name__" and "value". "labels" may be empty.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The rule group name of this alert in the corresponding
	//  Prometheus configuration file.
	//
	//  Some external tools may require this field to be populated correctly
	//  in order to refer to the original Prometheus configuration file.
	//  The rule group name and the alert name are necessary to update the
	//  relevant AlertPolicies in case the definition of the rule group changes
	//  in the future.
	//
	//  This field is optional. If this field is not empty, then it must
	//  contain a valid UTF-8 string.
	//  This field may not exceed 2048 Unicode characters in length.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.rule_group
	RuleGroup *string `json:"ruleGroup,omitempty"`

	// Optional. The alerting rule name of this alert in the corresponding
	//  Prometheus configuration file.
	//
	//  Some external tools may require this field to be populated correctly
	//  in order to refer to the original Prometheus configuration file.
	//  The rule group name and the alert name are necessary to update the
	//  relevant AlertPolicies in case the definition of the rule group changes
	//  in the future.
	//
	//  This field is optional. If this field is not empty, then it must be a
	//  [valid Prometheus label
	//  name](https://prometheus.io/docs/concepts/data_model/#metric-names-and-labels).
	//  This field may not exceed 2048 Unicode characters in length.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.alert_rule
	AlertRule *string `json:"alertRule,omitempty"`

	/* NOTYET - backcompat with terraform
	// Optional. Whether to disable metric existence validation for this
	//  condition.
	//
	//  This allows alerting policies to be defined on metrics that do not yet
	//  exist, improving advanced customer workflows such as configuring
	//  alerting policies using Terraform.
	//
	//  Users with the `monitoring.alertPolicyViewer` role are able to see the
	//  name of the non-existent metric in the alerting policy condition.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.PrometheusQueryLanguageCondition.disable_metric_validation
	DisableMetricValidation *bool `json:"disableMetricValidation,omitempty"`
	*/
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Condition
type AlertPolicy_Condition struct {
	// Required if the condition exists. The unique resource name for this
	//  condition. Its format is:
	//
	//      projects/[PROJECT_ID_OR_NUMBER]/alertPolicies/[POLICY_ID]/conditions/[CONDITION_ID]
	//
	//  `[CONDITION_ID]` is assigned by Cloud Monitoring when the
	//  condition is created as part of a new or updated alerting policy.
	//
	//  When calling the
	//  [alertPolicies.create][google.monitoring.v3.AlertPolicyService.CreateAlertPolicy]
	//  method, do not include the `name` field in the conditions of the
	//  requested alerting policy. Cloud Monitoring creates the
	//  condition identifiers and includes them in the new policy.
	//
	//  When calling the
	//  [alertPolicies.update][google.monitoring.v3.AlertPolicyService.UpdateAlertPolicy]
	//  method to update a policy, including a condition `name` causes the
	//  existing condition to be updated. Conditions without names are added to
	//  the updated policy. Existing conditions are deleted if they are not
	//  updated.
	//
	//  Best practice is to preserve `[CONDITION_ID]` if you make only small
	//  changes, such as those to condition thresholds, durations, or trigger
	//  values.  Otherwise, treat the change as a new condition and let the
	//  existing condition be deleted.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.name
	Name *string `json:"name,omitempty"`

	// A short name or phrase used to identify the condition in dashboards,
	//  notifications, and incidents. To avoid confusion, don't use the same
	//  display name for multiple conditions in the same policy.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A condition that compares a time series against a threshold.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.condition_threshold
	ConditionThreshold *AlertPolicy_Condition_MetricThreshold `json:"conditionThreshold,omitempty"`

	// A condition that checks that a time series continues to
	//  receive new data points.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.condition_absent
	ConditionAbsent *AlertPolicy_Condition_MetricAbsence `json:"conditionAbsent,omitempty"`

	// A condition that checks for log messages matching given constraints. If
	//  set, no other conditions can be present.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.condition_matched_log
	ConditionMatchedLog *AlertPolicy_Condition_LogMatch `json:"conditionMatchedLog,omitempty"`

	// A condition that uses the Monitoring Query Language to define
	//  alerts.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.condition_monitoring_query_language
	ConditionMonitoringQueryLanguage *AlertPolicy_Condition_MonitoringQueryLanguageCondition `json:"conditionMonitoringQueryLanguage,omitempty"`

	// A condition that uses the Prometheus query language to define alerts.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.condition_prometheus_query_language
	ConditionPrometheusQueryLanguage *AlertPolicy_Condition_PrometheusQueryLanguageCondition `json:"conditionPrometheusQueryLanguage,omitempty"`

	/* NOTYET - backcompat with terraform
	// A condition that periodically evaluates a SQL query result.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Condition.condition_sql
	ConditionSQL *AlertPolicy_Condition_SQLCondition `json:"conditionSQL,omitempty"`
	*/
}

// +kcc:proto=google.monitoring.v3.AlertPolicy.Documentation
type AlertPolicy_Documentation struct {
	// The body of the documentation, interpreted according to `mime_type`.
	//  The content may not exceed 8,192 Unicode characters and may not exceed
	//  more than 10,240 bytes when encoded in UTF-8 format, whichever is
	//  smaller. This text can be [templatized by using
	//  variables](https://cloud.google.com/monitoring/alerts/doc-variables#doc-vars).
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Documentation.content
	Content *string `json:"content,omitempty"`

	// The format of the `content` field. Presently, only the value
	//  `"text/markdown"` is supported. See
	//  [Markdown](https://en.wikipedia.org/wiki/Markdown) for more information.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Documentation.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Optional. The subject line of the notification. The subject line may not
	//  exceed 10,240 bytes. In notifications generated by this policy, the
	//  contents of the subject line after variable expansion will be truncated
	//  to 255 bytes or shorter at the latest UTF-8 character boundary. The
	//  255-byte limit is recommended by [this
	//  thread](https://stackoverflow.com/questions/1592291/what-is-the-email-subject-length-limit).
	//  It is both the limit imposed by some third-party ticketing products and
	//  it is common to define textual fields in databases as VARCHAR(255).
	//
	//  The contents of the subject line can be [templatized by using
	//  variables](https://cloud.google.com/monitoring/alerts/doc-variables#doc-vars).
	//  If this field is missing or empty, a default subject line will be
	//  generated.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Documentation.subject
	Subject *string `json:"subject,omitempty"`

	/* NOTYET - backcompat with terraform
	// Optional. Links to content such as playbooks, repositories, and other
	//  resources. This field can contain up to 3 entries.
	// +kcc:proto:field=google.monitoring.v3.AlertPolicy.Documentation.links
	Links []AlertPolicy_Documentation_Link `json:"links,omitempty"`
	*/
}

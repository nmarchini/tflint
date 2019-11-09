// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsRoute53HealthCheckInvalidCloudwatchAlarmRegionRule checks the pattern is valid
type AwsRoute53HealthCheckInvalidCloudwatchAlarmRegionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	enum          []string
}

// NewAwsRoute53HealthCheckInvalidCloudwatchAlarmRegionRule returns new rule with default attributes
func NewAwsRoute53HealthCheckInvalidCloudwatchAlarmRegionRule() *AwsRoute53HealthCheckInvalidCloudwatchAlarmRegionRule {
	return &AwsRoute53HealthCheckInvalidCloudwatchAlarmRegionRule{
		resourceType:  "aws_route53_health_check",
		attributeName: "cloudwatch_alarm_region",
		max:           64,
		min:           1,
		enum: []string{
			"us-east-1",
			"us-east-2",
			"us-west-1",
			"us-west-2",
			"ca-central-1",
			"eu-central-1",
			"eu-west-1",
			"eu-west-2",
			"eu-west-3",
			"ap-east-1",
			"me-south-1",
			"ap-south-1",
			"ap-southeast-1",
			"ap-southeast-2",
			"ap-northeast-1",
			"ap-northeast-2",
			"ap-northeast-3",
			"eu-north-1",
			"sa-east-1",
			"cn-northwest-1",
			"cn-north-1",
		},
	}
}

// Name returns the rule name
func (r *AwsRoute53HealthCheckInvalidCloudwatchAlarmRegionRule) Name() string {
	return "aws_route53_health_check_invalid_cloudwatch_alarm_region"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53HealthCheckInvalidCloudwatchAlarmRegionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53HealthCheckInvalidCloudwatchAlarmRegionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53HealthCheckInvalidCloudwatchAlarmRegionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53HealthCheckInvalidCloudwatchAlarmRegionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"cloudwatch_alarm_region must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"cloudwatch_alarm_region must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`cloudwatch_alarm_region is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsDirectoryServiceLogSubscriptionInvalidDirectoryIDRule checks the pattern is valid
type AwsDirectoryServiceLogSubscriptionInvalidDirectoryIDRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsDirectoryServiceLogSubscriptionInvalidDirectoryIDRule returns new rule with default attributes
func NewAwsDirectoryServiceLogSubscriptionInvalidDirectoryIDRule() *AwsDirectoryServiceLogSubscriptionInvalidDirectoryIDRule {
	return &AwsDirectoryServiceLogSubscriptionInvalidDirectoryIDRule{
		resourceType:  "aws_directory_service_log_subscription",
		attributeName: "directory_id",
		pattern:       regexp.MustCompile(`^d-[0-9a-f]{10}$`),
	}
}

// Name returns the rule name
func (r *AwsDirectoryServiceLogSubscriptionInvalidDirectoryIDRule) Name() string {
	return "aws_directory_service_log_subscription_invalid_directory_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDirectoryServiceLogSubscriptionInvalidDirectoryIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDirectoryServiceLogSubscriptionInvalidDirectoryIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDirectoryServiceLogSubscriptionInvalidDirectoryIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDirectoryServiceLogSubscriptionInvalidDirectoryIDRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`directory_id does not match valid pattern ^d-[0-9a-f]{10}$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsIAMAccessKeyInvalidUserRule checks the pattern is valid
type AwsIAMAccessKeyInvalidUserRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMAccessKeyInvalidUserRule returns new rule with default attributes
func NewAwsIAMAccessKeyInvalidUserRule() *AwsIAMAccessKeyInvalidUserRule {
	return &AwsIAMAccessKeyInvalidUserRule{
		resourceType:  "aws_iam_access_key",
		attributeName: "user",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMAccessKeyInvalidUserRule) Name() string {
	return "aws_iam_access_key_invalid_user"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMAccessKeyInvalidUserRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMAccessKeyInvalidUserRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMAccessKeyInvalidUserRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMAccessKeyInvalidUserRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"user must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"user must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`user does not match valid pattern ^[\w+=,.@-]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

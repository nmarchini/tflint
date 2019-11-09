// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsIAMRolePolicyInvalidNameRule checks the pattern is valid
type AwsIAMRolePolicyInvalidNameRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsIAMRolePolicyInvalidNameRule returns new rule with default attributes
func NewAwsIAMRolePolicyInvalidNameRule() *AwsIAMRolePolicyInvalidNameRule {
	return &AwsIAMRolePolicyInvalidNameRule{
		resourceType:  "aws_iam_role_policy",
		attributeName: "name",
		max:           128,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=,.@-]+$`),
	}
}

// Name returns the rule name
func (r *AwsIAMRolePolicyInvalidNameRule) Name() string {
	return "aws_iam_role_policy_invalid_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsIAMRolePolicyInvalidNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsIAMRolePolicyInvalidNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsIAMRolePolicyInvalidNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsIAMRolePolicyInvalidNameRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"name must be 128 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"name must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`name does not match valid pattern ^[\w+=,.@-]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

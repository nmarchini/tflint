// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCloudformationStackSetInvalidDescriptionRule checks the pattern is valid
type AwsCloudformationStackSetInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCloudformationStackSetInvalidDescriptionRule returns new rule with default attributes
func NewAwsCloudformationStackSetInvalidDescriptionRule() *AwsCloudformationStackSetInvalidDescriptionRule {
	return &AwsCloudformationStackSetInvalidDescriptionRule{
		resourceType:  "aws_cloudformation_stack_set",
		attributeName: "description",
		max:           1024,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCloudformationStackSetInvalidDescriptionRule) Name() string {
	return "aws_cloudformation_stack_set_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudformationStackSetInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCloudformationStackSetInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudformationStackSetInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudformationStackSetInvalidDescriptionRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 1024 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"description must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

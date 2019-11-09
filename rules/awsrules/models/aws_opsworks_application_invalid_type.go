// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsOpsworksApplicationInvalidTypeRule checks the pattern is valid
type AwsOpsworksApplicationInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsOpsworksApplicationInvalidTypeRule returns new rule with default attributes
func NewAwsOpsworksApplicationInvalidTypeRule() *AwsOpsworksApplicationInvalidTypeRule {
	return &AwsOpsworksApplicationInvalidTypeRule{
		resourceType:  "aws_opsworks_application",
		attributeName: "type",
		enum: []string{
			"aws-flow-ruby",
			"java",
			"rails",
			"php",
			"nodejs",
			"static",
			"other",
		},
	}
}

// Name returns the rule name
func (r *AwsOpsworksApplicationInvalidTypeRule) Name() string {
	return "aws_opsworks_application_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsOpsworksApplicationInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsOpsworksApplicationInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsOpsworksApplicationInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsOpsworksApplicationInvalidTypeRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`type is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

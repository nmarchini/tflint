// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsDynamoDBTableInvalidBillingModeRule checks the pattern is valid
type AwsDynamoDBTableInvalidBillingModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsDynamoDBTableInvalidBillingModeRule returns new rule with default attributes
func NewAwsDynamoDBTableInvalidBillingModeRule() *AwsDynamoDBTableInvalidBillingModeRule {
	return &AwsDynamoDBTableInvalidBillingModeRule{
		resourceType:  "aws_dynamodb_table",
		attributeName: "billing_mode",
		enum: []string{
			"PROVISIONED",
			"PAY_PER_REQUEST",
		},
	}
}

// Name returns the rule name
func (r *AwsDynamoDBTableInvalidBillingModeRule) Name() string {
	return "aws_dynamodb_table_invalid_billing_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsDynamoDBTableInvalidBillingModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsDynamoDBTableInvalidBillingModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsDynamoDBTableInvalidBillingModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsDynamoDBTableInvalidBillingModeRule) Check(runner *tflint.Runner) error {
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
					`billing_mode is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

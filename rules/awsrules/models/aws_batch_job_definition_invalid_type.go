// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsBatchJobDefinitionInvalidTypeRule checks the pattern is valid
type AwsBatchJobDefinitionInvalidTypeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsBatchJobDefinitionInvalidTypeRule returns new rule with default attributes
func NewAwsBatchJobDefinitionInvalidTypeRule() *AwsBatchJobDefinitionInvalidTypeRule {
	return &AwsBatchJobDefinitionInvalidTypeRule{
		resourceType:  "aws_batch_job_definition",
		attributeName: "type",
		enum: []string{
			"container",
			"multinode",
		},
	}
}

// Name returns the rule name
func (r *AwsBatchJobDefinitionInvalidTypeRule) Name() string {
	return "aws_batch_job_definition_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsBatchJobDefinitionInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsBatchJobDefinitionInvalidTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsBatchJobDefinitionInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsBatchJobDefinitionInvalidTypeRule) Check(runner *tflint.Runner) error {
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

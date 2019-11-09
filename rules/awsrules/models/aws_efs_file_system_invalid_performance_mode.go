// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsEfsFileSystemInvalidPerformanceModeRule checks the pattern is valid
type AwsEfsFileSystemInvalidPerformanceModeRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsEfsFileSystemInvalidPerformanceModeRule returns new rule with default attributes
func NewAwsEfsFileSystemInvalidPerformanceModeRule() *AwsEfsFileSystemInvalidPerformanceModeRule {
	return &AwsEfsFileSystemInvalidPerformanceModeRule{
		resourceType:  "aws_efs_file_system",
		attributeName: "performance_mode",
		enum: []string{
			"generalPurpose",
			"maxIO",
		},
	}
}

// Name returns the rule name
func (r *AwsEfsFileSystemInvalidPerformanceModeRule) Name() string {
	return "aws_efs_file_system_invalid_performance_mode"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEfsFileSystemInvalidPerformanceModeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEfsFileSystemInvalidPerformanceModeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEfsFileSystemInvalidPerformanceModeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEfsFileSystemInvalidPerformanceModeRule) Check(runner *tflint.Runner) error {
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
					`performance_mode is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

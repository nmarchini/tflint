// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsEcrRepositoryPolicyInvalidPolicyRule checks the pattern is valid
type AwsEcrRepositoryPolicyInvalidPolicyRule struct {
	resourceType  string
	attributeName string
	max           int
}

// NewAwsEcrRepositoryPolicyInvalidPolicyRule returns new rule with default attributes
func NewAwsEcrRepositoryPolicyInvalidPolicyRule() *AwsEcrRepositoryPolicyInvalidPolicyRule {
	return &AwsEcrRepositoryPolicyInvalidPolicyRule{
		resourceType:  "aws_ecr_repository_policy",
		attributeName: "policy",
		max:           10240,
	}
}

// Name returns the rule name
func (r *AwsEcrRepositoryPolicyInvalidPolicyRule) Name() string {
	return "aws_ecr_repository_policy_invalid_policy"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsEcrRepositoryPolicyInvalidPolicyRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsEcrRepositoryPolicyInvalidPolicyRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsEcrRepositoryPolicyInvalidPolicyRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsEcrRepositoryPolicyInvalidPolicyRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"policy must be 10240 characters or less",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

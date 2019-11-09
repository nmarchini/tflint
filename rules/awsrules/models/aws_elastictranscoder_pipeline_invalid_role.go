// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsElastictranscoderPipelineInvalidRoleRule checks the pattern is valid
type AwsElastictranscoderPipelineInvalidRoleRule struct {
	resourceType  string
	attributeName string
	pattern       *regexp.Regexp
}

// NewAwsElastictranscoderPipelineInvalidRoleRule returns new rule with default attributes
func NewAwsElastictranscoderPipelineInvalidRoleRule() *AwsElastictranscoderPipelineInvalidRoleRule {
	return &AwsElastictranscoderPipelineInvalidRoleRule{
		resourceType:  "aws_elastictranscoder_pipeline",
		attributeName: "role",
		pattern:       regexp.MustCompile(`^arn:aws:iam::\w{12}:role/.+$`),
	}
}

// Name returns the rule name
func (r *AwsElastictranscoderPipelineInvalidRoleRule) Name() string {
	return "aws_elastictranscoder_pipeline_invalid_role"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsElastictranscoderPipelineInvalidRoleRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsElastictranscoderPipelineInvalidRoleRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsElastictranscoderPipelineInvalidRoleRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsElastictranscoderPipelineInvalidRoleRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`role does not match valid pattern ^arn:aws:iam::\w{12}:role/.+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

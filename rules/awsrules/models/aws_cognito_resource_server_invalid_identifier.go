// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsCognitoResourceServerInvalidIdentifierRule checks the pattern is valid
type AwsCognitoResourceServerInvalidIdentifierRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsCognitoResourceServerInvalidIdentifierRule returns new rule with default attributes
func NewAwsCognitoResourceServerInvalidIdentifierRule() *AwsCognitoResourceServerInvalidIdentifierRule {
	return &AwsCognitoResourceServerInvalidIdentifierRule{
		resourceType:  "aws_cognito_resource_server",
		attributeName: "identifier",
		max:           256,
		min:           1,
		pattern:       regexp.MustCompile(`^[\x21\x23-\x5B\x5D-\x7E]+$`),
	}
}

// Name returns the rule name
func (r *AwsCognitoResourceServerInvalidIdentifierRule) Name() string {
	return "aws_cognito_resource_server_invalid_identifier"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCognitoResourceServerInvalidIdentifierRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCognitoResourceServerInvalidIdentifierRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCognitoResourceServerInvalidIdentifierRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCognitoResourceServerInvalidIdentifierRule) Check(runner *tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"identifier must be 256 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"identifier must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssue(
					r,
					`identifier does not match valid pattern ^[\x21\x23-\x5B\x5D-\x7E]+$`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

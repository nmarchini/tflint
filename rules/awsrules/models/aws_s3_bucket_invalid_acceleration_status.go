// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsS3BucketInvalidAccelerationStatusRule checks the pattern is valid
type AwsS3BucketInvalidAccelerationStatusRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsS3BucketInvalidAccelerationStatusRule returns new rule with default attributes
func NewAwsS3BucketInvalidAccelerationStatusRule() *AwsS3BucketInvalidAccelerationStatusRule {
	return &AwsS3BucketInvalidAccelerationStatusRule{
		resourceType:  "aws_s3_bucket",
		attributeName: "acceleration_status",
		enum: []string{
			"Enabled",
			"Suspended",
		},
	}
}

// Name returns the rule name
func (r *AwsS3BucketInvalidAccelerationStatusRule) Name() string {
	return "aws_s3_bucket_invalid_acceleration_status"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsS3BucketInvalidAccelerationStatusRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsS3BucketInvalidAccelerationStatusRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsS3BucketInvalidAccelerationStatusRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsS3BucketInvalidAccelerationStatusRule) Check(runner *tflint.Runner) error {
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
					`acceleration_status is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

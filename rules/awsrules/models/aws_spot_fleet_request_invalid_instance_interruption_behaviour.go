// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint/tflint"
)

// AwsSpotFleetRequestInvalidInstanceInterruptionBehaviourRule checks the pattern is valid
type AwsSpotFleetRequestInvalidInstanceInterruptionBehaviourRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSpotFleetRequestInvalidInstanceInterruptionBehaviourRule returns new rule with default attributes
func NewAwsSpotFleetRequestInvalidInstanceInterruptionBehaviourRule() *AwsSpotFleetRequestInvalidInstanceInterruptionBehaviourRule {
	return &AwsSpotFleetRequestInvalidInstanceInterruptionBehaviourRule{
		resourceType:  "aws_spot_fleet_request",
		attributeName: "instance_interruption_behaviour",
		enum: []string{
			"hibernate",
			"stop",
			"terminate",
		},
	}
}

// Name returns the rule name
func (r *AwsSpotFleetRequestInvalidInstanceInterruptionBehaviourRule) Name() string {
	return "aws_spot_fleet_request_invalid_instance_interruption_behaviour"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSpotFleetRequestInvalidInstanceInterruptionBehaviourRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSpotFleetRequestInvalidInstanceInterruptionBehaviourRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSpotFleetRequestInvalidInstanceInterruptionBehaviourRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSpotFleetRequestInvalidInstanceInterruptionBehaviourRule) Check(runner *tflint.Runner) error {
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
					`instance_interruption_behaviour is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}

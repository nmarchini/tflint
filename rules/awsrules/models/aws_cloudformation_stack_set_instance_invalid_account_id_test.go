// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsCloudformationStackSetInstanceInvalidAccountIDRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudformation_stack_set_instance" "foo" {
	account_id = "1234567890123"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudformationStackSetInstanceInvalidAccountIDRule(),
					Message: `account_id does not match valid pattern ^[0-9]{12}$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudformation_stack_set_instance" "foo" {
	account_id = "123456789012"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudformationStackSetInstanceInvalidAccountIDRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}

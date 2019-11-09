// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"testing"

	"github.com/terraform-linters/tflint/tflint"
)

func Test_AwsCloudhsmV2HsmInvalidIPAddressRule(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected tflint.Issues
	}{
		{
			Name: "It includes invalid characters",
			Content: `
resource "aws_cloudhsm_v2_hsm" "foo" {
	ip_address = "2001:4860:4860::8888"
}`,
			Expected: tflint.Issues{
				{
					Rule:    NewAwsCloudhsmV2HsmInvalidIPAddressRule(),
					Message: `ip_address does not match valid pattern ^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`,
				},
			},
		},
		{
			Name: "It is valid",
			Content: `
resource "aws_cloudhsm_v2_hsm" "foo" {
	ip_address = "8.8.8.8"
}`,
			Expected: tflint.Issues{},
		},
	}

	rule := NewAwsCloudhsmV2HsmInvalidIPAddressRule()

	for _, tc := range cases {
		runner := tflint.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		tflint.AssertIssuesWithoutRange(t, tc.Expected, runner.Issues)
	}
}

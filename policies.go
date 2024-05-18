package awsiamroles

import (
	"encoding/json"

	"github.com/awsdocs/aws-doc-sdk-examples/gov2/iam/actions"
	"github.com/micahhausler/aws-iam-policy/policy"
)

func ParsePolicy(b []byte) (*policy.Policy, error) {
	pol := &policy.Policy{}
	return pol, json.Unmarshal(b, pol)
}

func ParsePolicyDocument(b []byte) (*actions.PolicyDocument, error) {
	// https://docs.aws.amazon.com/cdk/api/v2/docs/aws-cdk-lib.aws_iam.PolicyDocument.html
	// https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/iam/actions/policies.go
	pd := &actions.PolicyDocument{}
	return pd, json.Unmarshal(b, pd)
}

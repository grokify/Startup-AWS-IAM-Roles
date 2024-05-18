package startup

import (
	_ "embed"

	"github.com/awsdocs/aws-doc-sdk-examples/gov2/iam/actions"
	awsiamroles "github.com/grokify/go-aws-iam-roles-examples"
	"github.com/micahhausler/aws-iam-policy/policy"
)

//go:embed 01_business-owner.json
var BusinessOwnerJSON []byte

func BusinessOwnerPolicyDocument() *actions.PolicyDocument {
	pd, err := awsiamroles.ParsePolicyDocument(BusinessOwnerJSON)
	if err != nil {
		panic(err)
	}
	return pd
}

func BusinessOwnerPolicy() *policy.Policy {
	pd, err := awsiamroles.ParsePolicy(BusinessOwnerJSON)
	if err != nil {
		panic(err)
	}
	return pd
}

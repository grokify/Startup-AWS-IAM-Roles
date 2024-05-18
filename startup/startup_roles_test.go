package startup

import (
	"testing"

	awsiamroles "github.com/grokify/go-aws-iam-roles-examples"
)

var parsePolicyTests = []struct {
	roleName    string
	v           []byte
	wantVersion string
	stmt0sid    string
}{
	{awsiamroles.RoleNameBusinessOwner, BusinessOwnerJSON, "2012-10-17", "Stmt1502902212539"},
}

func TestParsePolicy(t *testing.T) {
	for _, tt := range parsePolicyTests {
		pd, err := awsiamroles.ParsePolicy(tt.v)
		if err != nil {
			t.Errorf("awsiamroles.ParsePolicyDocument(...) Fail: error [%s]",
				err.Error())
		}
		if pd.Statements == nil || len(pd.Statements.Values()) == 0 {
			t.Errorf("awsiamroles.PolicyDocument as 0 statements(...) Fail on role [%s]",
				tt.roleName)
		}
		if pd.Version != tt.wantVersion {
			t.Errorf("awsiamroles.PolicyDocument Fail: mismatch on role name [%s] want [%s] got [%s]",
				tt.roleName, tt.wantVersion, pd.Version)
		}
	}
}

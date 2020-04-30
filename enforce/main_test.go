
package enforce

import (
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"testing"
)

const modelDef = `
[request_definition]
r = user, company, resource, perm

[policy_definition]
p = user, company, resource, perm

[policy_effect]
e = some(where (p.eft == allow))

[role_definition]
g = _, _
g2 = _, _

[matchers]
m = ((r.user == p.user || r.company == p.company) && (r.resource == p.resource || p.resource == '*' ) && r.perm == p.perm) || (g(r.company, p.company) && g2(r.resource, p.resource) && g2(r.perm, p.perm)) 
`

const policyDef = `
g, company1, clerkenwell_tenant
p, *, clerkenwell_tenant, door1, auth.use_door
p, *, company1, door2, auth.use_door
p, *, company1, door2, auth.blow_up
`

var enforcer Enforcer

func TestEnforce(t *testing.T) {
	uid := "user1"
	cid := "company1"
	permrequired := "auth.use_door"

	userPolicies := enforcer.enforcer.GetFilteredGroupingPolicy(0, uid)
	companyPolicies := enforcer.enforcer.GetFilteredGroupingPolicy(0, cid)

	policies := append(userPolicies, companyPolicies...)

	allPerms := [][]string{}
	for _, policy := range policies {
		allPerms = append(enforcer.enforcer.GetFilteredPolicy(1, policy[1]))
	}

	companyPerms := enforcer.enforcer.GetFilteredPolicy(1, cid)
	userPerms := enforcer.enforcer.GetFilteredPolicy(2, uid)

	allPerms = append(allPerms, companyPerms...)
	allPerms = append(allPerms, userPerms...)
	log.Println(allPerms)
	log.Println(permrequired)
	resourceids := []string{}
	for _, perm := range allPerms {
		if perm[3] == permrequired {
			resourceids = append(resourceids, perm[2])
		}
	}

	log.Println(resourceids)

	ok, err := enforcer.enforcer.Enforce("user1", "company1", "door1", "auth.use_door")

	log.Println(ok)
	log.Println(err)
	assert.Equal(t, false, true)
}
func init()  {
	m, _ := model.NewModelFromString(modelDef)
	adapter := fileadapter.NewFilteredAdapter("/tmp/policy.conf")
	ioutil.WriteFile("/tmp/policy.conf", []byte(policyDef), 0644)

	enforcer = NewEnforcer(m, adapter)
	enforcer.enforcer.LoadPolicy()
}
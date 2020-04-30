package enforce

import (
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

const modelDef = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`

const policyDef = `
p, alice, door1, auth.list_doors
p, alice, door2, auth.list_doors
p, alice, door3, auth.list_doors
p, dave, door1, auth.list_doors
`

var enforcer Enforcer

func TestEnforce(t *testing.T) {
	msg := NewAuthMessage([]string{"auth.list_doors"}, []string{"door1"}, false)
	res, _ := enforcer.Enforce([]string{"alice"}, msg)
	assert.Equal(t, true, res)
}

func TestHydrate(t *testing.T) {
	msg := NewAuthMessage([]string{"auth.list_doors"}, nil, false)
	outMsg, _ := enforcer.Hydrate([]string{"alice"}, msg)
	assert.ElementsMatch(t, []string{"door1", "door2", "door3"}, outMsg.XXX_AuthResourceIds())

	msg = NewAuthMessage([]string{"auth.list_doors"}, nil, false)
	outMsg, _ = enforcer.Hydrate([]string{"dave"}, msg)
	assert.Equal(t, []string{"door1"}, outMsg.XXX_AuthResourceIds())

	msg = NewAuthMessage([]string{"auth.list_doors"}, nil, false)
	_, err := enforcer.Hydrate([]string{"tim"}, msg)
	assert.Error(t, err)
}

func init()  {
	m, _ := model.NewModelFromString(modelDef)
	adapter := fileadapter.NewFilteredAdapter("/tmp/policy.conf")
	ioutil.WriteFile("/tmp/policy.conf", []byte(policyDef), 0644)
	ioutil.WriteFile("/tmp/policy.conf", []byte(policyDef), 0644)

	enforcer = NewEnforcer(m, adapter)
}
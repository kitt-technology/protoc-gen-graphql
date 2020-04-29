package enforce

import (
	"context"
	"github.com/casbin/casbin/v2"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/go-kit/kit/auth/casbin"
	"github.com/kitt-technology/kitt/lib/go/log"
	"io/ioutil"
)

func Enforce(ctx context.Context, model casbin.) (status bool, err error) {

	if err != nil {
		return
	}
	// See https://casbin.org/editor/ for reference
	text :=
		[]byte(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
`)

	p := []byte(`
p, ` + user.Id +`, ` + *msg.XXX_AuthResourceId() +`, OPEN_DOOR
`)

	err = ioutil.WriteFile("/tmp/policy.conf", p, 0644)
	err = ioutil.WriteFile("/tmp/model.conf", text, 0644)
	if err != nil {
		return
	}
	adapter := fileadapter.NewFilteredAdapter( "/tmp/policy.conf")

	e, err := casbin.NewEnforcer()
	e.InitWithAdapter("/tmp/model.conf", adapter)

	e.GetPolicy()
	if err != nil {
		return
	}
	e.LoadPolicy()

	if err != nil {
		return
	}


	if msg.XXX_PullResourceIds() {
		e.LoadFilteredPolicy(fileadapter.Filter{
			P: []string{},
		})
	}

	for _, perm := range msg.XXX_AuthPermissions() {
		userAllowed, err := e.Enforce(user.Id, msg.XXX_AuthResourceId(), perm)
		companyAllowed, err := e.Enforce(user.Id, msg.XXX_AuthResourceId(), perm)


		if err == nil && (userAllowed || companyAllowed) {
			log.WithContext(ctx).Info("User doesnt have permissions")
			log.WithContext(ctx).Info(status)
			return
		}
	}

	return e.Enforce(msg)
}


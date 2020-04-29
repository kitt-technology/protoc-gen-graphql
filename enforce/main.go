package enforce

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/kitt-technology/protoc-gen-auth/auth"
)

func NewEnforcer(model model.Model, adapter persist.FilteredAdapter) Enforcer {
	var err error
	enforcer, err := casbin.NewEnforcer()
	if err != nil {
		panic(err)
	}
	enforcer.InitWithModelAndAdapter(model, adapter)

	return Enforcer {
		enforcer,
	}
}

type Enforcer struct {
	enforcer *casbin.Enforcer
}

func (e Enforcer) Hydate(id string, msg auth.AuthMessage,) (auth.AuthMessage, error) {
	var filters []string
	filters = append(filters, id)

	err := e.enforcer.LoadFilteredPolicy(&fileadapter.Filter{P: filters})
	if err != nil {
		return nil, err
	}

	permissions := e.enforcer.GetPermissionsForUser(id)

	// count permissions for each returned object
	userPermsToObj := map[string]map[string]string{}
	for _, policy := range permissions {
		obj, act := policy[1], policy[2]
		if _, ok := userPermsToObj[obj]; !ok {
			userPermsToObj[obj] = make(map[string]string)
		}
		userPermsToObj[obj][act] = act
	}

	// valid objects are those with the correct amount of permissions
	permissionsRequired := len(msg.XXX_AuthPermissions())
	var validResourceIds []string
	for key, perms := range userPermsToObj {
		if len(perms) >= permissionsRequired {
			validResourceIds = append(validResourceIds, key)
		}
	}

	if len(validResourceIds) == 0 {
		return nil, fmt.Errorf("no permitted resources")
	}

	return msg.XXX_SetAuthResourceIds(validResourceIds), nil
}

func (e Enforcer) Enforce(id string, msg auth.AuthMessage) (bool, error) {
	var resourceIds []string
	if msg.XXX_AuthResourceId() != nil {
		resourceIds = append(resourceIds, *msg.XXX_AuthResourceId())
	}
	if msg.XXX_AuthResourceIds() != nil {
		resourceIds = append(resourceIds, msg.XXX_AuthResourceIds()...)
	}

	var filters []string
	filters = append(filters, id)

	e.enforcer.LoadFilteredPolicy(&fileadapter.Filter{P: filters})

	for _, resourceId := range resourceIds {
		for _, perm := range msg.XXX_AuthPermissions() {

			ok, err := e.enforcer.Enforce(id, resourceId, perm)

			// fail if any resource not permitted
			if !ok || err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

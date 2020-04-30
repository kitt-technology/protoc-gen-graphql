package enforce

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
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

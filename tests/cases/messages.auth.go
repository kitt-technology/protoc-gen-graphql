package cases

import (
	"github.com/kitt-technology/protoc-gen-auth/auth"
)

func (x *TestCommand) XXX_AuthPermission() string {
	return "test_command_permission1"
}

func (x *TestCommand) XXX_SetAuthResourceIds(resourceIds []string) auth.AuthResourceMessage {
	x.TestIds = resourceIds
	return x
}

func (x *TestCommand) XXX_AuthResourceIds() []string {
	resourceIds := []string{}
	resourceIds = append(resourceIds, x.TestId)
	resourceIds = append(resourceIds, x.TestIds...)
	return resourceIds
}

func (x *TestCommandNoIds) XXX_AuthPermission() string {
	return "test_command_permission2"
}

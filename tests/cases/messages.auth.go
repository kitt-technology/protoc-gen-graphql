package cases

import (
	"github.com/kitt-technology/protoc-gen-auth/auth"
)

func (x *TestCommand) XXX_AuthPermission() string {
	return "test_command_permission1"
}

func (x *TestCommand) XXX_AuthResourceIds() []string {
	resourceIds := []string{}
	resourceIds = append(resourceIds, x.TestId)
	resourceIds = append(resourceIds, x.TestIds...)
	return resourceIds
}

func (x *TestCommand) XXX_SetAuthResourceIds(resourceIds []string) auth.AuthMessage {
	x.TestIds = resourceIds
	return x
}

func (x *TestCommand) XXX_PullResourceIds() bool {
	return false
}

func (x *TestCommandNoIds) XXX_AuthPermission() string {
	return "test_command_permission1"
}

func (x *TestCommandNoIds) XXX_AuthResourceIds() []string {
	resourceIds := []string{}

	return resourceIds
}

func (x *TestCommandNoIds) XXX_SetAuthResourceIds(resourceIds []string) auth.AuthMessage {

	return x
}

func (x *TestCommandNoIds) XXX_PullResourceIds() bool {
	return false
}

func (x *TestCommandPullIds) XXX_AuthPermission() string {
	return "test_command_permission1"
}

func (x *TestCommandPullIds) XXX_AuthResourceIds() []string {
	resourceIds := []string{}

	resourceIds = append(resourceIds, x.TestIds...)
	return resourceIds
}

func (x *TestCommandPullIds) XXX_SetAuthResourceIds(resourceIds []string) auth.AuthMessage {
	x.TestIds = resourceIds
	return x
}

func (x *TestCommandPullIds) XXX_PullResourceIds() bool {
	return true
}

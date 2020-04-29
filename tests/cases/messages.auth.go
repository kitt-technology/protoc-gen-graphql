package cases

import (
	"github.com/kitt-technology/protoc-gen-auth/auth"
)

func (x *TestCommand) XXX_AuthPermissions() []string {
	return []string{
		"test_command_permission1",
		"test_command_permission2",
	}
}

func (x *TestCommand) XXX_AuthResourceId() *string {
	return &x.TestId
}

func (x *TestCommand) XXX_AuthResourceIds() []string {
	return x.TestIds
}

func (x *TestCommand) XXX_SetAuthResourceId(resourceId string) auth.AuthMessage {
	x.TestId = resourceId
	return x
}

func (x *TestCommand) XXX_SetAuthResourceIds(resourceIds []string) auth.AuthMessage {
	x.TestIds = resourceIds
	return x
}

func (x *TestCommand) XXX_PullResourceIds() bool {
	return false
}

func (x *TestCommandNoIds) XXX_AuthPermissions() []string {
	return []string{
		"test_command_permission1",
		"test_command_permission2",
	}
}

func (x *TestCommandNoIds) XXX_AuthResourceId() *string {
	return nil
}

func (x *TestCommandNoIds) XXX_AuthResourceIds() []string {
	return nil
}

func (x *TestCommandNoIds) XXX_SetAuthResourceId(resourceId string) auth.AuthMessage {

	return x
}

func (x *TestCommandNoIds) XXX_SetAuthResourceIds(resourceIds []string) auth.AuthMessage {

	return x
}

func (x *TestCommandNoIds) XXX_PullResourceIds() bool {
	return false
}

func (x *TestCommandPullIds) XXX_AuthPermissions() []string {
	return []string{
		"test_command_permission1",
		"test_command_permission2",
	}
}

func (x *TestCommandPullIds) XXX_AuthResourceId() *string {
	return nil
}

func (x *TestCommandPullIds) XXX_AuthResourceIds() []string {
	return x.TestIds
}

func (x *TestCommandPullIds) XXX_SetAuthResourceId(resourceId string) auth.AuthMessage {

	return x
}

func (x *TestCommandPullIds) XXX_SetAuthResourceIds(resourceIds []string) auth.AuthMessage {
	x.TestIds = resourceIds
	return x
}

func (x *TestCommandPullIds) XXX_PullResourceIds() bool {
	return true
}

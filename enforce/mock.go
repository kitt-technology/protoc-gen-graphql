package enforce

import (
	"github.com/kitt-technology/protoc-gen-auth/auth"
)

func NewAuthMessage(perms []string, resourceIds []string, pull bool) MockAuthMessage {
	return MockAuthMessage{
		perms: perms,
		resourceIds: resourceIds,
		pull: pull,
	}
}

type MockAuthMessage struct {
	auth.AuthMessage
	perms []string
	resourceIds []string
	pull bool
}

func (a MockAuthMessage) XXX_AuthPermissions() []string {
	return a.perms
}

func (a MockAuthMessage) XXX_AuthResourceIds() []string {
	return a.resourceIds
}

func (a MockAuthMessage) XXX_SetAuthResourceIds(resourceIds []string) auth.AuthMessage {
	a.resourceIds = resourceIds
	return a
}

func (a MockAuthMessage) XXX_PullResourceIds() bool {
	return a.pull
}

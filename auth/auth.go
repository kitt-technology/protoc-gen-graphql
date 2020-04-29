package auth

type AuthMessage interface {
	XXX_AuthPermissions() []string
	XXX_AuthResourceId() *string
	XXX_AuthResourceIds() []string
	XXX_SetAuthResourceId(resourceId string)
	XXX_SetAuthResourceIds(resourceIds string)
}

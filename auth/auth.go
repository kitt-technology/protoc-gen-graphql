package auth

type AuthMessage interface {
	XXX_AuthPermission() string
}

type AuthResourceMessage interface {
	XXX_AuthResourceIds() []string
	XXX_SetAuthResourceIds(resourceIds []string) AuthMessage
}
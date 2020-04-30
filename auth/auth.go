package auth

type AuthMessage interface {
	XXX_AuthPermission() string
	XXX_AuthResourceIds() []string
	XXX_SetAuthResourceIds(resourceIds []string) AuthMessage
    XXX_PullResourceIds() bool
}

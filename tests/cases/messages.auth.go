package cases

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

func (x *TestCommand) XXX_SetAuthResourceId(resourceId string) {
	x.TestId = resourceId
}

func (x *TestCommand) XXX_SetAuthResourceIds(resourceIds []string) {
	x.TestIds = resourceIds
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

func (x *TestCommandNoIds) XXX_SetAuthResourceId(resourceId string) {
	return
}

func (x *TestCommandNoIds) XXX_SetAuthResourceIds(resourceIds []string) {
	return
}

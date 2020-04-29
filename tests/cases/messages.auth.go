package cases

func (x *TestCommand) XXX_AuthPermissions() []string {
	return []string{
		"test_command_permission1",
		"test_command_permission2",
	}
}

func (x *TestCommand) XXX_AuthResourceId() string {
	return x.TestId
}

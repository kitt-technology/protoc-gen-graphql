package cases

func (x *OpenDoorCommand) XXX_AuthPermissions() []string {
	return []string{
		"open_door",
	}
}

func (x *OpenDoorCommand) XXX_AuthResourceId() *string {
	return &x.Id
}

func (x *OpenDoorCommand) XXX_AuthResourceIds() []string {
	return x.Ids
}

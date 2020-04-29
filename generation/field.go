package generation

const fieldTpl = `
func (x *TestCommand) XXX_ResourceId() string {
	if x != nil {
		return x.s
	}
	return ""
}
`

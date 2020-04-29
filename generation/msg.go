package generation

import (
    "bytes"
    "html/template"
)

const msgTpl = `
func (x *{{ .Type }}) XXX_AuthPermissions() []string {
	return []string{
    {{ range $perm := .Permissions }}"{{ $perm }}",
    {{ end }}
    }
}

func (x *{{ .Type }}) XXX_AuthResourceId() string {
	return x.{{ .ResourceId }}
}
`

type AuthMessage struct {
    Type string
    Permissions []string
    ResourceId string
}

func (a AuthMessage) Generate() string {
    var buf bytes.Buffer
    mTpl, err := template.New("msg").Parse(msgTpl)
    if err != nil {
        panic(err)
    }
    mTpl.Execute(&buf, a)

    return buf.String()
}

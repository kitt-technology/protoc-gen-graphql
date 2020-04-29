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

func (x *{{ .Type }}) XXX_AuthResourceId() *string {
    {{ if .ResourceId }} return &x.{{ .ResourceId }}{{ else }} return nil {{ end }}
}

func (x *{{ .Type }}) XXX_AuthResourceIds() []string {
    {{ if .ResourceIds }} return x.{{ .ResourceIds }}{{ else }}return nil {{ end }}
}

func (x *{{ .Type }}) XXX_SetAuthResourceId(resourceId string) {
    {{ if .ResourceId }}x.{{ .ResourceId }} = resourceId{{ else }}return {{ end }}
}

func (x *{{ .Type }}) XXX_SetAuthResourceIds(resourceIds []string) {
    {{ if .ResourceIds }}x.{{ .ResourceIds }} = resourceIds{{ else }}return {{ end }}
}

func (x *{{ .Type }}) XXX_PullResourceIds() bool {
    return {{ if .PullResourceIds }}true{{ else }}false{{ end }}
}
`

type AuthMessage struct {
    Type string
    Permissions []string
    ResourceId string
    ResourceIds string
    PullResourceIds bool
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

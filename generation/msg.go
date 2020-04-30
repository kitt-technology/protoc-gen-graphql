package generation

import (
    "bytes"
    "html/template"
)

const msgTpl = `
func (x *{{ .Type }}) XXX_AuthPermission() string {
	return "{{ .Permission }}"
}

func (x *{{ .Type }}) XXX_AuthResourceIds() []string {
    resourceIds := []string{}
    {{ if .ResourceId }} resourceIds = append(resourceIds,  x.{{ .ResourceId }}){{ end }}
    {{ if .ResourceIds }} resourceIds = append(resourceIds,  x.{{ .ResourceIds }}...){{ end }}
    return resourceIds
}

func (x *{{ .Type }}) XXX_SetAuthResourceIds(resourceIds []string) auth.AuthMessage {
    {{ if .ResourceIds }}x.{{ .ResourceIds }} = resourceIds{{ end }}
	return x
}

func (x *{{ .Type }}) XXX_PullResourceIds() bool {
    return {{ if .PullResourceIds }}true{{ else }}false{{ end }}
}
`

type AuthMessage struct {
    Type string
    Permission string
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

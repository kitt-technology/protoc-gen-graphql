package generation

import (
    "bytes"
    "html/template"
)

const msgTpl = `
func (x *{{ .Type }}) XXX_AuthPermission() string {
	return "{{ .Permission }}"
}

{{if or .ResourceIds}}
func (x *{{ .Type }}) XXX_SetAuthResourceIds(resourceIds []string) auth.AuthResourceMessage {
    {{ if .ResourceIds }}x.{{ .ResourceIds }} = resourceIds{{ end }}
	return x
}
{{end}}


{{if .ResourceIds}}
func (x *{{ .Type }}) XXX_AuthResourceIds() []string {
    resourceIds := []string{}
    {{ if .ResourceId }} resourceIds = append(resourceIds,  x.{{ .ResourceId }}){{ end }}
    {{ if .ResourceIds }} resourceIds = append(resourceIds,  x.{{ .ResourceIds }}...){{ end }}
    return resourceIds
}
{{end}}

`

type AuthMessage struct {
    Type string
    Permission string
    ResourceId *string
    ResourceIds *string
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

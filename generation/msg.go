package generation

import (
    "bytes"
    "html/template"
)

const msgTpl = `
func (x *{{ .Type }}) AuthMeta() string {
	return ""
}
`

type AuthMessage struct {
    Type string
    Permissions []string
    ResourceId string
}

func (a AuthMessage) Generate() string {
    var buf bytes.Buffer
    tpl, err := template.New("msg").Parse(msgTpl)
    if err != nil {
        panic(err)
    }
    tpl.Execute(&buf, a)
    return buf.String()
}

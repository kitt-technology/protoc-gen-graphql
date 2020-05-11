package generation

import (
    "bytes"
    "github.com/kitt-technology/protoc-gen-auth/auth"
    "google.golang.org/protobuf/compiler/protogen"
    "google.golang.org/protobuf/proto"
    "strings"
    "text/template"
)

const fileTpl = `
package {{ .Package }}

import (
	"github.com/kitt-technology/protoc-gen-auth/auth"
)
`

type File struct {
    Package protogen.GoPackageName
    AuthMessages []AuthMessage
}

func New(file *protogen.File) (f File)  {
    f.Package = file.GoPackageName

    for _, msg := range file.Proto.MessageType {
        authMessage := AuthMessage{
            Type: msg.GetName(),
        }
        if msg.Options != nil {
            authMessage.Permission = proto.GetExtension(msg.Options, auth.E_MessagePermission).(string)
        }

        for _, field := range msg.Field {
            if field.Options != nil {
                // TODO use proto-gen-go functionality for field names
                name := *field.Name

                switch proto.GetExtension(field.Options, auth.E_FieldBehaviour) {
                case auth.FieldBehaviour_ID:
                    resourceId := strings.ToUpper(string(name[0])) + name[1:]
                    authMessage.ResourceId = &resourceId
                    break;
                case auth.FieldBehaviour_IDS:
                    resourceIds := strings.ToUpper(string(name[0])) + name[1:]
                    authMessage.ResourceIds = &resourceIds
                    break
                }
            }
        }

        if authMessage.Permission != "" {
            f.AuthMessages = append(f.AuthMessages, authMessage)
        }
    }
    return f
}

func (f File) ToString() string {
    var buf bytes.Buffer
    tpl, err := template.New("file").Parse(fileTpl)
    if err != nil {
        panic(err)
    }
    tpl.Execute(&buf, f)

    out := buf.String()

    for _, msg := range f.AuthMessages {
       out += msg.Generate()
    }
    return out
}

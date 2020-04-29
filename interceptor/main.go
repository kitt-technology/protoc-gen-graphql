package interceptor

import (
	"context"
	"encoding/json"
	"github.com/kitt-technology/protoc-gen-auth/auth"
	"github.com/kitt-technology/protoc-gen-auth/enforce"
	"google.golang.org/grpc"
)

type AuthInterceptor struct {
	AttrKey  string
	Enforcer enforce.Enforcer
}

func (a AuthInterceptor) Interceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		switch msg := req.(type) {
		case auth.AuthMessage:
			attributejson := ctx.Value(a.AttrKey).(string)
			var attrs []string
			json.Unmarshal([]byte(attributejson), &attrs)

			var err error
			if msg.XXX_PullResourceIds() {
				req, err = a.Enforcer.Hydrate(attrs, msg)
			} else {
				ok, err := a.Enforcer.Enforce(attrs, msg)
				if !ok {
					return nil, err
				}
			}

			if err != nil {
				return nil, err
			}
		}
		return handler(ctx, req)
	}
}
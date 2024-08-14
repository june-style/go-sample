package interceptors

import (
	"context"
	"strings"

	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/domain/services"
	"github.com/june-style/go-sample/framework/protocol/pb"
	"github.com/june-style/go-sample/interface/logs"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Authorization(authorizer services.Authorizer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		admin, err := GetMethodOptionAdmin(req, info)
		if err != nil {
			return nil, derrors.Wrap(err)
		}

		if err := authorizer.VerifyApplicationKey(ctx); err != nil {
			return nil, derrors.Wrap(err)
		}

		if !admin.DisableToAuthAccessKey {
			ctx, err = authorizer.VerifyAccessKey(ctx)
			if err != nil {
				return nil, derrors.Wrap(err)
			}
		}

		if !admin.DisableToAuthSessionId {
			if err := authorizer.VerifySession(ctx); err != nil {
				return nil, derrors.Wrap(err)
			}
		}

		// access_log.go で既にコール済のためここでのエラーハンドリングはスキップする
		rid, _ := dcontext.GetHeader(ctx, dcontext.HeaderRequestID)

		logs.Access().
			SetRequestID(rid).
			SetUserID(dcontext.GetAuthenticatedUserID(ctx)).
			Msg("Authed-User")

		return handler(ctx, req)
	}
}

var (
	ErrInvalidRequest      = derrors.NewInternal("invalid request")
	ErrInvalidServiceName  = derrors.NewInternal("invalid service name")
	ErrInvalidMethodName   = derrors.NewInternal("invalid method name")
	ErrInvalidMethodOption = derrors.NewInternal("invalid method option")
)

func GetMethodOptionAdmin(req any, info *grpc.UnaryServerInfo) (*pb.Admin, error) {
	pm, ok := req.(proto.Message)
	if !ok || info == nil || info.FullMethod == "" {
		return nil, derrors.Wrap(ErrInvalidRequest)
	}
	ifm := infoFullMethod(info.FullMethod)
	svc := pm.ProtoReflect().Descriptor().ParentFile().
		Services().ByName(protoreflect.Name(ifm.service()))
	if svc == nil {
		return nil, derrors.Wrapf(ErrInvalidServiceName, "name is %s", ifm.service())
	}
	mtd := svc.Methods().ByName(protoreflect.Name(ifm.method()))
	if mtd == nil {
		return nil, derrors.Wrapf(ErrInvalidMethodName, "name is %s", ifm.method())
	}
	opts := mtd.Options().(*descriptorpb.MethodOptions)
	if pbad, ok := proto.GetExtension(opts, pb.E_Admin).(*pb.Admin); ok {
		return pbad, nil
	}
	return nil, derrors.Wrap(ErrInvalidMethodOption)
}

type infoFullMethod string

func (s infoFullMethod) service() string {
	var i = 1
	ss := strings.Split(string(s), "/")
	if len(ss) < i+1 {
		return ""
	}
	var j = 1
	sss := strings.Split(ss[i], ".")
	if len(sss) < j+1 {
		return ""
	}
	return sss[j]
}

func (s infoFullMethod) method() string {
	var i = 2
	ss := strings.Split(string(s), "/")
	if len(ss) < i+1 {
		return ""
	}
	return ss[i]
}

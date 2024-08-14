package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/framework/protocol/pb"
	"github.com/june-style/go-sample/framework/registry/injector"
	"github.com/june-style/go-sample/interface/interceptors"
	"github.com/june-style/go-sample/interface/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logs.Info().Str("msg", msgGreeting).Msg("Starting")

	ctx := context.Background()

	cfg, _, api, err := getApp(ctx)
	if err != nil {
		logs.Fatal().SetError(err).Msg("Stopped")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grpc.Port))
	if err != nil {
		logs.Fatal().SetError(err).Msg("Stopped")
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		getUnaryServerInterceptors(api)...),
	))
	reflection.Register(grpcServer)

	pb.RegisterHomeServiceServer(grpcServer, api.Controller.Home)
	pb.RegisterSignServiceServer(grpcServer, api.Controller.Sign)

	ctx, stop := signal.NotifyContext(ctx,
		syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, os.Interrupt,
	)
	defer stop()

	go func() {
		<-ctx.Done()
		logs.Info().Str("msg", msgStopedServerBySignal).Msg("Stopped")
		grpcServer.GracefulStop()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		logs.Fatal().SetError(err).Msg("Stopped")
	}
}

func getApp(ctx context.Context) (*dconfig.Config, *injector.DBClient, *injector.API, error) {
	cfg, err := dconfig.New()
	if err != nil {
		return nil, nil, nil, err
	}

	dbc, err := injector.InitDBClient(ctx, cfg)
	if err != nil {
		return nil, nil, nil, err
	}

	api, err := injector.InitAPI(dbc)
	if err != nil {
		return nil, nil, nil, err
	}

	return cfg, dbc, api, nil
}

func getUnaryServerInterceptors(api *injector.API) []grpc.UnaryServerInterceptor {
	interceptorSet := []grpc.UnaryServerInterceptor{
		interceptors.Recovery(),
		interceptors.ErrorStatus(),
		interceptors.ErrorLog(),
		interceptors.AccessLog(),
		interceptors.Authorization(api.Service.Authorizer),
		interceptors.Initialization(api.Service.Timer),
	}

	switch {
	case api.Config.Sys.IsProd(), api.Config.Sys.IsStg(), api.Config.Sys.IsDev():
		break
	case api.Config.Sys.IsLocal():
		interceptorSet = append(interceptorSet, interceptors.DeveloperLog())
	}

	return interceptorSet
}

const (
	msgGreeting             = "Hi! Go gRPC trial by june-style!"
	msgStopedServerBySignal = "Stoped the server gracefully by signal."
)

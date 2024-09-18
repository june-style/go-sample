package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/framework/protocol/pb"
	"github.com/june-style/go-sample/interface/logs"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logs.Info().Str("msg", msgGreeting).Msg("Starting")

	cfg, err := dconfig.New()
	if err != nil {
		logs.Fatal().SetError(err).Msg("Stopped")
	}

	flag.Parse()
	defer glog.Flush()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", cfg.Grpc.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logs.Fatal().SetError(err).Msg("Stopped")
	}
	defer conn.Close()

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(matcher))
	if err := pb.RegisterHomeServiceHandler(ctx, mux, conn); err != nil {
		logs.Fatal().SetError(err).Msg("Stopped")
	}
	if err := pb.RegisterSignServiceHandler(ctx, mux, conn); err != nil {
		logs.Fatal().SetError(err).Msg("Stopped")
	}

	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", cfg.Grpc.GatewayPort), mux); err != nil {
		logs.Error().SetError(err).Msg("Stopped")
		glog.Fatal(err)
	}
}

func matcher(key string) (string, bool) {
	if strings.HasPrefix(strings.ToLower(key), "x-") {
		return key, true
	}
	return "", false
}

const msgGreeting = "Hi! Go gRPC Gateway trial by june-style!"

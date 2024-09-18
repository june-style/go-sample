package main

import (
	"context"
	"fmt"
	"os"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/interface/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	cmdHomeGet = "home-get"
	cmdSignIn  = "sign-in"
	cmdSignUp  = "sign-up"
)

func main() {
	cmd := os.Args[1]

	logs.Info().Str("cmd", cmd).Msg("Start")

	cfg, err := dconfig.New()
	if err != nil {
		logs.Fatal().SetError(err).Msg("Finished")
	}

	conn, err := connector(cfg)
	if err != nil {
		logs.Fatal().SetError(err).Msg("Finished")
	}
	defer conn.Close()

	ctx := context.Background()
	var res any
	switch cmd {
	case cmdHomeGet:
		res, err = newHome(cfg, conn).get(ctx, os.Args[2], os.Args[3])
	case cmdSignIn:
		res, err = newSign(cfg, conn).in(ctx, os.Args[2])
	case cmdSignUp:
		res, err = newSign(cfg, conn).up(ctx, os.Args[2])
	default:
		logs.Fatal().Str("cmd", cmd).Msg("Finished")
	}

	if err != nil {
		logs.Error().SetError(err).Msg("Message")
		if s, ok := status.FromError(err); ok {
			logs.Debug().SetField("error-details", fmt.Sprintf("%+v", s.Details())).Send()
		}
	}

	logs.Debug().SetResponse(res).Msg("Response")
	logs.Info().Str("cmd", cmd).Msg("Finished")
}

func connector(cfg *dconfig.Config) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		fmt.Sprintf("local-api:%d", cfg.Grpc.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		derrors.Wrap(err)
	}
	return conn, nil
}

package interceptors_test

import (
	"context"
	"net"
	"testing"

	"github.com/june-style/go-sample/domain/dconfig"
	"github.com/june-style/go-sample/domain/dcontext"
	"github.com/june-style/go-sample/domain/derrors"
	"github.com/june-style/go-sample/domain/entities"
	"github.com/june-style/go-sample/domain/services"
	. "github.com/june-style/go-sample/interface/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/interop"
	pb "google.golang.org/grpc/interop/grpc_testing"
	"google.golang.org/grpc/test/bufconn"
)

func TestGRPC_Interceptors(t *testing.T) {
	ctx := context.TODO()

	testUnaryHandler := func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		// Note: Test code start
		// Target: Initialization()
		_, err := dcontext.GetTime(ctx)
		if err != nil {
			t.Errorf("not initialized")
		}
		return handler(ctx, req)
	}

	serviceTime, err := services.NewTimer(tcfg, &entities.Repository{})
	if err != nil {
		t.Errorf("failed to new service: %v", err)
	}

	options := []grpc.ServerOption{grpc.ChainUnaryInterceptor(
		// Note: Interceptor under test
		Recovery(),
		ErrorStatus(),
		ErrorLog(),
		AccessLog(),
		// Authorization(serviceAuth),
		Initialization(serviceTime),
		// Note: Handler showing test content
		testUnaryHandler,
	)}

	grpcServer := serve(options)
	defer t.Cleanup(func() { grpcServer.Stop() })

	dial := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(ctx,
		"dummy-target",
		grpc.WithContextDialer(dial),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Fialed to dial: %v", err)
	}
	defer t.Cleanup(func() { conn.Close() })

	client := pb.NewTestServiceClient(conn)
	interop.DoEmptyUnaryCall(ctx, client)
}

var tcfg = &dconfig.Config{
	Sys: dconfig.Sys{
		Env: dconfig.LOCAL,
		TZ:  "Asia/Tokyo",
	},
}

var lis *bufconn.Listener

func serve(options []grpc.ServerOption) *grpc.Server {
	lis = bufconn.Listen(1024 * 1024)
	grpcServer := grpc.NewServer(options...)
	pb.RegisterTestServiceServer(grpcServer, interop.NewTestServer())
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(derrors.Wrap(err))
		}
	}()
	return grpcServer
}

package cockroachgrpc

import (
	"context"
	"net"
	"os"
	"testing"

	"github.com/cockroachdb/errors/grpc/middleware"
	"github.com/hydrogen18/memlistener"
	"google.golang.org/grpc"
)

var Client EchoerClient

func TestMain(m *testing.M) {
	srv := &EchoServer{}

	lis := memlistener.NewMemoryListener()

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryServerInterceptor))
	RegisterEchoerServer(grpcServer, srv)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()

	dialOpts := []grpc.DialOption{
		grpc.WithContextDialer(func(_ context.Context, _ string) (net.Conn, error) {
			return lis.Dial("", "")
		}),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(middleware.UnaryClientInterceptor),
	}

	clientConn, err := grpc.Dial("", dialOpts...)
	if err != nil {
		panic(err)
	}

	Client = NewEchoerClient(clientConn)

	code := m.Run()

	grpcServer.Stop()
	if err := clientConn.Close(); err != nil {
		panic(err)
	}

	os.Exit(code)
}

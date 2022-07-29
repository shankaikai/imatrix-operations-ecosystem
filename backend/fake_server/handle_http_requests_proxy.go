package fake_server

import (
	"context"
	"flag"
	"net/http"

	"fmt"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gw "capstone.operations_ecosystem/backend/proto" // Update
)

//   var (
// 	// command-line options:
// 	// gRPC server endpoint
// 	grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:9090", "gRPC server endpoint")
//   )

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterWebAppServicesHandlerFromEndpoint(ctx, mux, "localhost:9090", opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":9089", mux)
}

func Proxy_main() {
	flag.Parse()
	defer glog.Flush()

	fmt.Println("Staring GW HTTP-Proxy server on 9089")

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

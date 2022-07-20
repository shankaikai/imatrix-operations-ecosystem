package server

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

func run(mainServerAddr *string, mainServerPort *int, webProxyAddr *string, webProxyPort *int) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gw.RegisterWebAppServicesHandlerFromEndpoint(ctx, mux,
		fmt.Sprintf("%s:%d", *mainServerAddr, *mainServerPort), opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", *webProxyAddr, *webProxyPort), mux)
}

func Proxy_main(mainServerAddr *string, mainServerPort *int, webProxyAddr *string, webProxyPort *int) {
	flag.Parse()
	defer glog.Flush()

	fmt.Println("Staring GW HTTP-Proxy server on", webProxyPort)

	if err := run(mainServerAddr, mainServerPort, webProxyAddr, webProxyPort); err != nil {
		glog.Fatal(err)
	}
}

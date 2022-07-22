package server

import (
	"context"
	"flag"
	"net/http"
	"regexp"

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
	return http.ListenAndServe(fmt.Sprintf("%s:%d", *webProxyAddr, *webProxyPort), appendCors(mux))
}

func allowedOrigin(origin string) bool {
	if matched, _ := regexp.MatchString("https://teleapp.aifs.lunarcloud.org", origin); matched {
		return true
	}
	return false
}

func appendCors(muxHandler http.Handler) http.Handler {
	newHandler := http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
		if allowedOrigin(req.Header.Get("Origin")) {
			respWriter.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
			respWriter.Header().Set("Access-Control-Allow-Methods", "POST")
			respWriter.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if req.Method == "OPTIONS" {
			return
		}
		muxHandler.ServeHTTP(respWriter, req)
	})
	return newHandler
}

// The proxy server is used to listen to HTTP requests to the backend server.
// This function starts the proxy server.
func Proxy_main(mainServerAddr *string, mainServerPort *int, webProxyAddr *string, webProxyPort *int) {
	flag.Parse()
	defer glog.Flush()

	fmt.Println("Starting GW HTTP-Proxy server on", *webProxyPort)

	if err := run(mainServerAddr, mainServerPort, webProxyAddr, webProxyPort); err != nil {
		glog.Fatal(err)
	}
}

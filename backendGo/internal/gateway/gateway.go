package gateway

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"

	//"crypto/tls"
	"fmt"
	"io/fs"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	//"google.golang.org/grpc/credentials/insecure"
	//"membership/insecure"
	pbExample "nucifera_backend/protos/membership"
	"nucifera_backend/third_party"
	"github.com/rs/cors"
)

// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenAPIHandler() http.Handler {
	mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

// Run runs the gRPC-Gateway, dialling the provided address.
func Run(dialAddr string) error {
	// Adds gRPC internal logs. This is quite verbose, so adjust as desired!
	log := grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)
	grpclog.SetLoggerV2(log)

	gwmux := runtime.NewServeMux()
	err := pbExample.RegisterDataServiceHandlerFromEndpoint(context.Background(), gwmux, dialAddr, []grpc.DialOption{grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}
	
	port := os.Getenv("PORT")
	if port == "" {
		port =  os.Getenv("HTTP_PORT")[1:]
	}

	c := cors.New(cors.Options {
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET"},
		AllowCredentials: true,
		AllowedHeaders:[]string{"*"},
		Debug: true,
	})

	gatewayAddr := ":" + port
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			gwmux.ServeHTTP(w, r)
			return
		}
	})
	
	log.Info("Serving gRPC-Gateway and OpenAPI Documentation on http://", gatewayAddr)
	return fmt.Errorf("serving gRPC-Gateway server: %w", http.ListenAndServe(gatewayAddr, c.Handler(handler)))
}

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
	"p99system/internal/membership/handlers"
	"p99system/internal/membership/middleware"
	//"google.golang.org/grpc/credentials/insecure"
	//"membership/insecure"
	pbExample "p99system/protos/membership"
	"p99system/third_party"
	//"github.com/rs/cors"
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

	oa := getOpenAPIHandler()

	tt := http.NewServeMux()
	// tt.HandleFunc("/membership/authorize", handlers.HandleAuth)
	// chainCommonMiddleware(tt, "/membership/response", handlers.HandleResponse, middleware.NewPostFormValidator(true))
	// chainCommonMiddleware(tt, "/membership/token", handlers.HandleToken, middleware.NewPostFormValidator(false))
	tt.HandleFunc("/membership/authorize", handlers.HandleAuth)
	tt.HandleFunc("/membership/response", handlers.HandleResponse)
	tt.HandleFunc("/membership/token", handlers.HandleToken)
	tt.HandleFunc("/membership/signin/oauth", handlers.HandleOauthSignIn)
	tt.HandleFunc("/membership/signin/response", handlers.HandleOauthSignInResponse)
	tt.HandleFunc("/membership/oauth2/userinfo", handlers.HandleUserInfo)

	// c := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowCredentials: true,
	// 	// Enable Debugging for testing, consider disabling in production
	// 	Debug: true,
	// })
	
	port := os.Getenv("PORT")
	if port == "" {
		port =  os.Getenv("HTTP_PORT")[1:]
	}
	gatewayAddr := ":" + port
	gwServer := http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				gwmux.ServeHTTP(w, r)
				return
			}
			if strings.HasPrefix(r.URL.Path, "/membership") {
			 	tt.ServeHTTP(w, r)
			 	return
			}
			oa.ServeHTTP(w, r)
		}),
	}
	
	log.Info("Serving gRPC-Gateway and OpenAPI Documentation on http://", gatewayAddr)
	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())
}

func chainCommonMiddleware(mux *http.ServeMux, pattern string, handler http.HandlerFunc, extras ...middleware.Middleware) {
	middlewareSlice := []middleware.Middleware{middleware.NewNotFoundMiddleware(pattern)}
	middlewareSlice = append(middlewareSlice, extras...)
	chain := middleware.Chain(handler, middlewareSlice...)
	mux.HandleFunc(pattern, chain)
}
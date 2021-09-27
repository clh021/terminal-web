package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	grpcServer := grpc.NewServer()
	RegisterShellServiceServer(grpcServer, &ShellService{})

	port := 7900
	h := func(rw http.ResponseWriter, r *http.Request) {
		grpcweb.WrapServer(
			grpcServer,
			grpcweb.WithOriginFunc(func(origin string) bool { return true }),
			grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
			grpcweb.WithWebsockets(true),
		).ServeHTTP(rw, r)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(h),
	}
	grpclog.Printf("starting server. http port: %d ", port)
	log.Printf("starting server. http port: %d ", port)
	err := httpServer.ListenAndServe()
	if err != nil {
		grpclog.Fatalf(" failed starting http server: %v", err)
	}
}

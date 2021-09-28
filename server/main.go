package main

import (
	context "context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

func director(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
	// 得到 serviceName 和 methodName
	log.Println("fullMethodName:", fullMethodName)
	// 将 proxy 的请求根据 serviceName 转发给 真正的服务提供 者
	md, _ := metadata.FromIncomingContext(ctx)
	address := "grpc21:7901"
	conn, err := grpc.DialContext(ctx, address, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithCodec(proxy.Codec()))
	if err != nil {
		log.Println("error[conn]:", err)
		return nil, nil, err
	}
	// 再把提供者的响应转发给 请求者
	if conn != nil {
		return metadata.NewOutgoingContext(ctx, md), conn, nil
	}
	return nil, nil, fmt.Errorf("not support this proxy rpc message:%q", fullMethodName)
}

func main() {
	grpcServer := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)),
	)
	RegisterShellServiceServer(grpcServer, &ShellService{})

	port := 7900
	h := func(rw http.ResponseWriter, r *http.Request) {
		grpcweb.WrapServer(
			grpcServer,
			grpcweb.WithOriginFunc(func(origin string) bool { return true }),
			grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
			grpcweb.WithWebsockets(true),
			grpcweb.WithWebsocketPingInterval(time.Second*10),
			grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		).ServeHTTP(rw, r)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(h),
	}
	log.Printf("starting server. http port: %d ", port)
	err := httpServer.ListenAndServe()
	if err != nil {
		grpclog.Fatalf(" failed starting http server: %v", err)
	}
}

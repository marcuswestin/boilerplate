package main

import (
	"examples/conf"
	"examples/grpc-echo"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := conf.Get("GRPC_ECHO_SERVER_SERVICE_PORT", "9090")
	run(":" + port)
}

func run(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	echo.RegisterEchoServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Println("Listen:", address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// server is used to implement echo.EchoServer.
type server struct{}

// Ping implements echo.EchoServer
func (s *server) Ping(ctx context.Context, in *echo.PingReq) (*echo.PingRes, error) {
	log.Println("Req: SayHello -", in)
	return &echo.PingRes{Pong: "pong " + in.Ping}, nil
}

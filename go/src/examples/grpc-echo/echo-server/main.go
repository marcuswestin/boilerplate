package main

import (
	"examples/grpc-echo"
	"log"
	"net"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := env("GRPC_ECHO_SERVER_SERVICE_PORT", "9090")
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

// Greet implements echo.EchoServer
func (s *server) Greet(ctx context.Context, in *echo.GreetReq) (*echo.GreetRes, error) {
	log.Println("Req: SayHello -", in)
	return &echo.GreetRes{Greeting: "Yo, " + in.Name}, nil
}

func env(key, defaultValue string) (resultValue string) {
	defer func() { log.Println("config: " + key + "=" + resultValue) }()
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

package main

import (
	"context"
	"examples/conf"
	"examples/grpc-echo"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("missing argument: name")
		os.Exit(1)
	}

	echoHost := conf.Get("GRPC_ECHO_SERVER_SERVICE_HOST", "localhost")
	echoPort := conf.Get("GRPC_ECHO_SERVER_SERVICE_PORT", "9090")
	run(echoHost+":"+echoPort, os.Args[1])
}

func run(echoAddress, name string) {
	// Connect to grpc-echo server
	conn, err := grpc.Dial(echoAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := echo.NewEchoClient(conn)

	echoRes, err := c.Ping(context.Background(), &echo.PingReq{Ping: name})
	if err != nil {
		log.Fatalf("Echo service returned error: %v", err)
	}
	log.Printf("Response: %s", echoRes.Pong)
}

package main

import (
	"context"
	"examples/conf"
	"examples/grpc-echo"
	"log"
	"net/http"
	"strings"

	"google.golang.org/grpc"
)

func main() {
	port := conf.Get("GRPC_ECHO_HTTP_BRIDGE_SERVICE_PORT", "8080")
	echoHost := conf.Get("GRPC_ECHO_SERVER_SERVICE_HOST", "localhost")
	echoPort := conf.Get("GRPC_ECHO_SERVER_SERVICE_PORT", "9090")
	run(":"+port, echoHost+":"+echoPort)
}

func run(address, echoAddress string) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(echoAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := echo.NewEchoClient(conn)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Request:", req.Method, req.URL.String())
		name := strings.Replace(req.URL.Path, "/", "", 1)

		echoRes, err := c.Ping(context.Background(), &echo.PingReq{Ping: name})
		if err != nil {
			http.Error(w, "Echo service returned error: "+err.Error(), 500)
			return
		}

		w.Write([]byte(echoRes.Pong))
	})
	log.Println("Listen", address)
	err = http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("Error listening and serving: %v", err)
	}
}

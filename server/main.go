package main

import (
	"context"
	"fmt"
	token "github.com/hollo08/go-test/server/token"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	go GrpcServer()
	go GrpcClient()
	var a string
	fmt.Scan(&a)
}

// main start a gRPC server and waits for connection
func GrpcServer() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	//s := token.Server{}等效于new(token.Server)
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	//token.RegisterPingServer(grpcServer, &s)
	token.RegisterPingServer(grpcServer, new(token.Server))
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
func GrpcClient() {
	var conn *grpc.ClientConn
	//call Login
	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := token.NewPingClient(conn)
	loginReply, err := c.Login(context.Background(), &token.LoginRequest{Username: "gavin", Password: "gavin"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	fmt.Println("Login Reply:", loginReply.Status)
	//Call SayHello
	requestToken := new(token.AuthToekn)
	requestToken.Token = loginReply.Token
	conn, err = grpc.Dial(":7777", grpc.WithInsecure(), grpc.WithPerRPCCredentials(requestToken))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c = token.NewPingClient(conn)
	helloreply, err := c.SayHello(context.Background(), &token.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", helloreply.Greeting)
}
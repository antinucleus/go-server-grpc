package main

import (
	"context"
	"fmt"
	"log"
	"loginpb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	loginpb.UnimplementedLoginServiceServer
}

func (c *server) Login(ctx context.Context, req *loginpb.LoginRequest) (*loginpb.LoginResponse, error) {
	fmt.Printf("Login function was invoked with %v\n", req)
	username := req.GetUsername()
	password := req.GetPassword()

	var r string
	if username == "test" && password == 123 {
		r = "Successful"
	} else {
		r = "Error"
	}
	res := &loginpb.LoginResponse{
		Result: r,
	}
	return res, nil

}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	loginpb.RegisterLoginServiceServer(s, &server{})
	fmt.Println("Server is started")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}

}

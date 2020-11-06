package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/mjafari98/go-grpc-adder/adder"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedAdderServer
}

func (s *server) AddNumber(ctx context.Context, in *pb.NumberRequest) (*pb.NumberResponse, error) {
	log.Printf("Received: %d", in.GetNum())
	log.Printf("Returning: %d", in.GetNum()+17)
	log.Printf("-------------------")
	return &pb.NumberResponse{
		Num:    in.GetNum() + 17,
		Status: fmt.Sprintf("your number ( %d ) + 17 = %d", in.GetNum(), in.GetNum()+17),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAdderServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

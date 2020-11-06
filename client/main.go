package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/mjafari98/go-grpc-adder/adder"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAdderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var number int32
	log.Printf("Enter a Number. You have 10 seconds until the timeout...\n")
	_, _ = fmt.Scanf("%d", &number)

	r, err := c.AddNumber(ctx, &pb.NumberRequest{Num: number})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Number: %d", r.GetNum())
	log.Printf("Status: %s", r.GetStatus())
	log.Printf("-------------------")
}

package main

import (
	pb "helloworld/helloworld"
	"log"
	"os"
	"time"

	"context"
	"google.golang.org/grpc"
)
const (
	address = "localhost:8888"
	defaultName = "world"
)

func main(){
	conn,err := grpc.Dial(address,grpc.WithInsecure(),grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect: %v",err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1:]
	}
	ctx,cancel :=context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	r,err := c.SayHello(ctx,&pb.HelloRequest{Name:name})
	if err != nil {
		log.Fatalf("could not greet: %v",err)
	}
	log.Printf("Greeting: %s",r.GetMessage())
}
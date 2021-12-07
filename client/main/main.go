package main

import (
	"context"
	"google.golang.org/grpc"
	hello "grpc/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1314", grpc.WithInsecure())
	if nil != err {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := hello.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &hello.StringReq{Value: "client"})
	if nil != err {
		log.Fatalln(err)
	}

	log.Println(reply)
}

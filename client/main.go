package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/ddld93/sms-outboubd-microservices/gen/smspb"
	"google.golang.org/grpc"
)

func main() {
	//connect to server
	conn, err := grpc.Dial("localhost:3333", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting to grpc server %v\n", err)
	}
	//create new client from service
	client := smspb.NewSMSServiceApiClient(conn)
	req := &smspb.GetSMSRequest{
		Id: "",
	}
	//invoke
	
	stream, err := client.GetAll(context.Background(),req)
	if err != nil {
		log.Fatalf("Error reading server stream RPC  %v\n", err)
	}
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			log.Println("Server stream data finished!!!!!")
		}
		if err != nil {
			log.Fatalf("Error reading content of server stream %v\n", err)
		}
		fmt.Printf("Data Received >>> %+v\n", resp)
	}

}

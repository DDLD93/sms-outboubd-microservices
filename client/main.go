package main

import (
	"context"
	"fmt"
	


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
	// req := &smspb.GetAllSMSRequest{
	// 	Message: "",
	// }	
	one := &smspb.GetSMSRequest{
		Id: "61d8aa5a5112a97ad094b2d1",
	}
	resp,err := client.Send(context.TODO(),&smspb.SendSMSRequest{
		SenderPhone: "07063600093",
		ReceiverPhone: "08035736994",
		TextSubject: "abujere",
		TextBody: "to maryam jere",
	})

	fmt.Println(resp)

	// message, err := client.GetOne(context.Background(),req)
	//  fmt.Println(message)


	//invoke
	message,err := client.GetOne(context.TODO(),one)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(message)

	// resp,err := client.Send(context.TODO(),one)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(message)
	
	// stream, err := client.GetAll(context.Background(),req)
	// if err != nil {
	// 	log.Fatalf("Error reading server stream RPC  %v\n", err)
	// }
	// for {
	// 	resp, err := stream.Recv()
	// 	if err == io.EOF {
	// 		log.Println("Server stream data finished!!!!!")
	// 	}
	// 	if err != nil {
	// 		log.Fatalf("Error reading content of server stream %v\n", err)
	// 	}
	// 	fmt.Printf("Data Received >>> %+v\n", resp)
	// }

}

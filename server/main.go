package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/ddld93/sms-outboubd-microservices/gen/smspb"
	"github.com/ddld93/sms-outboubd-microservices/server/controller"
	"github.com/ddld93/sms-outboubd-microservices/server/model"
	"github.com/ddld93/sms-outboubd-microservices/server/routes"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type smsServiceApiServer struct {
	smspb.UnimplementedSMSServiceApiServer
}

func init() {

	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var dbPort, _ = strconv.Atoi(os.Getenv("DB_PORT"))
var dbInt = routes.SMSRoute{SMSCtrl: controller.NewSMSCtrl("localhost", dbPort)}

func (c *smsServiceApiServer) Send(ctx context.Context, req *smspb.SendSMSRequest) (*smspb.SendSMSResponse, error) {
	senderPhone := req.SenderPhone
	receiverPhone := req.ReceiverPhone
	textSubject := req.TextSubject
	textBody := req.TextBody
	text := model.SmsOutbound{SenderPhone: senderPhone, ReceiverPhone: receiverPhone, TextSubject: textSubject, TextBody: textBody}
	// send message by calling the send message route
	err := dbInt.SendSMS(text)
	if err != nil {
		return nil, err
	}

	log.Printf("Sending Message >> %+v\n", text)
	return &smspb.SendSMSResponse{Message: fmt.Sprintf("Message sent Successfully")}, nil
}

func (c *smsServiceApiServer) Get(ctx context.Context, req *smspb.GetSMSRequest) (*smspb.GetSMSResponse, error) {
	id := req.Id
	// send message by calling the send message route
	message, err := dbInt.GetSMS(id)
	if err != nil {
		return nil, err
	}
	return &smspb.GetSMSResponse{
		SenderPhone:   message.SenderPhone,
		ReceiverPhone: message.ReceiverPhone,
		TextSubject:   message.TextSubject,
		TextBody:      message.TextBody}, nil
}

// func (c *smsServiceApiServer) GetAll(req *smspb.GetSMSRequest, stream *smspb.)  error{
// 	id := req.Id
// 	// send message by calling the send message route

// }

func main() {

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	host := os.Getenv("HOST")

	url := fmt.Sprintf("%s:%d", host, port)

	// listner
	listener, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatalf("Error starting listener %v\n", err)
	}
	//grpc server
	grpcServer := grpc.NewServer()
	//register Service
	smspb.RegisterSMSServiceApiServer(grpcServer, &smsServiceApiServer{})
	//serve
	fmt.Println("Grpc server listening on port 5000!")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Error serving grpc services %v\n", err)
	}

}

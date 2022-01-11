package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	//"strconv"

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

var dbPort,_ = strconv.Atoi(os.Getenv("DBPORT"))
var dbInt = routes.SMSRoute{SMSCtrl: controller.NewSMSCtrl("localhost", 1515)}

func (c *smsServiceApiServer) Send(ctx context.Context, req *smspb.SendSMSRequest) (*smspb.SendSMSResponse, error) {

	text := model.SmsOutbound{
		SenderPhone:   req.SenderPhone,
		ReceiverPhone: req.ReceiverPhone,
		TextSubject:   req.TextSubject,
		TextBody:      req.TextBody,
	}

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
		TextBody:      message.TextBody,
	}, nil
}

func (c *smsServiceApiServer) GetAll(req *smspb.GetSMSRequest, stream smspb.SMSServiceApi_GetAllServer) error {

	// send message by calling the send message route
	messages, err := dbInt.GetAllSMS()
	if err != nil {
		return err
	}
	for _, msg := range messages {
		resp := &smspb.GetSMSResponse{
			SenderPhone:   msg.SenderPhone,
			ReceiverPhone: msg.ReceiverPhone,
			TextSubject:   msg.TextSubject,
			TextBody:      msg.TextBody,
		}
		stream.Send(resp)
	}

	return nil
}

func main() {
	

   
	

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	fmt.Println(port,"\n",host)

	url := fmt.Sprintf("%s:%s", "0.0.0.0", "3333")

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
	fmt.Printf("Grpc server listening on port %v!", port)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("Error serving grpc services %v\n", err)
	}

}

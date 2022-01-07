package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ddld93/sms-outboubd-microservices/controller"
	"github.com/joho/godotenv"

	// "github.com/ddld93/sms-outboubd-microservices/model"
	"github.com/ddld93/sms-outboubd-microservices/routes"
)

func init() {

    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }
}


func main()  {

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	smsCtrl := controller.NewSMSCtrl("localhost", port)
	route := routes.SMSRoute{SMSCtrl: smsCtrl}

	// newSMS := model.SmsOutbound{
	// 	SenderPhone: "0810520235",
	// 	ReceiverPhone: "0803344554",
	// 	TextSubject: "test 2",
	// 	TextBody: "testing body 2",
	// }

	// err := route.SendSMS(newSMS)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// message,err := route.GetSMS("guygiuhhjklj")
	messages,err :=	route.GetSMS("61d8aa5a5112a97ad094b2d1")

 	if err != nil {
 	fmt.Println(err)
	 }
	 fmt.Println(messages)
}
package main

import (
	"fmt"

	"github.com/ddld93/sms-outboubd-microservices/controller"
	// "github.com/ddld93/sms-outboubd-microservices/model"
	"github.com/ddld93/sms-outboubd-microservices/routes"
)


func main()  {
	smsCtrl := controller.NewSMSCtrl("localhost", 4321)
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
	messages,err :=	route.GetSMS("61d8aa5a5112a97ad094b2d")

 	if err != nil {
 	fmt.Println(err)
	 }
	 fmt.Println(messages)
}
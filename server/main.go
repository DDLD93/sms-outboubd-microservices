package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ddld93/sms-outboubd-microservices/server/controller"

	"github.com/joho/godotenv"

	// "github.com/ddld93/sms-outboubd-microservices/server/model"
	"github.com/ddld93/sms-outboubd-microservices/server/routes"
)

func init() {

	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}



func main()  {

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	smsCtrl := controller.NewSMSCtrl("localhost", port)
	route := routes.SMSRoute{SMSCtrl: smsCtrl}
	messsages ,err :=route.GetAllSMS()
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(messsages)

	messsage ,err2 :=route.GetSMS("61d8aa5a5112a97ad094b2d1")
	if err2 != nil{
		fmt.Sprint(err)
	}
	fmt.Sprint(messsage)

	
}
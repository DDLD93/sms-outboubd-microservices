package routes

import (
	"github.com/ddld93/sms-outboubd-microservices/controller"
	"github.com/ddld93/sms-outboubd-microservices/model"
)

type SMSRoute struct{
	SMSCtrl *controller.SMSCtrl
}

func (s SMSRoute) SendSMS(sms model.SmsOutbound) error{
	err := s.SMSCtrl.Send(&sms)
	if err != nil{
		return err
	}
	return nil
}
func (s *SMSRoute) GetSMS(id string)  (model.SmsOutbound,error){
	sms,err := s.SMSCtrl.GetOne(id)
	if err != nil{
		return sms,err
	}
	return sms,nil
}
func (s *SMSRoute) GetAllSMS()  ([]model.SmsOutbound,error){
	sms,err :=  s.SMSCtrl.GetAll()
	if err != nil{
		return sms,err
	}
	return sms,nil
}
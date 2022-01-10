package controller

import (
	"errors"
	"fmt"
	"log"

	"github.com/ddld93/sms-outboubd-microservices/server/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SMSCtrl struct {
	Session *mgo.Session
}

func NewSMSCtrl(host string, port int) *SMSCtrl {
	url := fmt.Sprintf("mongodb://%s:%d", host, port)
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal("Error connecting to mongo database", err)
	}
	return &SMSCtrl{Session: session}
}
func (s *SMSCtrl) Save(sms *model.SmsOutbound) error {
	sms.Id = bson.NewObjectId()
	err := s.Session.DB("smsdb").C("outbound-sms").Insert(sms)
	if err != nil {
		return errors.New("failed to insert message to database")
	}
	return nil
}

func (s *SMSCtrl) Send(sms *model.SmsOutbound) error {
	// make a calls sms api
	//if  no error
	err := s.Save(sms)
	if err != nil {
		return err
	}
	fmt.Println("message sent successifully")
	return nil
}

func (s *SMSCtrl) GetAll() ([]model.SmsOutbound, error) {
	smsSlice := []model.SmsOutbound{}
	err := s.Session.DB("smsdb").C("outbound-sms").Find(bson.M{}).All(&smsSlice)
	if err != nil {
		return smsSlice, errors.New("error getting messages to database")
	}
	return smsSlice, nil
}
func (s *SMSCtrl) GetOne(Id string) (model.SmsOutbound, error) {
	sms := model.SmsOutbound{}
	id := bson.ObjectIdHex(Id)
	err := s.Session.DB("smsdb").C("outbound-sms").FindId(id).One(&sms)
	if err != nil {
		return sms, errors.New("error getting messages to database")
	}
	return sms, nil
}

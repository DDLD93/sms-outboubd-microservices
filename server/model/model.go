package model

import "gopkg.in/mgo.v2/bson"


type SmsOutbound struct {
	Id    bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	SenderPhone string        `json:"senderPhone" bson:"senderPhone"`
	ReceiverPhone string        `json:"receiverPhone" bson:"receiverPhone"`
	TextSubject string        `json:"textSubject" bson:"textSubject"`
	TextBody string        `json:"textBody" bson:"textBody"`
	Date bson.MongoTimestamp
}
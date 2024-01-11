package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Catatan struct {
	ID           primitive.ObjectID 	   `bson:"_id,omitempty" json:"_id,omitempty"`
	Title         		string             `bson:"title,omitempty" json:"title,omitempty"`
	Note 				string             `bson:"note,omitempty" json:"note,omitempty"`
	Date        		string             `bson:"date,omitempty" json:"date,omitempty"`
	StartTime 			string             `bson:"starttime,omitempty" json:"starttime,omitempty"`
	EndTime 			string             `bson:"endtime,omitempty" json:"endtime,omitempty"`
	Remind        		string             `bson:"remind,omitempty" json:"remind,omitempty"`
	Repeat        		string             `bson:"repeat,omitempty" json:"repeat,omitempty"`
}

type User struct{
	ID              primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
	FullName	    string          	`bson:"fullname,omitempty" json:"fullname,omitempty"`
	Email 			string             	`bson:"email,omitempty" json:"email,omitempty"`
	Password        string          	`bson:"password,omitempty" json:"password,omitempty"`
	PhoneNumber     string          	`bson:"phonenumber,omitempty" json:"phonenumber,omitempty"`
	Salt        	string              `bson:"salt,omitempty" json:"salt,omitempty"`
	
}
type Credential struct {
	Status  int    `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Role    string `json:"role,omitempty" bson:"role,omitempty"`
}

type Response struct {
	Status  int    `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Payload struct {
	Id    primitive.ObjectID `json:"id"`
	Email string             `json:"email"`
	Exp   time.Time          `json:"exp"`
	Iat   time.Time          `json:"iat"`
	Nbf   time.Time          `json:"nbf"`
}
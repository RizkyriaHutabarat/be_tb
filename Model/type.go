package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Catatan struct {
	ID           primitive.ObjectID 	   `bson:"_id,omitempty" json:"_id,omitempty"`
	Judul_Tugas         string             `bson:"judul_tugas,omitempty" json:"judul_tugas,omitempty"`
	Matkul          	string             `bson:"matkul,omitempty" json:"matkul,omitempty"`
	Deskripsi_Tugas 	string             `bson:"deskripsi_tugas,omitempty" json:"deskripsi_tugas,omitempty"`
	Tanggal_Deadline 	string             `bson:"tanggal_deadline,omitempty" json:"tanggal_deadline,omitempty"`
	Tanggal_Submit 		string             `bson:"tanggal_submit,omitempty" json:"tanggal_submit,omitempty"`
}

type User struct{
	ID              primitive.ObjectID 	`bson:"_id,omitempty" json:"_id,omitempty"`
	Username 		string             	`bson:"username,omitempty" json:"username,omitempty"`
	Password        string          	`bson:"password,omitempty" json:"password,omitempty"`
}

type Token struct{
	Token_String              string          	`bson:"tokenstring,omitempty" json:"tokenstring,omitempty"`
}

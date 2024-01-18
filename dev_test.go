package Catatan

import (
	"fmt"
	"testing"

	model "github.com/RizkyriaHutabarat/be_tb/Model"
	module "github.com/RizkyriaHutabarat/be_tb/Module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = module.MongoConnect("MONGOSTRING", "db_note")
var collectionnameUser = "user"

// var collectionnameFishingspot = "fishingspot"

func TestGenerateKey(t *testing.T) {
	privateKey, publicKey := module.GenerateKey()
	fmt.Println("privateKey : ", privateKey)
	fmt.Println("publicKey : ", publicKey)
}

func TestSignUp(t *testing.T) {
	conn := db
	var user model.User
	user.FullName = "lina"
	user.Email = "lina@gmail.com"
	user.Password = "lina1234"
	user.PhoneNumber = "627483918273"
	email, err := module.SignUp(conn, collectionnameUser, user)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil SignUp : ", email)
	}
}

func TestLogInn(t *testing.T) {
	conn := db
	var user model.User
	user.Email = "lina@gmail.com"
	user.Password = "lina1234"
	user, _ = module.LogIn(conn, collectionnameUser, user)
	tokenstring, err := module.Encode(user.ID, user.Email, "33186fcfc13ba9946bf200cf6c7808e6ebfc605140f65809e06648985b08ebda2df976efd75eacf2a37b1ce184deec8d3b72cb78f7881ed5e7a02d97351c2aef")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil LogIn : ", user.FullName)
		fmt.Print("Berhasil LogIn : " + tokenstring)
	}
}

func TestToken(*testing.T) {
	token := "v4.public.eyJleHAiOiIyMDI0LTAxLTA0VDExOjI1OjU0WiIsImZ1bGxuYW1lIjoiYWRtaW5AZ21haWwuY29tIiwiaWF0IjoiMjAyNC0wMS0wNFQwOToyNTo1NFoiLCJpZCI6IjY1OTY1ZWNkY2MxOGQxNmNkNGNhNGY4YSIsIm5iZiI6IjIwMjQtMDEtMDRUMDk6MjU6NTRaIn22kA21UMcQv-Q3SGvd2AdY6B1UMk13v97NlVu2HDGYnLaO5erzeLLET7R47uqk0klAWctireNQDVGAaIeRNjf4F"
	tokenstring, err := module.Decode("2df976efd75eacf2a37b1ce184deec8d3b72cb78f7881ed5e7a02d97351c2aef", token)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("Id Token : " + tokenstring.Id.Hex())
		fmt.Print("Email Token : " + tokenstring.Email)
	}
}


func TestDeleteCatatan(t *testing.T) {
	conn := db
	id := "659fcfff7775a76bc7286670"
	objectId, err := primitive.ObjectIDFromHex(id)
	err = module.DeleteCatatan(objectId, "catatan", conn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil Delete Catatan")
	}
}

func TestInsertCatatan(t *testing.T) {
	id, err := module.InsertOneDoc(db, "catatan", model.Catatan{
		ID:       primitive.NewObjectID(),
		Title : "Pemrograman IV",
		Note : "Membuat  Flutter",
		Date : "18/1/2024",
		StartTime : "19:55 PM",
		EndTime : "20:20 PM",
		Remind : "0 minutes early",
		Repeat : "none",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil InsertCatatan : ", id)
	}
}

func TestUpdateCatatan(t *testing.T) {
	id := "659fce28399bcb212f1f497e"
	objectId, err := primitive.ObjectIDFromHex(id)

	data := module.UpdateOneDoc(objectId, db, "catatan", model.Catatan{
		ID:       objectId,
		Title : "Pemrograman IV",
		Note : "Membuat Tugas Besar",
		Date : "11/1/2024",
		StartTime : "09:50 PM",
		EndTime : "10:12 PM",
		Remind : "0 minutes early",
		Repeat : "none",
		
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil UpdateCatatan", data)
	}
}
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
	user.FullName = "kia"
	user.Email = "kia@gmail.com"
	user.Password = "kia181103"
	user.PhoneNumber = "621219882869"
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
	user.Email = "admin@gmail.com"
	user.Password = "admin12345678"
	user, _ = module.LogIn(conn, collectionnameUser, user)
	tokenstring, err := module.Encode(user.ID, user.Email, "4280fd42157198546b8bff929bddeaaf2cc6c3cedea412fa307904814fd45d2fbf735cfe7d545c9e53360f4271673dd7ad15d0fd56039298908360591b86d6b5")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Berhasil LogIn : ", user.FullName)
		fmt.Print("Berhasil LogIn : " + tokenstring)
	}
}

func TestToken(*testing.T) {
	token := "v4.public.eyJlbWFpbCI6IiIsImV4cCI6IjIwMjQtMDEtMTFUMjA6MTY6NTcrMDc6MDAiLCJpYXQiOiIyMDI0LTAxLTExVDE4OjE2OjU3KzA3OjAwIiwiaWQiOiIwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAiLCJuYmYiOiIyMDI0LTAxLTExVDE4OjE2OjU3KzA3OjAwIn3Mp6gy60hsgsHUlMnuz6m_nMOiES-fBdM2C0JC3NYqZsGF2MED_pBAllwPvHHTa1g2owAC44qHBXrqRwAg3gYD"
	tokenstring, err := module.Decode("bf735cfe7d545c9e53360f4271673dd7ad15d0fd56039298908360591b86d6b5", token)
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
		Note : "Membuat Tugas Besar Flutter",
		Date : "11/1/2024",
		StartTime : "09:55 PM",
		EndTime : "10:20 PM",
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
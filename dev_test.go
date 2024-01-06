package Catatan

import (
	"context"
	"fmt"
	"testing"

	model "github.com/RizkyriaHutabarat/be_tb/Model"
	module "github.com/RizkyriaHutabarat/be_tb/Module"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAll(t *testing.T) {
	data := module.GetAllCatatan(module.MongoConn, "catatan")
	fmt.Println(data)
}

func TestGetCatatanFromID(t *testing.T) {
	id := "6596d0cc9d8c2918f00f9663"
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}
	ctt, err := module.GetCatatanFromID(objectID, module.MongoConn, "catatan")
	if err != nil {
		t.Fatalf("error calling GetCatatanFromID: %v", err)
	}
	fmt.Println(ctt)
}

func TestInsertCatatan(t *testing.T) {

	judul_tugas := "Tugas Pemrog 4"
	matkul := "Pemrograman"
	deskripsi_tugas := "Membuat Tampilan UI dengan menggunakan flutter"
	tanggal_deadline := "Sabtu 12 Januari 2024 Jam 14.00"
	tanggal_submit := "Sabtu 12 Januari 2024"
	
	insertedID, err := module.InsertCatatan(module.MongoConn, "catatan", judul_tugas, matkul, deskripsi_tugas, tanggal_deadline, tanggal_submit)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestDeleteCatatanByID(t *testing.T) {
	id := "65980cc866bb7aa8379c47e7" // ID data yang ingin dihapus
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to ObjectID: %v", err)
	}

	err = module.DeleteCatatanByID(objectID, module.MongoConn, "catatan")
	if err != nil {
		t.Fatalf("error calling DeleteCatatanByID: %v", err)
	}

	// Verifikasi bahwa data telah dihapus dengan melakukan pengecekan menggunakan GetPresensiFromID
	_, err = module.GetCatatanFromID(objectID, module.MongoConn, "catatan")
	if err == nil {
		t.Fatalf("expected data to be deleted, but it still exists")
	}
}

func TestUpdateCatatan(t *testing.T) {
    col := "catatan"

    // Define a test document
    doc := model.Catatan{
        ID:               primitive.NewObjectID(),
        Judul_Tugas:      "Tugas Pemrog 4",
        Matkul:           "Pemrograman",
        Deskripsi_Tugas:  "Membuat Tampilan UI dengan menggunakan flutter",
        Tanggal_Deadline: "Sabtu 12 Januari 2024 Jam 14.00",
        Tanggal_Submit:   "Sabtu 12 Januari 2024",
    }

    // Insert the test document into the collection
    if _, err := module.MongoConn.Collection(col).InsertOne(context.Background(), doc); err != nil {
        t.Fatalf("Failed to insert test document: %v", err)
    }

    // Define updated fields
    judul_tugas := "Tugas Pemrograman 4"
    matkul := "Pemrograman"
    deskripsi_tugas := "Tugas baru dengan flutter UI"
    tanggal_deadline := "Minggu 13 Januari 2024 Jam 15.00"
    tanggal_submit := "Minggu 13 Januari 2024"

    // Call UpdateCatatan with the test document ID and updated fields
    if err := module.UpdateCatatan(module.MongoConn, col, doc.ID, judul_tugas, matkul, deskripsi_tugas, tanggal_deadline, tanggal_submit); err != nil {
        t.Fatalf("UpdateCatatan failed: %v", err)
    }

    // Retrieve the updated document from the collection
    var updatedDoc model.Catatan
    if err := module.MongoConn.Collection(col).FindOne(context.Background(), bson.M{"_id": doc.ID}).Decode(&updatedDoc); err != nil {
        t.Fatalf("Failed to retrieve updated document: %v", err)
    }
}



//TB
//login
// func TestLoginUser(t *testing.T) {
// 	username := "admin"
// 	password := "admin"

// 	authenticated, err := module.LoginUser(module.MongoConn, "user", username, password)
// 	if err != nil {
// 		t.Errorf("Error authenticating user: %v", err)
// 	}

// 	if authenticated {
// 		fmt.Println("user authenticated successfully")
// 	} else {
// 		t.Errorf("user authentication failed")
// 	}
// }
func TestCreateUser(t *testing.T) {
	username := "yuhuu"
	password := "ahayy"

	insertedID, err := module.CreateUser(module.MongoConn, "user", username, password)
	if err != nil {
		t.Errorf("Error creating user: %v", err)
	}

	if insertedID.IsZero() {
		t.Fatal("Invalid inserted user ID")
	}
}

func TestLogin_Success(t *testing.T) {
	username := "username"
	password := "password"

	loggedIn, _, err := module.Login(username, password, module.MongoConn, "user")
	if err != nil {
		t.Errorf("Error logging in: %v", err)
	}

	if !loggedIn {
		t.Error("Login should be successful but it failed")
	}
}

func TestLogin_Failure(t *testing.T) {
	username := "username"
	password := "wrongpassword"

	loggedIn, _, err := module.Login(username, password, module.MongoConn, "user")
	if err != nil {
		t.Errorf("Error logging in: %v", err)
	}

	if loggedIn {
		t.Error("Login should fail but it succeeded")
	}
}



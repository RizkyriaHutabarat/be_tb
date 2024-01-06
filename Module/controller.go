package module

import (
	"context"
	"errors"
	"fmt"
	"os"

	model "github.com/RizkyriaHutabarat/be_tb/Model"
	"github.com/aiteung/atdb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "db_note",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

// func InsertCatatan(db *mongo.Database, col string, judul_tugas string, matkul string, deskripsi_tugas string, tanggal_deadline string, tanggal_submit string) (InsertedID interface{}) {
// 	var catatan model.Catatan
// 	catatan.Judul_Tugas = judul_tugas
// 	catatan.Matkul = matkul
// 	catatan.Deskripsi_Tugas = deskripsi_tugas
// 	catatan.Tanggal_Deadline = tanggal_deadline
// 	catatan.Tanggal_Submit = tanggal_submit
// 	return InsertOneDoc(db, col, catatan)
// }


func GetAllCatatan(db *mongo.Database, col string) (data []model.Catatan) {
	catatan := db.Collection(col)
	filter := bson.M{}
	cursor, err := catatan.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetCatatanFromID(_id primitive.ObjectID, db *mongo.Database, col string) (cat model.Catatan, errs error) {
	catatan := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := catatan.FindOne(context.TODO(), filter).Decode(&cat)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return cat, fmt.Errorf("no data found for ID %s", _id)
		}
		return cat, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return cat, nil
}


func InsertCatatan(db *mongo.Database, col string, judul_tugas string, matkul string, deskripsi_tugas string, tanggal_deadline string, tanggal_submit string) (insertedID primitive.ObjectID, err error) {
	catatan := bson.M{
		"judul_tugas":	judul_tugas,
		"matkul":     	matkul,
		"deskripsi_tugas":     	deskripsi_tugas,
		"tanggal_deadline":   	tanggal_deadline,
		"tanggal_submit": tanggal_submit,
		
	}
	result, err := db.Collection(col).InsertOne(context.Background(), catatan)
	if err != nil {
		fmt.Printf("InsertCatatan: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func DeleteCatatanByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
	catatan := db.Collection(col)
	filter := bson.M{"_id": _id}

	result, err := catatan.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

func UpdateCatatan(db *mongo.Database, col string,id primitive.ObjectID, judul_tugas string, matkul string, deskripsi_tugas string, tanggal_deadline string, tanggal_submit string) (err error) {
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{
		"judul_tugas":	judul_tugas,
		"matkul":     	matkul,
		"deskripsi_tugas":     	deskripsi_tugas,
		"tanggal_deadline":   	tanggal_deadline,
		"tanggal_submit": tanggal_submit,
		},
	}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Printf("UpdateCatatan: %v\n", err)
		return
	}
	if result.ModifiedCount == 0 {
		err = errors.New("No data has been changed with the specified ID")
		return
	}
	return nil
}

//TB

//login
// func LoginUser(db *mongo.Database, col string, username string, password string) (authenticated bool, err error) {
// 	filter := bson.M{
// 		"username": username,
// 		"password": password,
// 	}

// 	result, err := db.Collection(col).CountDocuments(context.Background(), filter)
// 	if err != nil {
// 		fmt.Printf("LoginUser: %v\n", err)
// 		return false, err
// 	}

// 	if result == 1 {
// 		return true, nil
// 	}

// 	return false, nil
// }

// func InsertUser(db *mongo.Database, col string, username string, password string) (insertedID primitive.ObjectID, err error) {
// 		admin := bson.M{
// 			"username":	username,
// 			"password": password,
// 		}
// 		result, err := db.Collection(col).InsertOne(context.Background(), admin)
// 		if err != nil {
// 			fmt.Printf("InsertAdmin: %v\n", err)
// 			return
// 		}
// 		insertedID = result.InsertedID.(primitive.ObjectID)
// 		return insertedID, nil
// }
	
// 	func GetAdmin(db *mongo.Database, col string) (data []model.Admin) {
// 		admin := db.Collection(col)
// 		filter := bson.M{}
// 		cursor, err := admin.Find(context.TODO(), filter)
// 		if err != nil {
// 			fmt.Println("GetALLData :", err)
// 		}
// 		err = cursor.All(context.TODO(), &data)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		return
// 	}
	
// 	func GetAdminFromID(_id primitive.ObjectID, db *mongo.Database, col string) (adm model.Admin, errs error) {
// 		admin := db.Collection(col)
// 		filter := bson.M{"_id": _id}
// 		err := admin.FindOne(context.TODO(), filter).Decode(&adm)
// 		if err != nil {
// 			if errors.Is(err, mongo.ErrNoDocuments) {
// 				return adm, fmt.Errorf("no data found for ID %s", _id)
// 			}
// 			return adm, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
// 		}
// 		return adm, nil
// 	}
// login

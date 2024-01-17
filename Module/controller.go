package module

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	model "github.com/RizkyriaHutabarat/be_tb/Model"
	"github.com/badoux/checkmail"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/argon2"
)

// var MongoString string = os.Getenv("MONGOSTRING")

// var MongoInfo = atdb.DBInfo{
// 	DBString: MongoString,
// 	DBName:   "db_note",
// }

// var MongoConn = atdb.MongoConnect(MongoInfo)

// mongodb
func MongoConnect(MongoString, dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MongoString)))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

//crud

func GetAllDocs(db *mongo.Database, col string, docs interface{}) interface{} {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error GetAllDocs %s: %s", col, err)
	}
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		return err
	}
	return docs
}

func InsertOneDoc(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		fmt.Println("error insert : ", err)
		return insertedID, fmt.Errorf("kesalahan server : insert")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}


func UpdateOneDoc(id primitive.ObjectID, db *mongo.Database, col string, doc interface{}) (err error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		return fmt.Errorf("error update: %v", err)
	}
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("tidak ada data yang diubah")
		return
	}
	return nil
}

func DeleteOneDoc(_id primitive.ObjectID, db *mongo.Database, col string) error {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// func GetAllCatatan(db *mongo.Database, col string) (data []model.Catatan) {
// 	catatan := db.Collection(col)
// 	filter := bson.M{}
// 	cursor, err := catatan.Find(context.TODO(), filter)
// 	if err != nil {
// 		fmt.Println("GetALLData :", err)
// 	}
// 	err = cursor.All(context.TODO(), &data)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	return
// }

func GetAllCatatan(db *mongo.Database, col string) (docs []model.Catatan, err error) {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return docs, fmt.Errorf("kesalahan server")
	}
	err = cursor.All(context.Background(), &docs)
	if err != nil {
		return docs, fmt.Errorf("kesalahan server")
	}
	return docs, nil
}

// func GetCatatanFromID(_id primitive.ObjectID, db *mongo.Database, col string) (cat model.Catatan, errs error) {
// 	catatan := db.Collection(col)
// 	filter := bson.M{"_id": _id}
// 	err := catatan.FindOne(context.TODO(), filter).Decode(&cat)
// 	if err != nil {
// 		if errors.Is(err, mongo.ErrNoDocuments) {
// 			return cat, fmt.Errorf("no data found for ID %s", _id)
// 		}
// 		return cat, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
// 	}
// 	return cat, nil
// }

func GetCatatanById(db *mongo.Database, col string, idparam primitive.ObjectID) (doc model.Catatan, err error) {
	collection := db.Collection(col)
	filter := bson.M{"_id": idparam}
	err = collection.FindOne(context.Background(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return doc, fmt.Errorf("data tidak ditemukan untuk ID %s", idparam)
		}
		return doc, fmt.Errorf("kesalahan server")
	}
	return doc, nil
}

// func InsertCatatan(db *mongo.Database, col string, judul_tugas string, matkul string, deskripsi_tugas string, tanggal_deadline string, tanggal_submit string) (insertedID primitive.ObjectID, err error) {
// 	catatan := bson.M{
// 		"judul_tugas":	judul_tugas,
// 		"matkul":     	matkul,
// 		"deskripsi_tugas":     	deskripsi_tugas,
// 		"tanggal_deadline":   	tanggal_deadline,
// 		"tanggal_submit": tanggal_submit,
		
// 	}
// 	result, err := db.Collection(col).InsertOne(context.Background(), catatan)
// 	if err != nil {
// 		fmt.Printf("InsertCatatan: %v\n", err)
// 		return
// 	}
// 	insertedID = result.InsertedID.(primitive.ObjectID)
// 	return insertedID, nil
// }

func InsertCatatan(db *mongo.Database, col string, r *http.Request) (bson.M, error) {
	title := r.FormValue("title")
	note := r.FormValue("note")
	date := r.FormValue("date")
	starttime := r.FormValue("starttime")
	endtime := r.FormValue("endtime")
	remind := r.FormValue("remind")
	repeat := r.FormValue("repeat")

	if title == "" || note == "" || date == "" || starttime == "" || endtime == ""|| remind == "" || repeat == ""  {
		return bson.M{}, fmt.Errorf("mohon untuk melengkapi data")
	}
	
	catatan := bson.M{
		"_id":      primitive.NewObjectID(),
		"title":    	   title,
		"note":     note,
		"date": 			   date,
		"starttime":   	   starttime,
		"endtime":       endtime,
		"remind":              remind,
		"repeat":              repeat,
	}
	_, err := InsertOneDoc(db, col, catatan)
	if err != nil {
		return bson.M{}, err
	}
	return catatan, nil
}

// func DeleteCatatanByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
// 	catatan := db.Collection(col)
// 	filter := bson.M{"_id": _id}

// 	result, err := catatan.DeleteOne(context.TODO(), filter)
// 	if err != nil {
// 		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
// 	}

// 	if result.DeletedCount == 0 {
// 		return fmt.Errorf("data with ID %s not found", _id)
// 	}

// 	return nil
// }
func DeleteCatatan(_id primitive.ObjectID, col string, db *mongo.Database) error {
	err := DeleteOneDoc(_id, db, col)
	if err != nil {
		return err
	}
	return nil
}

// func UpdateCatatan(db *mongo.Database, col string,id primitive.ObjectID, judul_tugas string, matkul string, deskripsi_tugas string, tanggal_deadline string, tanggal_submit string) (err error) {
// 	filter := bson.M{"_id": id}
// 	update := bson.M{
// 		"$set": bson.M{
// 		"judul_tugas":	judul_tugas,
// 		"matkul":     	matkul,
// 		"deskripsi_tugas":     	deskripsi_tugas,
// 		"tanggal_deadline":   	tanggal_deadline,
// 		"tanggal_submit": tanggal_submit,
// 		},
// 	}
// 	result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
// 	if err != nil {
// 		fmt.Printf("UpdateCatatan: %v\n", err)
// 		return
// 	}
// 	if result.ModifiedCount == 0 {
// 		err = errors.New("No data has been changed with the specified ID")
// 		return
// 	}
// 	return nil
// }

func UpdateCatatan(_id primitive.ObjectID, db *mongo.Database, col string, r *http.Request) (bson.M, error) {
	judul_tugas := r.FormValue("judul_tugas")
	deskripsi_tugas := r.FormValue("deskripsi_tugas")
	date := r.FormValue("date")
	starttime := r.FormValue("starttime")
	endtime := r.FormValue("endtime")
	remind := r.FormValue("remind")
	repeat := r.FormValue("repeat")

	if judul_tugas == "" || deskripsi_tugas == "" || date == "" || starttime == "" || endtime == ""|| remind == "" || repeat == ""  {
		return bson.M{}, fmt.Errorf("mohon untuk melengkapi data")
	}

	catatan := bson.M{
		"_id":      primitive.NewObjectID(),
		"judul_tugas":    	   judul_tugas,
		"deskripsi_tugas":     deskripsi_tugas,
		"date": 			   date,
		"starttime":   	   starttime,
		"endtime":       endtime,
		"remind":              remind,
		"repeat":              repeat,
	}
	err := UpdateOneDoc(_id, db, col, catatan)
	if err != nil {
		return bson.M{}, err
	}
	return catatan, nil
}

// get user login
func GetUserLogin(PASETOPUBLICKEYENV string, r *http.Request) (model.Payload, error) {
	tokenstring := r.Header.Get("Authorization")
	payload, err := Decode(os.Getenv(PASETOPUBLICKEYENV), tokenstring)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

// get id
func GetID(r *http.Request) string {
	return r.URL.Query().Get("id")
}

// return struct
func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}


//user
func GetUserFromID(_id primitive.ObjectID, db *mongo.Database) (doc model.User, err error) {
	collection := db.Collection("user")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return doc, fmt.Errorf("no data found for ID %s", _id)
		}
		return doc, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return doc, nil
}

func GetUserFromEmail(email string, db *mongo.Database) (doc model.User, err error) {
	collection := db.Collection("user")
	filter := bson.M{"email": email}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return doc, fmt.Errorf("email tidak ditemukan")
		}
		return doc, fmt.Errorf("kesalahan server")
	}
	return doc, nil
}

func GetUserFromPhonenumber(phonenumber string, db *mongo.Database) (doc model.User, err error) {
	collection := db.Collection("user")
	filter := bson.M{"no_hp": phonenumber}
	err = collection.FindOne(context.Background(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return doc, fmt.Errorf("nomor telepon tidak ditemukan")
		}
		return doc, fmt.Errorf("kesalahan server")
	}
	return doc, nil
}

func ValidatePhoneNumber(phoneNumber string) (bool, error) {
	// Define the regular expression pattern for numeric characters
	numericPattern := `^[0-9]+$`

	// Compile the numeric pattern
	numericRegexp, err := regexp.Compile(numericPattern)
	if err != nil {
		return false, err
	}
	// Check if the phone number consists only of numeric characters
	if !numericRegexp.MatchString(phoneNumber) {
		return false, nil
	}

	// Define the regular expression pattern for "62" followed by 6 to 12 digits
	pattern := `^62\d{6,13}$`

	// Compile the regular expression
	regexpPattern, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	
	// Test if the phone number matches the pattern
	isValid := regexpPattern.MatchString(phoneNumber)

	return isValid, nil
}

func SignUp(db *mongo.Database, col string, insertedDoc model.User) (string, error) {
	if insertedDoc.FullName == "" || insertedDoc.Email == "" || insertedDoc.Password == "" || insertedDoc.PhoneNumber == ""{
		return "", fmt.Errorf("mohon untuk melengkapi data")
	}
	if err := checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return "", fmt.Errorf("email tidak valid")
	}
	userExists, _ := GetUserFromEmail(insertedDoc.Email, db)
	if insertedDoc.Email == userExists.Email {
		return "", fmt.Errorf("email sudah terdaftar")
	}
	validatePhoneNumber, _ := ValidatePhoneNumber(insertedDoc.PhoneNumber)
	if !validatePhoneNumber {
		return "", fmt.Errorf("nomor telepon tidak valid")
	}
	PhoneNumberExists, _ := GetUserFromPhonenumber(insertedDoc.PhoneNumber, db)
	if insertedDoc.PhoneNumber == PhoneNumberExists.PhoneNumber {
		return "", fmt.Errorf("nomor telepon sudah terdaftar")
	}
	if strings.Contains(insertedDoc.Password, " ") {
		return "", fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Password) < 8 {
		return "", fmt.Errorf("password terlalu pendek")
	}
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", fmt.Errorf("kesalahan server : salt")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	user := bson.M{
		"FullName": insertedDoc.FullName,
		"Email": insertedDoc.Email,
		"Password": hex.EncodeToString(hashedPassword),
		"Phonenumber": insertedDoc.PhoneNumber,
		"Salt": hex.EncodeToString(salt),
	}
	_, err = InsertOneDoc(db, col, user)
	if err != nil {
		return "", err
	}
	return insertedDoc.Email, nil
}

func LogIn(db *mongo.Database, col string, insertedDoc model.User) (user model.User, err error) {
	if insertedDoc.Email == "" || insertedDoc.Password == "" {
		return user, fmt.Errorf("mohon untuk melengkapi data")
	}
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return user, fmt.Errorf("email tidak valid")
	}
	existsDoc, err := GetUserFromEmail(insertedDoc.Email, db)
	if err != nil {
		return
	}
	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		return user, fmt.Errorf("kesalahan server : salt")
	}
	hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != existsDoc.Password {
		return user, fmt.Errorf("password salah")
	}
	return existsDoc, nil
}

//TB





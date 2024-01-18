package module

import (
	"encoding/json"
	"net/http"
	"os"

	model "github.com/RizkyriaHutabarat/be_tb/Model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	credential model.Credential
	response   model.Response
	user       model.User
)

func SignUpHandler(MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	email, err := SignUp(conn, collectionname, user)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Berhasil SignUp"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data": bson.M{
			"email": email,
		},
	}
	return GCFReturnStruct(responData)
}

func LogInHandler(PASETOPRIVATEKEYENV, MONGOCONNSTRINGENV, dbname, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(response)
	}
	user, err := LogIn(conn, collectionname, user)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	tokenstring, err := Encode(user.ID, user.Email, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		response.Message = "Gagal Encode Token : " + err.Error()
		return GCFReturnStruct(response)
	}
	//
	credential.Message = "Selamat Datang " + user.FullName
	credential.Token = tokenstring
	credential.Status = 200
	responData := bson.M{
		"status":  credential.Status,
		"message": credential.Message,
		"data": bson.M{
			"token": credential.Token,
			"email": user.Email,
		},
	}
	return GCFReturnStruct(responData)
}


// func EditEmailHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
// 	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
// 	response.Status = 400
// 	//
// 	user_login, err := GetUserLogin(PASETOPUBLICKEYENV, r)
// 	if err != nil {
// 		response.Message = "Gagal Decode Token : " + err.Error()
// 		return GCFReturnStruct(response)
// 	}
// 	err = json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		response.Message = "error parsing application/json: " + err.Error()
// 		return GCFReturnStruct(response)
// 	}
// 	data, err := EditEmail(user_login.Id, conn, user)
// 	if err != nil {
// 		response.Message = err.Error()
// 		return GCFReturnStruct(response)
// 	}
// 	//
// 	response.Status = 200
// 	response.Message = "Berhasil mengubah email" + user_login.Email
// 	responData := bson.M{
// 		"status":  response.Status,
// 		"message": response.Message,
// 		"data":    data,
// 	}
// 	return GCFReturnStruct(responData)
// }

func TambahCatatanHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	// user, err := GetUserLogin(PASETOPUBLICKEYENV, r)
	// if err != nil {
	// 	response.Message = err.Error()
	// 	return GCFReturnStruct(response)
	// }
	// if user.Email != "kia@gmail.com" {
	// 	response.Message = "Anda tidak memiliki akses, email anda : " + user.Email
	// 	return GCFReturnStruct(response)
	// }
	data, err := InsertCatatan(conn, collectionname, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 201
	response.Message = "Berhasil menambah catatan"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return GCFReturnStruct(responData)
}

func GetCatatanHandler(MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	id := GetID(r)
	if id == "" {
		data, err := GetAllCatatan(conn, collectionname)
		if err != nil {
			response.Message = err.Error()
			return GCFReturnStruct(response)
		}
		responData := bson.M{
			"status":  200,
			"message": "Get Success",
			"data":    data,
		}
		//
		return GCFReturnStruct(responData)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	catatan, err := GetCatatanById(conn, collectionname, idparam)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Get Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    catatan,
	}
	return GCFReturnStruct(responData)
}

func EditUpdateHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	if user.Email != "kia@gmail.com" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	id := GetID(r)
	if id == "" {
		response.Message = "Wrong parameter"
		return GCFReturnStruct(response)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid id parameter"
		return GCFReturnStruct(response)
	}
	data, err := UpdateCatatan(idparam, conn, collectionname, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Berhasil mengubah catatan"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data":    data,
	}
	return GCFReturnStruct(responData)
}

func DeleteCatatanHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, collectionname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	user, err := GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	if user.Email != "kia@gmail.com" {
		response.Message = "Anda tidak memiliki akses"
		return GCFReturnStruct(response)
	}
	id := GetID(r)
	if id == "" {
		response.Message = "Wrong parameter"
		return GCFReturnStruct(response)
	}
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Message = "Invalid id parameter"
		return GCFReturnStruct(response)
	}
	err = DeleteCatatan(idparam, collectionname, conn)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 204
	response.Message = "Berhasil menghapus catatan"
	return GCFReturnStruct(response)
}



func GetProfileHandler(PASETOPUBLICKEYENV, MONGOCONNSTRINGENV, dbname string, r *http.Request) string {
	conn := MongoConnect(MONGOCONNSTRINGENV, dbname)
	response.Status = 400
	//
	payload, err := GetUserLogin(PASETOPUBLICKEYENV, r)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	user, err := GetUserFromID(payload.Id, conn)
	if err != nil {
		response.Message = err.Error()
		return GCFReturnStruct(response)
	}
	//
	response.Status = 200
	response.Message = "Get Success"
	responData := bson.M{
		"status":  response.Status,
		"message": response.Message,
		"data": bson.M{
			"_id":          user.ID,
			"nama_lengkap": user.FullName,
			"email":        user.Email,
			"phonenumber":        user.PhoneNumber,
		},
	}
	return GCFReturnStruct(responData)
}
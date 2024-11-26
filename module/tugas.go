package module

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/mhrndiva/kemahasiswaan/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func MongoConnect(dbname string) (db *mongo.Database) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(dbname)
}

func InsertOneDoc(db string, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := MongoConnect(db).Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertMahasiswa(nama string, phonenumber string, jurusan string, npm int, alamat string, email string, poin int) (InsertedID interface{}) {
    var mahasiswa model.Mahasiswa // Menggunakan model.Mahasiswa
    mahasiswa.Nama = nama
    mahasiswa.Npm = npm
    mahasiswa.Phone_number = phonenumber
    mahasiswa.Jurusan = jurusan
    mahasiswa.Alamat = alamat
    mahasiswa.Email = email
    mahasiswa.Poin = poin
    return InsertOneDoc("data_mahasiswa", "mahasiswa", mahasiswa)
}


func GetMahasiswaFromID(_id primitive.ObjectID, db *mongo.Database, col string) (mahasiswa model.Mahasiswa, errs error) {
	karyawan := db.Collection(col)
	filter := bson.M{"_id": _id}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&mahasiswa)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return mahasiswa, fmt.Errorf("no data found for ID %s", _id)
		}
		return mahasiswa, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return mahasiswa, nil
}

func GetMahasiswaFromNPM(npm int) (staf model.Mahasiswa) {
    mahasiswa := MongoConnect("data_mahasiswa").Collection("mahasiswa")
    filter := bson.M{"npm": npm}
    err := mahasiswa.FindOne(context.TODO(), filter).Decode(&staf)
    if err != nil {
        fmt.Printf("GetMahasiswaFromNPM: %v\n", err)
    }
    return staf
}


func GetAllMahasiswa() (data [] model.Mahasiswa) {
	mahasiswa := MongoConnect("data_mahasiswa").Collection("mahasiswa")
	filter := bson.M{}
	cursor, err := mahasiswa.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func UpdateMahasiswa(npm int, updatedMahasiswa model.Mahasiswa) (bool, error) {
	// Connect to the mahasiswa collection
	mahasiswaCollection := MongoConnect("data_mahasiswa").Collection("mahasiswa")

	// Create the filter to find the mahasiswa by npm
	filter := bson.M{"npm": npm}

	// Create the update document with the fields to update
	update := bson.M{
		"$set": bson.M{
			"nama":          updatedMahasiswa.Nama,
			"phone_number":  updatedMahasiswa.Phone_number,
			"jurusan":       updatedMahasiswa.Jurusan,
			"alamat":        updatedMahasiswa.Alamat,
			"email":         updatedMahasiswa.Email,
			"poin":          updatedMahasiswa.Poin,
		},
	}

	// Perform the update operation
	result, err := mahasiswaCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, fmt.Errorf("failed to update mahasiswa: %v", err)
	}

	// Check if any document was matched and updated
	if result.MatchedCount == 0 {
		return false, fmt.Errorf("no mahasiswa found with npm %s", npm)
	}

	return true, nil
}

func DeleteMahasiswaByNPM(npm int) error {
    mahasiswaCollection := MongoConnect("data_mahasiswa").Collection("mahasiswa")
    filter := bson.M{"npm": npm}

    result, err := mahasiswaCollection.DeleteOne(context.TODO(), filter)
    if err != nil {
        return fmt.Errorf("error deleting data for NPM %d: %v", npm, err)
    }

    if result.DeletedCount == 0 {
        return fmt.Errorf("data with NPM %d not found", npm)
    }

    return nil
}


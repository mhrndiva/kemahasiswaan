package module

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/mhrndiva/kemahasiswaan/model"
	//"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//var MongoString string = os.Getenv("MONGOSTRING")

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

func InsertMahasiswa(nama string, phonenumber string, jurusan string, npm int, alamat string, email string) (InsertedID interface{}) {
	var mahasiswa model.Mahasiswa
	mahasiswa.Nama = nama
	mahasiswa.Npm = npm
	mahasiswa.Phone_number = phonenumber
	mahasiswa.Jurusan = jurusan
	mahasiswa.Alamat = alamat
	mahasiswa.Email = email
	return InsertOneDoc("kemahasiswaan", "mahasiswa", mahasiswa)
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

func GetMahasiswaFromNPM(npm string) (staf model.Mahasiswa) {
	mahasiswa:= MongoConnect("kemahasiswaan").Collection("mahasiswa")
	filter := bson.M{"npm": npm}
	err := mahasiswa.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getMahasiswaFromNPM: %v\n", err)
	}
	return staf
}

func GetAllMahasiswa() (data [] model.Mahasiswa) {
	mahasiswa := MongoConnect("kemahasiswaan").Collection("mahasiswa")
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

func InsertMatkul(namamatkul string, jadwal string, sks int, dosen string) (InsertedID interface{}) {
    var matkul model.Matkul
    matkul.Nama_matkul = namamatkul
    matkul.Jadwal = jadwal
    matkul.Sks = sks
    matkul.Dosen = dosen
    return InsertOneDoc("kemahasiswaan", "matkul", matkul)
}

func GetAllMatkul() (data [] model.Matkul) {
	matkul := MongoConnect("kemahasiswaan").Collection("matkul")
	filter := bson.M{}
	cursor, err := matkul.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func InsertPresensi(db *mongo.Database,col string, npm int, matkul model.Matkul, biodata model.Mahasiswa, checkin string) (insertedID primitive.ObjectID, err error) {
	presensi := bson.M{
		"npm":      npm,
		"datetime": primitive.NewDateTimeFromTime(time.Now().UTC()),
		"matkul":   matkul,
		"biodata":	biodata,
		"checkin":	checkin,
	}
	result, err := db.Collection(col).InsertOne(context.Background(), presensi)
	if err != nil {
		fmt.Printf("InsertPresensi: %v\n", err)
		return
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func GetAllPresensi() (data [] model.Presensi) {
	presensi := MongoConnect("kemahasiswaan").Collection("presensi")
	filter := bson.M{}
	cursor, err := presensi.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}


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

func GetAllDosen() (data [] model.Dosen) {
	dosen := MongoConnect("data_mahasiswa").Collection("dosen")
	filter := bson.M{}
	cursor, err := dosen.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("GetALLData :", err)
	}
	err = cursor.All(context.TODO(), &data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
// func UpdateMahasiswaByID(id string, updatedMahasiswa Mahasiswa) (bool, error) {
// 	// Connect ke koleksi mahasiswa
// 	mahasiswaCollection := MongoConnect("data_mahasiswa").Collection("mahasiswa")

// 	// Konversi ID string ke ObjectID MongoDB
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return false, fmt.Errorf("invalid ID format: %v", err)
// 	}

// 	// Filter untuk menemukan mahasiswa berdasarkan ID
// 	filter := bson.M{"_id": objectID}

// 	// Dokumen update dengan data baru
// 	update := bson.M{
// 		"$set": bson.M{
// 			"nama":          updatedMahasiswa.Nama,
// 			"npm":           updatedMahasiswa.Npm,
// 			"phone_number":  updatedMahasiswa.Phone_number,
// 			"jurusan":       updatedMahasiswa.Jurusan,
// 			"alamat":        updatedMahasiswa.Alamat,
// 			"email":         updatedMahasiswa.Email,
// 			"poin":          updatedMahasiswa.Poin,
// 		},
// 	}

// 	// Eksekusi update
// 	result, err := mahasiswaCollection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		return false, fmt.Errorf("failed to update mahasiswa: %v", err)
// 	}

// 	// Cek apakah ada dokumen yang ditemukan dan diupdate
// 	if result.MatchedCount == 0 {
// 		return false, fmt.Errorf("no mahasiswa found with ID %s", id)
// 	}

// 	return true, nil
// }

func UpdateMahasiswaByID(db *mongo.Database, col string, id primitive.ObjectID, nama string, npm int, phone_number string, jurusan string, alamat string, email string, poin int) error {
    // Filter untuk menemukan dokumen berdasarkan ID
    filter := bson.M{"_id": id}

    // Dokumen update
    update := bson.M{
        "$set": bson.M{
            "nama":         nama,
            "npm":          npm,
            "phone_number": phone_number,
            "jurusan":      jurusan,
            "alamat":       alamat,
            "email":        email,
            "poin":         poin, // Menambahkan poin ke update
        },
    }

    // Eksekusi update
    result, err := db.Collection(col).UpdateOne(context.Background(), filter, update)
    if err != nil {
        return fmt.Errorf("failed to update mahasiswa with ID %v: %w", id, err)
    }

    // Periksa apakah dokumen ditemukan
    if result.MatchedCount == 0 {
        return errors.New("no document found with the specified ID")
    }

    // Periksa apakah dokumen diubah
    if result.ModifiedCount == 0 {
        return errors.New("no data has been changed with the specified ID")
    }

    return nil
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

func DeleteDosenByKodeDosen(kode_dosen int) error {
    dosenCollection := MongoConnect("data_mahasiswa").Collection("dosen")
    filter := bson.M{"kode_dosen": kode_dosen}

    result, err := dosenCollection.DeleteOne(context.TODO(), filter)
    if err != nil {
        return fmt.Errorf("error deleting data for Kode Dosen %d: %v", kode_dosen, err)
    }

    if result.DeletedCount == 0 {
        return fmt.Errorf("data with Kode Dosen %d not found", kode_dosen)
    }

    return nil
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

//untuk dosen
func InsertDosen(nama string, kode_dosen int, phone_number string, matkul string, email string) (insertedID interface{}) {
    var dosen model.Dosen
    dosen.Nama = nama
    dosen.Kode_dosen = kode_dosen
    dosen.Phone_number = phone_number
    dosen.Matkul = matkul
    dosen.Email = email

    return InsertOneDoc("data_mahasiswa", "dosen", dosen)
}

// func UpdateDosen(kode_dosen int, updatedDosen model.Dosen) (bool, error) {
//     // Connect to the mahasiswa collection
//     dosenCollection := MongoConnect("data_mahasiswa").Collection("dosen")

//     // Create the filter to find the dosen by kode_dosen
//     filter := bson.M{"kode_dosen": kode_dosen}

//     // Create the update document with the fields to update
//     update := bson.M{
//         "$set": bson.M{
//             "nama":          updatedDosen.Nama,
//             "phone_number":  updatedDosen.Phone_number,
//             "matkul":        updatedDosen.Matkul,
//             "email":         updatedDosen.Email,
//         },
//     }

//     // Perform the update operation
//     result, err := dosenCollection.UpdateOne(context.TODO(), filter, update)
//     if err != nil {
//         return false, fmt.Errorf("failed to update dosen: %v", err)
//     }

//     // Check if any document was matched and updated
//     if result.MatchedCount == 0 {
//         return false, fmt.Errorf("no dosen found with kode_dosen %d", kode_dosen)
//     }

//     return true, nil
// }

// func DeleteDosenByKodeDosen(kode_dosen int) error {
//     dosenCollection := MongoConnect("data_mahasiswa").Collection("dosen")
//     filter := bson.M{"kode_dosen": kode_dosen}

//     // Melakukan penghapusan berdasarkan kode_dosen
//     result, err := dosenCollection.DeleteOne(context.TODO(), filter)
//     if err != nil {
//         return fmt.Errorf("error deleting data for kode_dosen %d: %v", kode_dosen, err)
//     }

//     // Jika tidak ada data yang dihapus
//     if result.DeletedCount == 0 {
//         return fmt.Errorf("data with kode_dosen %d not found", kode_dosen)
//     }

//     return nil
// }

func GetDosenFromKodeDosen(kode_dosen int) (staf model.Dosen) {
    dosen := MongoConnect("data_mahasiswa").Collection("dosen")
    filter := bson.M{"kode_dosen": kode_dosen}
    err := dosen.FindOne(context.TODO(), filter).Decode(&staf)
    if err != nil {
        fmt.Printf("GetDosenFromKodeDosen: %v\n", err)
    }
    return staf
}
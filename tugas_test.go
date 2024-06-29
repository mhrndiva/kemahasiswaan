package _714220050

import (
    "fmt"
    "testing"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/mhrndiva/kemahasiswaan/model"
    "github.com/mhrndiva/kemahasiswaan/module"
)

func TestInsertPresensi(t *testing.T) {
	var biodata = model.Mahasiswa{
		Nama:     "Devi Wulandari",
		Npm:  714220504,
		Jurusan: "Informatika",
		Phone_number:        "081329563526",
		Alamat:       "cikutra",
		Email:      "devi@gmail.com",
		
	}

	long := 98.345345
	lat := 123.561651
	lokasi := "Amsterdam"
	phonenumber := "6811110023231"
	checkin := "masuk"
	biodata := model.Karyawan{
		Nama:        "Ruud Gullit",
		Phone_number: "628456456222222",
		Jabatan:     "Football Player",
		Jam_kerja:   []model.JamKerja{jamKerja1, jamKerja2},
		Hari_kerja:  []string{"Senin", "Selasa"},
	}
	insertedID, err := module.InsertPresensi(module.MongoConn, "presensi", long, lat, lokasi, phonenumber, checkin, biodata)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}
// func TestInsertMahasiswa(t *testing.T) {
//     // Menjalankan fungsi InsertMahasiswa
//     insertedID := module.InsertMahasiswa("Devi", "71422054", "informatika")
//     if insertedID == nil {
//         t.Error("Expected non-nil value, got nil")
//     }
//     fmt.Println(insertedID)
// }

// func TestGetMahasiswaFromPhoneNumber(t *testing.T) {
//     // Menjalankan fungsi GetMahasiswaFromPhoneNumber
//     phoneNumber := "714220054"
//     mahasiswa := module.GetMahasiswaFromPhoneNumber(phoneNumber)
//     if mahasiswa.Nama == "Devi" {
//         t.Error("Expected non-empty name, got empty")
//     }
//     fmt.Println(mahasiswa.Nama)
// }

// func TestGetAllMahasiswa(t *testing.T) {
//     // Menjalankan fungsi GetAllMahasiswa
//     allMahasiswa := module.GetAllMahasiswa()
//     if len(allMahasiswa) == 0 {
//         t.Error("Expected non-empty slice, got empty slice")
//     }
//     fmt.Println(allMahasiswa)
// }


// func TestInsertMatkul(t *testing.T) {
//     // Menjalankan fungsi InsertMatkul
//     namaMatkul := "pemograman"
//     jamMasuk := "09:30"
//     hari := []string{"sabtu", "kamis"}
//     sks := 2
//     dosen := "pak indra"

//     insertedID := module.InsertMatkul(namaMatkul, jamMasuk, hari, sks, dosen)
//     if insertedID == nil {
//         t.Error("Expected non-nil value, got nil")
//     }
//     fmt.Println(insertedID)
// }

// func TestInsertPresensi(t *testing.T) {
//     // Menjalankan fungsi InsertPresensi
//     phoneNumber := "714220054"
//     datetime := primitive.NewDateTimeFromTime(time.Now().UTC())
//     biodata := model.Mahasiswa{Nama: "Devi", Phone_number: "714220054", Jurusan: "informatika"}

//     insertedID := module.InsertPresensi(phoneNumber, datetime, biodata,)
//     if insertedID == nil {
//         t.Error("Expected non-nil value, got nil")
//     }
//     fmt.Println(insertedID)
// }



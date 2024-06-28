package _714220050

import (
    "fmt"
    "testing"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/mhrndiva/kemahasiswaan/model"
    "github.com/mhrndiva/kemahasiswaan/module"
)

func TestInsertMahasiswa(t *testing.T) {
    // Menjalankan fungsi InsertMahasiswa
    insertedID := module.IsertMahasiswa("Devi", "71422054", "informatika")
    if insertedID == nil {
        t.Error("Expected non-nil value, got nil")
    }
    fmt.Println(insertedID)
}

func TestGetMahasiswaFromPhoneNumber(t *testing.T) {
    // Menjalankan fungsi GetMahasiswaFromPhoneNumber
    phoneNumber := "714220054"
    mahasiswa := module.GetMahasiswaFromPhoneNumber(phoneNumber)
    if mahasiswa.Nama == "Devi" {
        t.Error("Expected non-empty name, got empty")
    }
    fmt.Println(mahasiswa.Nama)
}

func TestGetAllMahasiswa(t *testing.T) {
    // Menjalankan fungsi GetAllMahasiswa
    allMahasiswa := module.GetAllMahasiswa()
    if len(allMahasiswa) == 0 {
        t.Error("Expected non-empty slice, got empty slice")
    }
    fmt.Println(allMahasiswa)
}


func TestInsertMatkul(t *testing.T) {
    // Menjalankan fungsi InsertMatkul
    namaMatkul := "pemograman"
    jamMasuk := "09:30"
    hari := []string{"sabtu", "kamis"}
    sks := 2
    dosen := "pak indra"

    insertedID := module.InsertMatkul(namaMatkul, jamMasuk, hari, sks, dosen)
    if insertedID == nil {
        t.Error("Expected non-nil value, got nil")
    }
    fmt.Println(insertedID)
}

func TestInsertPresensi(t *testing.T) {
    // Menjalankan fungsi InsertPresensi
    phoneNumber := "714220054"
    datetime := primitive.NewDateTimeFromTime(time.Now().UTC())
    biodata := model.Mahasiswa{Nama: "Devi", Phone_number: "714220054", Jurusan: "informatika"}

    insertedID := module.InsertPresensi(phoneNumber, datetime, biodata)
    if insertedID == nil {
        t.Error("Expected non-nil value, got nil")
    }
    fmt.Println(insertedID)
}



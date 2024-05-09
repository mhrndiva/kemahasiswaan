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
    insertedID := module.IsertMahasiswa("Abel", "087654321", "pertanian")
    if insertedID == nil {
        t.Error("Expected non-nil value, got nil")
    }
    fmt.Println(insertedID)
}

func TestGetMahasiswaFromPhoneNumber(t *testing.T) {
    // Menjalankan fungsi GetMahasiswaFromPhoneNumber
    phoneNumber := "087654321"
    mahasiswa := module.GetMahasiswaFromPhoneNumber(phoneNumber)
    if mahasiswa.Nama == "Abel" {
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
    namaMatkul := "Proteksi Tanaman"
    jamMasuk := "09:30"
    hari := []string{"sabtu", "kamis"}
    sks := 2
    dosen := "Bu sinta"

    insertedID := module.InsertMatkul(namaMatkul, jamMasuk, hari, sks, dosen)
    if insertedID == nil {
        t.Error("Expected non-nil value, got nil")
    }
    fmt.Println(insertedID)
}

func TestInsertPresensi(t *testing.T) {
    // Menjalankan fungsi InsertPresensi
    phoneNumber := "087654321"
    datetime := primitive.NewDateTimeFromTime(time.Now().UTC())
    biodata := model.Mahasiswa{Nama: "Abel", Phone_number: "087654321", Jurusan: "pertanian"}

    insertedID := module.InsertPresensi(phoneNumber, datetime, biodata)
    if insertedID == nil {
        t.Error("Expected non-nil value, got nil")
    }
    fmt.Println(insertedID)
}



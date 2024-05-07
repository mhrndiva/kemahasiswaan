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
    insertedID := module.IsertMahasiswa("John Doe", "123456789", "Teknik Informatika")
    if insertedID == nil {
        t.Error("Expected non-nil value, got nil")
    }
    fmt.Println(insertedID)
}

func TestGetMahasiswaFromPhoneNumber(t *testing.T) {
    // Menjalankan fungsi GetMahasiswaFromPhoneNumber
    phoneNumber := "123456789"
    mahasiswa := module.GetMahasiswaFromPhoneNumber(phoneNumber)
    if mahasiswa.Nama == "" {
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
    namaMatkul := "Matematika Dasar"
    jamMasuk := "08:00"
    hari := []string{"Senin", "Rabu"}
    sks := 3
    dosen := "Dr. Andi"

    insertedID := module.InsertMatkul(namaMatkul, jamMasuk, hari, sks, dosen)
    if insertedID == nil {
        t.Error("Expected non-nil value, got nil")
    }
    fmt.Println(insertedID)
}

func TestInsertPresensi(t *testing.T) {
    // Menjalankan fungsi InsertPresensi
    phoneNumber := "123456789"
    datetime := primitive.NewDateTimeFromTime(time.Now().UTC())
    biodata := model.Mahasiswa{Nama: "John Doe", Phone_number: "123456789", Jurusan: "Teknik Informatika"}

    insertedID := module.InsertPresensi(phoneNumber, datetime, biodata)
    if insertedID == nil {
        t.Error("Expected non-nil value, got nil")
    }
    fmt.Println(insertedID)
}



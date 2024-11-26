package _714220050

import (
    "fmt"
    "testing"
    "github.com/mhrndiva/kemahasiswaan/module"
)

func TestInsertMahasiswa(t *testing.T) {
    nama := "Ariadiva"
    npm := 714220050
    phonenumber := "081329503526"
    jurusan := "Informatika"
    alamat := "cikutra"
    email := "devi@gmail.com"
    poin := 20

    hasil := module.InsertMahasiswa(nama, phonenumber, jurusan, npm, alamat, email, poin)
    fmt.Println(hasil)

    if hasil == nil {
        t.Errorf("Expected a non-nil result, got %v", hasil)
    }
}

func TestGetMahasiswaFromNPM(t *testing.T) {
	npm := 714220050 // Use the correct NPM for testing

	// Call the GetMahasiswaFromNPM function
	mahasiswa := module.GetMahasiswaFromNPM(npm)

	// Check if the result is valid
	if mahasiswa.Nama == "" {
		t.Errorf("Expected mahasiswa with npm '%d', but got an empty result", npm)
	} else {
		fmt.Println(mahasiswa)
	}
}
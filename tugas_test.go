package _714220050

import (
    "fmt"
    "testing"
    "github.com/mhrndiva/kemahasiswaan/module"
)

func TestInsertMahasiswa(t *testing.T) {
    nama := "Devi Wulandari"
    npm := 714220504
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

func TestUpdateMahasiswa(t *testing.T) {

	npm := "714220504" 
	updatedMahasiswa := module.Mahasiswa{
		Nama:         "Devi Wulandari",
		Phone_number: "081329503527", 
		Jurusan:      "Informatika",
		Alamat:       "Cikutra",
		Email:        "devi_updated@gmail.com",
		Poin:         25, // Updated points
	}

	// Call the UpdateMahasiswa function
	success, err := module.UpdateMahasiswa(npm, updatedMahasiswa)

	// Check for errors and assert the result
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	// If the update was not successful, fail the test
	if !success {
		t.Errorf("Expected update to be successful, but it wasn't.")
	}

	// Optionally, retrieve the updated mahasiswa and assert the changes
	// This requires accessing the database and checking the values
	mahasiswa := module.GetMahasiswaFromNPM(npm)
	if mahasiswa.Nama != updatedMahasiswa.Nama {
		t.Errorf("Expected Name to be '%s', but got '%s'", updatedMahasiswa.Nama, mahasiswa.Nama)
	}
	if mahasiswa.Phone_number != updatedMahasiswa.Phone_number {
		t.Errorf("Expected Phone_number to be '%s', but got '%s'", updatedMahasiswa.Phone_number, mahasiswa.Phone_number)
	}
	if mahasiswa.Jurusan != updatedMahasiswa.Jurusan {
		t.Errorf("Expected Jurusan to be '%s', but got '%s'", updatedMahasiswa.Jurusan, mahasiswa.Jurusan)
	}
	if mahasiswa.Alamat != updatedMahasiswa.Alamat {
		t.Errorf("Expected Alamat to be '%s', but got '%s'", updatedMahasiswa.Alamat, mahasiswa.Alamat)
	}
	if mahasiswa.Email != updatedMahasiswa.Email {
		t.Errorf("Expected Email to be '%s', but got '%s'", updatedMahasiswa.Email, mahasiswa.Email)
	}
	if mahasiswa.Poin != updatedMahasiswa.Poin {
		t.Errorf("Expected Poin to be '%d', but got '%d'", updatedMahasiswa.Poin, mahasiswa.Poin)
	}

	// Print success for clarity
	fmt.Println("Test passed, mahasiswa updated successfully.")
}
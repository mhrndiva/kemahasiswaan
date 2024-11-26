package _714220050

import (
    "fmt"
    "testing"
    "github.com/mhrndiva/kemahasiswaan/module"
)

func TestInsertMahasiswa(t *testing.T) {
    nama := "Serli Pariela"
    npm := 714220023
    phonenumber := "081329503526"
    jurusan := "Informatika"
    alamat := "Batu Jajar"
    email := "Serli@gmail.com"
    poin := 25

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

func TestGetAllMahasiswa(t *testing.T) {
	// Call the GetAllMahasiswa function to retrieve all mahasiswa data
	mahasiswaList := module.GetAllMahasiswa()

	// Check if data is returned
	if len(mahasiswaList) == 0 {
		t.Errorf("Expected non-empty mahasiswa list, but got an empty list")
	} else {
		// Print out the list of mahasiswa
		fmt.Println("Mahasiswa List:", mahasiswaList)
	}
}

func TestDeleteMahasiswaByNPM(t *testing.T) {
	// Define the NPM of the mahasiswa to delete
	npm := 714220050 // Use the correct NPM for testing

	// Call the DeleteMahasiswaByNPM function
	err := module.DeleteMahasiswaByNPM(npm)

	// Check if there was an error
	if err != nil {
		t.Errorf("Failed to delete mahasiswa with npm %d: %v", npm, err)
	} else {
		// Print confirmation message
		fmt.Printf("Mahasiswa with npm %d deleted successfully\n", npm)
	}
}

//ini buat dosen
func TestInsertDosen(t *testing.T) {
    nama := "Nisa"
    kode_dosen := 102
    phone_number := "081234567890"
    matkul := "Pemrograman V"
    email := "Nisa@ulbi.com"

   
    hasil := module.InsertDosen(nama, kode_dosen, phone_number, matkul, email)
    fmt.Println(hasil)

    if hasil == nil {
        t.Errorf("Expected a non-nil result, got %v", hasil)
    }
}

func TestDeleteDosenByKodeDosen(t *testing.T) {
    kode_dosen := 102  // Ganti dengan kode dosen yang valid

    err := module.DeleteDosenByKodeDosen(kode_dosen)

    // Check if there was an error
    if err != nil {
        t.Errorf("Failed to delete dosen with kode_dosen %d: %v", kode_dosen, err)
    } else {
        // Print confirmation message
        fmt.Printf("Dosen with kode_dosen %d deleted successfully\n", kode_dosen)
    }
}

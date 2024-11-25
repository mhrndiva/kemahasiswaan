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

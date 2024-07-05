package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Mahasiswa struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty" json:"nama,omitempty"`
	Npm          int				 `bson:"npm,omitempty" json:"npm,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	Jurusan      string             `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	Alamat		 string				 `bson:"alamat,omitempty" json:"alamat,omitempty"`
	Email		 string				 `bson:"email,omitempty" json:"email,omitempty"`
}

type Matkul struct {
	Nama_matkul   string   `bson:"namamatkul,omitempty" json:"namamatkul,omitempty"`
	Jadwal        string  `bson:"jadwal,omitempty" json:"jadwal,omitempty"`
	Sks      	  int      `bson:"sks,omitempty" json:"sks,omitempty"`
	Dosen         string   `bson:"dosen,omitempty" json:"dosen,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Npm			 int            `bson:"npm,omitempty" json:"npm,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty" json:"datetime,omitempty"`
	Matkul       Matkul               `bson:"matkul,omitempty" json:"matkul,omitempty"`
	Biodata      Mahasiswa           `bson:"biodata,omitempty" json:"biodata,omitempty"`
	Checkin		 string					`bson:"checkin,omitempty" json:"checkin,omitempty"`
}

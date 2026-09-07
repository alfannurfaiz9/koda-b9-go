package model

type RiwayatPendidikan struct {
	Nama    string
	Jurusan string
}

type Biodata struct {
	Nama              string
	Foto              string
	Email             string
	Umur              uint8
	NomorTelepon      string
	StatusPernikahan  bool
	RiwayatPendidikan RiwayatPendidikan
}

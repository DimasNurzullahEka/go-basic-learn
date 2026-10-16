package entities

type Mahasiswa struct {
	Id            int64  `json:"id"`
	NamaLengkap   string `json:"nama_lengkap"`
	JenisKelamin  string `json:"jenis_kelamin"`
	TempatLahir   string `json:"tempat_lahir"`
	TanggalLahir  string `json:"tanggal_lahir"`
	Alamat        string `json:"alamat"`
}

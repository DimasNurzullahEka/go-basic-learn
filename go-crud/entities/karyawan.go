package entities

type Karyawan struct {
	ID            int64  `json:"id"`
	Nama_karyawan string `json:"nama_karyawan"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Alamat        string `json:"alamat"`
}

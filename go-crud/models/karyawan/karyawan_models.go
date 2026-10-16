package karyawan

import (
	"database/sql"

	"github.com/jeypc/go-crud-modal/config"
	"github.com/jeypc/go-crud-modal/entities"
)

type Karyawanmodels struct {
	db *sql.DB
}

func New() *Karyawanmodels {
	db, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &Karyawanmodels{db: db}
}
func (m *Karyawanmodels) FindAll(karyawan *[]entities.Karyawan) error {
	rows, err := m.db.Query("select * from karyawan")
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var data entities.Karyawan
		rows.Scan(
			&data.ID,
			&data.Nama_karyawan,
			&data.Jenis_kelamin,
			&data.Alamat)
		*karyawan = append(*karyawan, data)
	}

	return nil
}

package karyawan

import (
	"database/sql"

	"github.com/jeypc/go-crud-modal/config"
	"github.com/jeypc/go-crud-modal/entities"
)

type KaryawanModel struct {
	db *sql.DB
}

func New() *KaryawanModel {
	db, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &KaryawanModel{db: db}
}

func (m *KaryawanModel) FindAll(karyawan *[]entities.Karyawan) error {
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

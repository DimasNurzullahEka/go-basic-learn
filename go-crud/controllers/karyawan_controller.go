package controllers

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/jeypc/go-crud-modal/entities"
	karyawanModel "github.com/jeypc/go-crud-modal/models/karyawan"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/karyawan/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := temp.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetDataKaryawan() string {
	buffer := &bytes.Buffer{}

	temp, err := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/karyawan/data.html")
	if err != nil {
		return "Error template: " + err.Error()
	}

	var karyawan []entities.Karyawan
	model := karyawanModel.New()
	err = model.FindAll(&karyawan)
	if err != nil {
		return "Error: " + err.Error()
	}

	data := map[string]interface{}{
		"karyawan": karyawan,
	}

	err = temp.Execute(buffer, data)
	if err != nil {
		return "Error execute: " + err.Error()
	}

	return buffer.String()
}

package controllers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/jeypc/go-crud-modal/entities"
	karyawanModel "github.com/jeypc/go-crud-modal/models/karyawan"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"data": template.HTML(GetDataKaryawan()),
	}

	path, err := karyawanViewPath("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temp, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := temp.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetDataKaryawan() string {
	buffer := &bytes.Buffer{}

	path, err := karyawanViewPath("data.html")
	if err != nil {
		return "Error template: " + err.Error()
	}

	temp, err := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles(path)
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

func GetForm(w http.ResponseWriter, r *http.Request) {
	path, err := karyawanViewPath("form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	temp, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := temp.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func karyawanViewPath(name string) (string, error) {
	_, sourceFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("tidak dapat menentukan lokasi template karyawan")
	}

	return filepath.Join(filepath.Dir(sourceFile), "..", "views", "karyawan", name), nil
}

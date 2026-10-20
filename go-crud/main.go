package main

import (
	"log"
	nethttp "net/http"

	karyawancontroller "github.com/jeypc/go-crud-modal/controllers"
)

func main() {
	http := struct {
		HandleFunc     func(string, func(nethttp.ResponseWriter, *nethttp.Request))
		listenAndServe func(string, nethttp.Handler) error
	}{
		HandleFunc:     nethttp.HandleFunc,
		listenAndServe: nethttp.ListenAndServe,
	}

	http.HandleFunc("/", karyawancontroller.Index)
	http.HandleFunc("/karyawan/get-form", karyawancontroller.GetForm)
	log.Fatal(http.listenAndServe(":8080", nil))
}

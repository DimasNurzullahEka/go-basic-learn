package main

import "fmt"

func main() {
	namaInisal := "Abdillah Dimas"
	message := Hello(Person{Name: namaInisal})
	fmt.Printf("Halo, %v! %v", namaInisal, message)
}

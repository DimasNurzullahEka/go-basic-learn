package main

import "fmt"

func Ekspresi() {
	ekspresi_mertua := 3
	switch ekspresi_mertua {
	case 1:
		fmt.Println("Ekspresi mertua: Senang")
	case 2:
		fmt.Println("Ekspresi mertua: Sedih")
	case 3:
		fmt.Println("Ekspresi mertua: Marah")
	default:
		fmt.Println("Ekspresi mertua: Tidak diketahui")
	}
}

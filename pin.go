package main

import "fmt"

func main() {
	userpin := "2006"
	var pin string

	for percobaan := 1; percobaan <= 2; percobaan++  {
		fmt.Print("masukan pin anda: ")
	    fmt.Scanln(&pin)
		if userpin == pin {
			fmt.Println("pin benar")
			break

		}
		fmt.Println("pin salah")
		if percobaan == 2 {
			fmt.Println("akses di blokir")
		}
	
	}
	

}
package main

import "fmt"

func main (){

	var angka1, angka2 float64
	var operasiBilangan string

	fmt.Print("Masukan Operasi Bilangan (+ - * /): ")
	fmt.Scanln(&operasiBilangan)

	fmt.Print("Masukan Angka Pertama: ")
	fmt.Scanln(&angka1)

	fmt.Print("Masukan Angka Kedua: ")
	fmt.Scanln(&angka2)


	switch operasiBilangan {

	case "+" :
		fmt.Printf("%f %s %f = %f", angka1, operasiBilangan, angka2, angka1 + angka2)
	
	case "-" :
		fmt.Printf("%f %s %F = %f", angka1, operasiBilangan, angka2, angka1 - angka2)

	case "*" :
		fmt.Printf("%f %s %f = %f", angka1, operasiBilangan, angka2, angka1 * angka2)

	case "/" :
		if angka2 == 0.0 {
			fmt.Println("membagi dengan situasi nol")
		}else{
			fmt.Printf("%f %s %f = %f", angka1, operasiBilangan, angka2, angka1 / angka2)
		}

	default:
		fmt.Println("kalkulator tidak mengerti input user")

	}


}
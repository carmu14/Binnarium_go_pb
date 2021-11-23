package main

import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	nombre := "Eduardo"
	fmt.Println(string(nombre[4]))

	for i := 0; i < 7; i++ {
		fmt.Println(i,string(nombre[i]))
	}

	//supliendo while

	numero := 1

	for numero != 1 { //si es falsa se produce un bucle infinito
		fmt.Println("Ejecuto el codigo dentro del for")
	}
	fmt.Println("Texto despues del FOR")

}
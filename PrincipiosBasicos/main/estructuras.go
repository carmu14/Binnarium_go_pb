package main

import "fmt"

func main() {

	numero := 45
	//if
	if numero == 40 {
		fmt.Println("Ejecuto codigo que hay en el if")
	} else if numero == 41 {
		fmt.Println("Ejecuto el codigo de las llaves del else if")
	} else { //es opcional
		//si es falso se ejecuta
		fmt.Println("Ejecuto el codigo del else")
	}

	//ejercicio

	//edad := 30 corresponde al if
	//edad := 18 corresponde al if
	//edad := 7 corresponde al else if
	edad := -2 //corresponde al else

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else if edad < 18 && edad >= 0 {
		fmt.Println("Eres menor de edad")
	} else {
		fmt.Println("La edad no es valida")
	}
}

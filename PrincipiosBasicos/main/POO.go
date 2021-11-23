package main

import "fmt"

type persona struct {
	nombre string
	apellido string
	edad int
}


func main() {

	persona1 := persona{nombre: "Carlos", apellido: "Munoz", edad: 32}

	persona2 := persona{nombre: "Maria", apellido: "Palacios", edad: 28}

	fmt.Println("Persona1 es igual a", persona1)
	fmt.Println("Persona2 es igual a", persona2)

	persona2.edad = 47 //cambiar valor a la variable

	fmt.Println("Persona2 es igual a", persona2)

}

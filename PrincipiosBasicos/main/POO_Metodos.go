package main

import "fmt"

type persona struct {
	nombre string
	apellido string
	edad int
}

func (p persona) saludar (saludo string)  {
	fmt.Println(saludo + ", " + p.nombre)
}

func (pers persona) cumple () int {
	return pers.edad + 1
}

func main() {

	persona1 := persona{nombre: "Carlos", apellido: "Munoz", edad: 32}

	persona2 := persona{nombre: "Maria", apellido: "Palacios", edad: 28}

	fmt.Println("Persona1 es igual a", persona1)
	fmt.Println("Persona2 es igual a", persona2)

	persona2.edad = 47 //cambiar valor a la variable

	fmt.Println("Persona2 es igual a", persona2)

	persona1.saludar("Hola")
	persona2.saludar("Hola")

	fmt.Println(persona1.cumple())

	edad_persona2 := persona2.cumple()
	fmt.Println(edad_persona2)
}

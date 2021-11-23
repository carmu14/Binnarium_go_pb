package main

import "fmt"

func main() {
	mi_funcion()
	mi_funcion()
	mi_funcion()

	saludar("Carlos")

	sumar(5, 8)
	fmt.Println(sumar(100, 10))

	suma := sumar(10,41)
	fmt.Println(suma)
}

func mi_funcion()  {
	fmt.Println("Esta es una linea de codigo q voy a cambiar")
	fmt.Println("Esta es otra linea de codigo")
}

func saludar(nombre string)  {
	fmt.Println("Hola " + nombre)
}

//Funciones que nos retorna un valor

func sumar(numero1 int, numero2 int) int {
	suma := numero1 + numero2
	return suma
}
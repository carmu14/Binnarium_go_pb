package main

import "fmt"

func main() {

	var entero int
	var decimal float64
	var texto string
	var booleana bool

	entero = 5
	decimal = 10.4
	texto = "Bienvenido"
	booleana = true

	fmt.Println(entero,decimal,texto, booleana)

	//la declaracion de variables en una sola linea
	numero := 10

	fmt.Println(numero)

	//comentario

	//operadores matematicos
	suma := 10 + 5
	resta := 10 - 5
	multiplicacion := 10 * 5
	division := 10 / 5
	restodiv := 10 % 3

	fmt.Println(suma,resta,multiplicacion,division,restodiv)

	//operadores de incremento y
	suma++
	fmt.Println(suma)

	suma--
	fmt.Println(suma)

	suma += 10 //es como decir suma = suma + 10
	fmt.Println(suma)

	//operadores logicos
	fmt.Println(10 > 5)
	fmt.Println(10 < 5)
	fmt.Println(10 >= 10)
	fmt.Println(10 <= 5)
	fmt.Println(10 == 10) //para comparar utilizar dos iguales
	fmt.Println(10 != 5)

	//operadores y
	fmt.Println(10 < 5 && 10 > 4)
	fmt.Println(10 > 5 && 10 > 4)

	//operador o
	fmt.Println(10 < 5 || 10 > 5)
}



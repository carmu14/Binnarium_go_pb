package main

import "fmt"

func main() {

	frutas := [] string {"Pera", "Manzana", "Naranja"}

	//fmt.Println(frutas[1])

	frutas = append(frutas, "Maduro", "Melocoton", "Sandia") //agregar elementos a lista

	frutas [0] = "Mandarina"  //remplazar un elemnto en la lista

	for i := 0; i < 6; i++{
		fmt.Println(frutas[i])
	}

	//buscar elemento
	for i := 0; i < len(frutas); i++{ //len averigua cuantos elementos hay en la lista
		if frutas[i] == "Limon" {

			fmt.Println("Hay un elemnto en la lista")

		}

	}
}

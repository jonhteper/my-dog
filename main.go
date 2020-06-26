package main

import (
	"fmt"
	"github.com/jonhteper/my-dog/game"
)

func main() {
	myDog := game.Husky{
		Age: 1,
	}
	sam := game.Husky{
		Name:  "Sam",
		Age:   2,
		Color: "azúl",
	}

	game.ImprimirLinea()
	fmt.Println("--------- CREAR HUSKY ---------")
	fmt.Println("Nombre de tu husky: ")
	_, err := fmt.Scan(&myDog.Name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Color de tu husky: ")
	_, err = fmt.Scan(&myDog.Color)
	if err != nil {
		fmt.Println(err)
		return
	}

	game.LimpiarVista()
	game.ImprimirLinea()
	myDog.Ladrar()
	fmt.Println("¿Salir a jugar?(si/no)")
	var verificador string
	_, err = fmt.Scan(&verificador)
	if err != nil {
		fmt.Println(err)
		return
	}

	if verificador == "si" {
		game.PrimeraEscena(myDog)
	} else {
		myDog.Despedida()
		return
	}

	sam.Responder(myDog)
	return
}

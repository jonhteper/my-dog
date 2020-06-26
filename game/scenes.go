package game

import (
	"fmt"
	"time"
)

func PrimeraEscena(h Husky) {
	paloma := Animal{
		Species: "ave",
		Name:    "paloma",
		Life:    3,
		Speed:   15,
		Smart:   3,
		Food:    true,
	}
	LimpiarVista()
	ImprimirLinea()
	barraProgreso()
	fmt.Printf("*%v sale*\n", h.Name)
	time.Sleep(1 * time.Second)
	LimpiarVista()
	ImprimirLinea()
	fmt.Printf("%v ha visto un %v. ¿Lo dejas acercarse?(si/no)\n", h.Name, paloma.Species)
	var verificador string
	_, err := fmt.Scan(&verificador)
	if err != nil {
		fmt.Println(err)
		return
	}

	if verificador == "no" {
		fmt.Printf("*%v se aleja*\n", h.Name)
		return
	} else {
		LimpiarVista()
		ImprimirLinea()
		fmt.Printf("Es una %v.\n\t—¡Guau!¿Puedo atraparla?(Si/No/En-casa-hay-comida)\n", paloma.Name)
		_, err := fmt.Scan(&verificador)
		if err != nil {
			fmt.Println(err)
			return
		}

		if verificador == "No" || verificador == "En-casa-hay-comida" {
			h.RespuestaTriste()
			return
		} else {
			fmt.Println("Continuará...")
			return
		}
	}
}

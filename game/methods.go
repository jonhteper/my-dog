package game

import (
	"fmt"
	"time"
)

func (h Husky) Ladrar() {
	time.Sleep(1 * time.Second)
	fmt.Printf("—¡Guau!, soy %v. ¡Ya quiero jugar!.\n", h.Name)
}

func (h Husky) RespuestaTriste() {
	time.Sleep(1 * time.Second)
	fmt.Printf(":( —Está bien...guau\n")
}

func (h Husky) Saludar() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hola, soy", h.Name, ".¿Cómo te llamas?")
}

func (h Husky) Responder(o Husky) {
	time.Sleep(1 * time.Second)
	fmt.Println("Hola,", o.Name, ". Yo soy", h.Name)
}
func (h Husky) Despedida() {
	time.Sleep(1 * time.Second)
	LimpiarVista()
	ImprimirLinea()
	fmt.Println("—¡Guau! Fue divertido. ¡Adiós!")
	ImprimirLinea()
}

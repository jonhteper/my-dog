// +build js,wasm

//Portación del juego a WebAssembly
package main

import (
	"github.com/jonhteper/my-dog/game"
	"log"
	"syscall/js"
	"time"
)

var (
	document               = js.Global().Get("document")
	container              = document.Call("getElementById", "container")
	interactive            = document.Call("getElementById", "interactive")
	myDog       game.Husky = game.Husky{
		Age: 1,
	}
)

func createButton() js.Value {
	button := document.Call("createElement", "button")
	button.Set("appendChild", button)
	button.Set("id", "my-btn")
	button.Set("innerHTML", "Iniciar Juego")

	return button
}

func inicio() (title js.Value, input js.Value) {
	title = document.Call("createElement", "h2")
	title.Set("innerText", "--------- CREAR HUSKY ---------")
	input = document.Call("createElement", "div")
	label := document.Call("createElement", "label")
	text := document.Call("createElement", "input")
	label.Set("innerText", "Nombre de tu husky: ")
	label.Set("htmlFor", "name_husky")
	text.Set("type", "text")
	text.Set("id", "name_husky")
	input.Call("append", label, text)

	input.Call("addEventListener", "change", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		leerNombre(&myDog)
		log.Println(myDog.Name)
		return nil
	}))

	return
}

func leerNombre(h *game.Husky) (nombre string) {
	input := document.Call("getElementById", "name_husky")
	nombre = input.Get("value").String()
	h.Name = nombre
	return
}

func main() {

	channel := make(chan struct{})
	title, input := inicio()
	container.Set("innerHTML", "Loading...")

	go func() {
		time.Sleep(1 * time.Second)
		container.Set("innerHTML", "")
		container.Call("appendChild", title)
		interactive.Set("innerHTML", "")
		interactive.Call("appendChild", input)
	}()

	<-channel
}

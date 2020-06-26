package game

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func ImprimirLinea() {
	fmt.Println("*******************************************************")
}
func LimpiarVista() {
	if runtime.GOOS == "linux" {
		console := exec.Command("clear")
		console.Stdout = os.Stdout
		console.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func barraProgreso() {
	const col = 5
	// Clear the screen by printing \x0c.
	bar := fmt.Sprintf("[%%-%vs]", col)
	for i := 0; i < col; i++ {
		fmt.Printf(bar, strings.Repeat("=", i)+">")
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("")
}

package main

import (
	"Clase20_08/Palindromos"
	"Clase20_08/Prueba_Swich"
	"fmt"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {

	Prueba_Swich.PruebaSwich()
	comprobacion := Palindromos.EncontrarPalindromo("ana")
	if comprobacion == true {
		fmt.Println("Es un palindromo")
	} else {
		fmt.Println("No es un palindromo")
	}

}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

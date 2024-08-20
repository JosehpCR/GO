package Palindromos

import "strings"

func EncontrarPalindromo(palabra string) bool {
	palabra = strings.ToLower(palabra)

	palabra = strings.Replace(palabra, " ", "", -1)

	n := len(palabra)
	for i := 0; i < n/2; i++ {
		if palabra[i] != palabra[n-i-1] {
			return false
		}
	}
	return true
}

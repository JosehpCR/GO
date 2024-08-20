package Prueba_Swich

import (
	"fmt"
	"time"
)

func PruebaSwich() {
	today := time.Now()
	switch today.Day() {
	case 5:
		fmt.Println("Hoy es 5th ")
		fallthrough
	case 10:
		fmt.Println("Hoy es 10th ")
		fallthrough
	case 14:
		fmt.Println("Hoy es 14th ")
		fallthrough
	case 15:
		fmt.Println("Hoy es 15th ")
	case 25, 26, 27:
		fmt.Println("Hoy es 25th ")
	case 31:
		fmt.Println("Fiesta esta noche ")
	default:
		fmt.Println("No information available")

	}

}

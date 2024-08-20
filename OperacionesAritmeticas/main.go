package main

import (
	"fmt"
	"math"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	// Datos de prueba
	ladCuadrado := 6.0
	baseRectangulo := 5.0
	alturaRectangulo := 8.0
	baseTriangulo := 2.0
	alturaTriangulo := 6.0
	diagonalMayorRombo := 9.0
	diagonalMenorRombo := 3.0
	ladoPentagono := 10.0
	apotemaPentagono := 12.0
	ladoHexagono := 4.0
	radioCirculo := 5.0
	baseMayorTrapecio := 12.0
	baseMenorTrapecio := 6.0
	alturaTrapecio := 9.0
	baseParalelogramo := 7.0
	alturaParalelogramo := 5.0

	// Cálculo de áreas
	areaCuadrado := ladCuadrado * ladCuadrado
	areaRectangulo := baseRectangulo * alturaRectangulo
	areaTriangulo := (baseTriangulo * alturaTriangulo) / 2
	areaRombo := (diagonalMayorRombo * diagonalMenorRombo) / 2
	areaPentagono := (5 * ladoPentagono * apotemaPentagono) / 2
	areaHexagono := (3 * math.Sqrt(3) * ladoHexagono * ladoHexagono) / 2
	areaCirculo := math.Pi * radioCirculo * radioCirculo
	areaTrapecio := ((baseMayorTrapecio + baseMenorTrapecio) * alturaTrapecio) / 2
	areaParalelogramo := baseParalelogramo * alturaParalelogramo

	// Mostrar resultados
	fmt.Println("Área del Cuadrado:", areaCuadrado)
	fmt.Println("Área del Rectángulo:", areaRectangulo)
	fmt.Println("Área del Triángulo:", areaTriangulo)
	fmt.Println("Área del Rombo:", areaRombo)
	fmt.Println("Área del Pentágono:", areaPentagono)
	fmt.Println("Área del Hexágono:", areaHexagono)
	fmt.Println("Área del Círculo:", areaCirculo)
	fmt.Println("Área del Trapecio:", areaTrapecio)
	fmt.Println("Área del Paralelogramo:", areaParalelogramo)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

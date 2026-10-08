package main

import "fmt"

func main() {
	notas := [6][4]float64{
		{5.5, 6.0, 4.8, 6.2},
		{4.0, 5.0, 3.5, 4.5},
		{6.8, 7.0, 6.5, 6.9},
		{3.2, 4.5, 5.0, 4.0},
		{5.0, 5.5, 6.0, 5.8},
		{6.0, 6.5, 6.2, 5.9},
	}

	var sumaTotal float64
	totalNotas := 0

	for i := 0; i < len(notas); i++ {
		estudianteNotas := notas[i][:]
		suma := 0.0
		min := estudianteNotas[0]
		max := estudianteNotas[0]

		for _, nota := range estudianteNotas {
			suma += nota
			if nota < min {
				min = nota
			}
			if nota > max {
				max = nota
			}
		}

		promedio := suma / float64(len(estudianteNotas))
		sumaTotal += suma
		totalNotas += len(estudianteNotas)

		fmt.Printf("Estudiante %d:\n", i+1)
		fmt.Printf("  Promedio: %.2f\n", promedio)
		fmt.Printf("  Nota mas alta: %.2f\n", max)
		fmt.Printf("  Nota mas baja: %.2f\n\n", min)
	}

	promedioGeneral := sumaTotal / float64(totalNotas)
	fmt.Printf("Promedio general de la clase: %.2f\n", promedioGeneral)
}
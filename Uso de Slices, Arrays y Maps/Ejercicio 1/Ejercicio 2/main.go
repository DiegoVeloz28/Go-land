package main

import (
	"fmt"
	"strings"
)

func obtenerGanador(votos map[string]int) string {
	ganador := ""
	maxVotos := -1

	for actividad, cantidad := range votos {
		if cantidad > maxVotos {
			maxVotos = cantidad
			ganador = actividad
		}
	}

	return ganador
}

func main() {
	votos := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"musica":      0,
	}

	fmt.Println("Actividades disponibles: deportes, videojuegos, cine, musica")

	for i := 1; i <= 5; i++ {
		var eleccion string
		for {
			fmt.Printf("Ingrese el voto %d: ", i)
			fmt.Scan(&eleccion)
			eleccion = strings.ToLower(eleccion)

			if _, existe := votos[eleccion]; existe {
				votos[eleccion]++
				break
			}

			fmt.Println("Opcion no valida. Intente nuevamente.")
		}
	}

	fmt.Println("\nResultados de la votacion:")
	for actividad, cantidad := range votos {
		fmt.Printf("- %s: %d votos\n", actividad, cantidad)
	}

	ganador := obtenerGanador(votos)
	fmt.Printf("\nActividad con mayor aceptacion: %s (%d votos)\n", ganador, votos[ganador])
}
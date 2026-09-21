package main

import "fmt"

func main() {
	var op int
	fmt.Println("Bienvenido al menú")
	fmt.Println("Dispone de todas las siguientes opciones:")
	fmt.Println("Opción 1: Notas de los estudiantes")
	fmt.Println("Opción 2: Suma de números")
	fmt.Println("Opción 3: Convertir Celsius a Fahrenheit")
	fmt.Println("Opción 4: Convertir Fahrenheit a Celsius")
	fmt.Println("Opción 0: Salir")
	fmt.Print("Elija una opción: ")
	fmt.Scan(&op)

	switch op { 
	case 0:
		fmt.Println("Saliendo del programa...")
	case 1:
		var n int
		fmt.Print("Escriba el número de estudiantes que desea ingresar las notas: ")
		fmt.Scan(&n)
		notas(n) 
	case 2:
		suma()
	case 3:
		celsiusToFahrenheit()
	case 4:
		fahrenheitToCelsius()
	default:
		fmt.Println("Opción no válida")
	}
}

// Opción 1
func notas(n int) {
	var nota float64
	var sum float64

	for i := 1; i <= n; i++ {
		fmt.Printf("Ingrese la nota del estudiante %d (0-100): ", i)
		fmt.Scan(&nota)
		sum += nota
	}

	
	promedio := averageGrade(sum, n)
	fmt.Printf("\nEl promedio del curso es: %.2f\n", promedio)

	
	if promedio >= 70 {
		fmt.Println("Estado: Aprobado")
	} else {
		fmt.Println("Estado: Reprobado")
	}

	
	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Rendimiento: Excellent performance")
	case promedio >= 80 && promedio < 90:
		fmt.Println("Rendimiento: Good performance")
	case promedio >= 70 && promedio < 80:
		fmt.Println("Rendimiento: Satisfactory performance")
	default:
		fmt.Println("Rendimiento: Needs improvement")
	}
}

func averageGrade(sum float64, count int) float64 {
	if count == 0 {
		return 0
	}
	return sum / float64(count) 
}

// Opción 2
func suma() {
	var n int
	var total int
	fmt.Print("Ingrese un número límite (n): ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		total += i
	}
	fmt.Printf("La suma de los números del 1 al %d es: %d\n", n, total)
}

// Opción 3
func celsiusToFahrenheit() {
	var c float64
	fmt.Print("Ingrese la temperatura en grados Celsius: ")
	fmt.Scan(&c)
	f := (c * 9 / 5) + 32
	fmt.Printf("%.2f °C equivalen a %.2f °F\n", c, f)
}

// Opción 4
func fahrenheitToCelsius() {
	var f float64
	fmt.Print("Ingrese la temperatura en grados Fahrenheit: ")
	fmt.Scan(&f)
	c := (f - 32) * 5 / 9
	fmt.Printf("%.2f °F equivalen a %.2f °C\n", f, c)
}
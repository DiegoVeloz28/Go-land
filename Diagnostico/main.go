package main
import "fmt"

var productosvendidos [] string
var subtotalesventas [] float64

func RegistroVentas (nombre string, precio float64, cantidad int) {
	subtotal:= precio * float64(cantidad)
	productosvendidos = append(productosvendidos, nombre)
	subtotalesventas = append(subtotalesventas, subtotal)
	 fmt.Printf(" Venta registrada: %d x %s (Subtotal: $%.2f)\n", cantidad, nombre, subtotal)
}
func MostrarEstadisticas(){
	if len(subtotalesVentas) == 0 {
		fmt.Println(" No existen ventas registradas aún.")
		return
	}

	var total float64
	fmt.Println(" --- RESUMEN DE VENTAS ---")
	for i := 0; i < len(productosVendidos); i++ {
		fmt.Printf("Producto: %s | Subtotal: $%.2f\n", productosVendidos[i], subtotalesVentas[i])
		total += subtotalesVentas[i]
	}
	fmt.Printf("----------------------------\n")
	fmt.Printf(" Total recaudado: $%.2f\n", total)
}

func main (){
	fmt.Println("Diagnostico")
	fmt.Println("Bienvenido al menú")
	for {
		fmt.Println("Elija la opción que desee realizar \n Opción 0 = Salir del menú \n Opción 1 = ")
	}

}
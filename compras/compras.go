package compras

import "fmt"

// Calcular el IVA
func CalcularIVA(subtotal float64) float64 {
	return subtotal * 0.15
}

// Calcular el total de la compra
func CalcularTotal(subtotal, iva float64) float64 {
	return subtotal + iva
}

// Mostrar el resumen de la compra
func MostrarResumen(subtotal, iva, total float64) {
	fmt.Println("\n===== RESUMEN DE COMPRA =====")
	fmt.Printf("Subtotal: $%.2f\n", subtotal)
	fmt.Printf("IVA:      $%.2f\n", iva)
	fmt.Printf("TOTAL:    $%.2f\n", total)
	fmt.Println("=============================")
}

package carrito

// Validar la cantidad seleccionada
func ValidarCantidad(cantidad int) bool {
	return cantidad > 0
}

// Calcular el subtotal de un producto
func CalcularSubtotal(precio float64, cantidad int) float64 {
	return precio * float64(cantidad)
}

// Calcular el subtotal de varios productos
func CalcularTotalCarrito(subtotales []float64) float64 {
	total := 0.0

	for _, subtotal := range subtotales {
		total += subtotal
	}

	return total
}

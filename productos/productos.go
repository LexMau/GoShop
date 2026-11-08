package productos

import "fmt"

// Mostrar el catalogo de productos
func MostrarProductos() {
	fmt.Println("===== CATALOGO GOSHOP =====")
	fmt.Println("1. Laptop - $650.00")
	fmt.Println("2. Mouse - $15.00")
	fmt.Println("3. Teclado - $25.00")
	fmt.Println("4. Monitor - $180.00")
}

// Obtener el nombre del producto
func ObtenerNombre(codigo int) string {
	switch codigo {
	case 1:
		return "Laptop"
	case 2:
		return "Mouse"
	case 3:
		return "Teclado"
	case 4:
		return "Monitor"
	default:
		return "Producto no encontrado"
	}
}

// Obtener el precio del producto
func ObtenerPrecio(codigo int) float64 {
	switch codigo {
	case 1:
		return 650.00
	case 2:
		return 15.00
	case 3:
		return 25.00
	case 4:
		return 180.00
	default:
		return 0
	}
}

// Validar si existe el producto
func ProductoValido(codigo int) bool {
	return codigo >= 1 && codigo <= 4
}

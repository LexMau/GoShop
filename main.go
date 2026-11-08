package main

import (
	"GoShop/carrito"
	"GoShop/compras"
	"GoShop/productos"
	"fmt"
)

func main() {

	var opcion int

	for {

		fmt.Println("\n==========================")
		fmt.Println("       BIENVENIDO A")
		fmt.Println("          GOSHOP")
		fmt.Println("==========================")
		fmt.Println("1. Ver productos")
		fmt.Println("2. Realizar una compra")
		fmt.Println("3. Salir")
		fmt.Println("==========================")
		fmt.Print("Seleccione una opcion: ")

		fmt.Scan(&opcion)

		switch opcion {

		case 1:
			productos.MostrarProductos()

		case 2:

			var subtotales []float64
			var codigo int
			var cantidad int
			var continuar string

			continuar = "s"

			for continuar == "s" {

				productos.MostrarProductos()

				fmt.Print("\nIngrese el codigo del producto: ")
				fmt.Scan(&codigo)

				if !productos.ProductoValido(codigo) {
					fmt.Println("Error: el producto seleccionado no existe.")
					fmt.Println("Por favor, seleccione un codigo del 1 al 4.")
					continue
				}

				for {
					fmt.Print("Ingrese la cantidad: ")
					fmt.Scan(&cantidad)

					if carrito.ValidarCantidad(cantidad) {
						break
					}

					fmt.Println("Cantidad no valida. Intente nuevamente.")
				}

				nombre := productos.ObtenerNombre(codigo)
				precio := productos.ObtenerPrecio(codigo)

				subtotal := carrito.CalcularSubtotal(precio, cantidad)

				subtotales = append(subtotales, subtotal)

				fmt.Println("\nProducto:", nombre)
				fmt.Println("Precio: $", precio)
				fmt.Println("Cantidad:", cantidad)
				fmt.Printf("Subtotal: $%.2f\n", subtotal)

				fmt.Print("\nDesea agregar otro producto? (s/n): ")
				fmt.Scan(&continuar)
			}

			if len(subtotales) > 0 {

				subtotalGeneral := carrito.CalcularTotalCarrito(subtotales)

				iva := compras.CalcularIVA(subtotalGeneral)

				total := compras.CalcularTotal(subtotalGeneral, iva)

				compras.MostrarResumen(subtotalGeneral, iva, total)

			} else {
				fmt.Println("No se agregaron productos.")
			}

		case 3:
			fmt.Println("Gracias por utilizar GoShop.")
			return

		default:
			fmt.Println("Opcion no valida. Intente nuevamente.")
		}
	}
}

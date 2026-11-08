# 🛒 GoShop

### Sistema de Gestión de Comercio Electrónico

GoShop es un proyecto académico desarrollado utilizando el lenguaje de programación Go. Su finalidad es simular las operaciones principales de una tienda virtual mediante una aplicación interactiva ejecutada desde la consola.

## 👨‍💻 Autor

**Alexander Mauricio Macias Medina**

Universidad Internacional del Ecuador

## 🎯 Objetivo

Desarrollar un sistema básico de comercio electrónico que permita consultar productos, seleccionar cantidades y calcular los valores de una compra mediante funciones y paquetes organizados.

## ⚙️ Funcionalidades principales

- Visualización del catálogo de productos.
- Consulta de precios.
- Selección de productos.
- Validación de códigos y cantidades.
- Cálculo de subtotales.
- Cálculo del IVA.
- Cálculo del total de compra.
- Presentación del resumen de compra.

## 📁 Estructura del proyecto

```text
GoShop/
├── carrito/
│   └── carrito.go
├── compras/
│   └── compras.go
├── productos/
│   └── productos.go
├── go.mod
├── main.go
└── README.md
```

## 🧩 Descripción de los módulos

**Productos:** administra el catálogo y permite consultar los nombres y precios de los artículos.

**Carrito:** valida las cantidades y calcula los subtotales de los productos seleccionados.

**Compras:** calcula el IVA y el importe total, además de presentar el resumen de la compra.

**Main:** coordina los módulos y presenta el menú principal del sistema.

## 🛠️ Tecnologías utilizadas

- Lenguaje Go.
- GoLand.
- Git.
- GitHub.
- Biblioteca estándar fmt.

## ▶️ Ejecución del sistema

Para ejecutar el proyecto, abre una terminal en la carpeta principal y utiliza:

```bash
go run .
```

## 📚 Contenidos académicos aplicados

- Sintaxis del lenguaje Go.
- Estructuras condicionales.
- Ciclos iterativos.
- Creación y utilización de funciones.
- Parámetros y valores de retorno.
- Importación y organización de paquetes.
- Programación funcional mediante funciones reutilizables.

## 📌 Alcance

GoShop es una simulación académica de comercio electrónico. Esta versión funciona desde la consola y no incluye base de datos, pagos electrónicos ni gestión de inventario persistente.

---

**Proyecto académico – Programación con Go**
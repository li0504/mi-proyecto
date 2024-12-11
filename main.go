package main

import (
	"fmt"
	"sistema_libros/base"
	"sistema_libros/datos"
	"sistema_libros/funciones"
)

func main() {
	biblioteca := base.BibliotecaElectronica{}
	funciones.SaludarUsuario()

	for {
		fmt.Println("\n=== Menú Principal ===")
		fmt.Println("1. Agregar libro electrónico")
		fmt.Println("2. Listar libros electrónicos")
		fmt.Println("3. Salir")

		var opcion int
		fmt.Print("Ingrese una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var id int
			var titulo, autor, link string

			fmt.Print("Ingrese el ID del libro: ")
			fmt.Scanln(&id)
			fmt.Print("Ingrese el título del libro: ")
			fmt.Scanln(&titulo)
			fmt.Print("Ingrese el autor del libro: ")
			fmt.Scanln(&autor)
			fmt.Print("Ingrese el link del libro: ")
			fmt.Scanln(&link)

			nuevoLibro := datos.Ebook{ID: id, Titulo: titulo, Autor: autor, Link: link}
			biblioteca.AgregarLibro(nuevoLibro)

		case 2:
			biblioteca.ListarLibros()

		case 3:
			fmt.Println("Saliendo del sistema... ¡Hasta luego!")
			return

		default:
			fmt.Println("Opción no válida, intente de nuevo.")
		}
	}
}
	
package base

import (
	"fmt"
	"sistema_libros/datos"
)

type BibliotecaElectronica struct {
	Libros []datos.Ebook
}

// Agregar un libro electrónico
func (b *BibliotecaElectronica) AgregarLibro(libro datos.Ebook) {
	b.Libros = append(b.Libros, libro)
	fmt.Println("¡Libro electrónico agregado con éxito!")
}

// Listar libros electrónicos
func (b *BibliotecaElectronica) ListarLibros() {
	if len(b.Libros) == 0 {
		fmt.Println("No hay libros electrónicos registrados.")
		return
	}

	fmt.Println("\n=== Lista de Libros Electrónicos ===")
	for _, libro := range b.Libros {
		fmt.Printf("ID: %d | Título: %s | Autor: %s | Link: %s\n", libro.ID, libro.Titulo, libro.Autor, libro.Link)
	}
}

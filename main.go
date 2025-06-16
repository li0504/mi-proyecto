// ==============================================
// SISTEMA DE GESTIÓN DE E-COMMERCE
// Objetivo: Automatizar procesos de ventas en línea
// autor: Lirio David Villon Perez
// Universidad Internacional del ecuador
// ==============================================

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// =======================
// INTERFACES GENERALES
// =======================
type CRUD interface {
	Create()
	Read()
	Update()
	Delete()
}

// =======================
// ESTRUCTURAS
// =======================
type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Stock       int
	CategoryID  int
}

type Category struct {
	ID   int
	Name string
}

type Customer struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type Order struct {
	ID         int
	CustomerID int
	Products   []Product
	Total      float64
}

// =======================
// GESTORES (MANAGERS)
// =======================
type ProductManager struct {
	Products []Product
	nextID   int
}

func (pm *ProductManager) Create() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nombre del producto: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Descripción: ")
	desc, _ := reader.ReadString('\n')
	desc = strings.TrimSpace(desc)

	fmt.Print("Precio: ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	price, _ := strconv.ParseFloat(priceStr, 64)

	fmt.Print("Stock: ")
	stockStr, _ := reader.ReadString('\n')
	stockStr = strings.TrimSpace(stockStr)
	stock, _ := strconv.Atoi(stockStr)

	product := Product{
		ID:          pm.nextID,
		Name:        name,
		Description: desc,
		Price:       price,
		Stock:       stock,
	}
	pm.Products = append(pm.Products, product)
	pm.nextID++
	fmt.Println("Producto creado con éxito!")
}

func (pm *ProductManager) Read() {
	fmt.Println("\n=== LISTA DE PRODUCTOS ===")
	for _, p := range pm.Products {
		fmt.Printf("ID: %d, Nombre: %s, Precio: %.2f, Stock: %d\n", p.ID, p.Name, p.Price, p.Stock)
	}
}

func (pm *ProductManager) Update() {
	fmt.Println("Funcionalidad de actualización no implementada aún")
}

func (pm *ProductManager) Delete() {
	fmt.Println("Funcionalidad de eliminación no implementada aún")
}

// =======================
// MAIN
// =======================
func main() {
	fmt.Println("==== SISTEMA DE E-COMMERCE INICIADO ====")

	pm := ProductManager{nextID: 1}

	for {
		fmt.Println("\nMenú Principal:")
		fmt.Println("1. Crear producto")
		fmt.Println("2. Ver productos")
		fmt.Println("0. Salir")
		fmt.Print("Seleccione una opción: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			pm.Create()
		case "2":
			pm.Read()
		case "0":
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}

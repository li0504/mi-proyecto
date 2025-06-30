package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ProductManager struct {
	products []Product
	nextID   int
}

func (pm *ProductManager) Create() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nombre del producto: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Descripción: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("Precio: ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		fmt.Println("Precio inválido")
		return
	}

	fmt.Print("Stock: ")
	stockStr, _ := reader.ReadString('\n')
	stockStr = strings.TrimSpace(stockStr)
	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		fmt.Println("Stock inválido")
		return
	}

	product := Product{
		ID:          pm.nextID,
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	pm.products = append(pm.products, product)
	pm.nextID++

	fmt.Println("Producto creado exitosamente")
}

func (pm *ProductManager) Read() {
	if len(pm.products) == 0 {
		fmt.Println("No hay productos registrados")
		return
	}
	fmt.Println("Lista de productos:")
	for _, p := range pm.products {
		fmt.Printf("ID: %d | Nombre: %s | Descripción: %s | Precio: %.2f | Stock: %d\n", p.ID, p.Name, p.Description, p.Price, p.Stock)
	}
}

func (pm *ProductManager) Update() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese ID del producto a actualizar: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID inválido")
		return
	}

	index := -1
	for i, p := range pm.products {
		if p.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Producto no encontrado")
		return
	}

	fmt.Print("Nuevo nombre (dejar vacío para no cambiar): ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name != "" {
		pm.products[index].Name = name
	}

	fmt.Print("Nueva descripción (dejar vacío para no cambiar): ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	if description != "" {
		pm.products[index].Description = description
	}

	fmt.Print("Nuevo precio (dejar vacío para no cambiar): ")
	priceStr, _ := reader.ReadString('\n')
	priceStr = strings.TrimSpace(priceStr)
	if priceStr != "" {
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			fmt.Println("Precio inválido")
			return
		}
		pm.products[index].Price = price
	}

	fmt.Print("Nuevo stock (dejar vacío para no cambiar): ")
	stockStr, _ := reader.ReadString('\n')
	stockStr = strings.TrimSpace(stockStr)
	if stockStr != "" {
		stock, err := strconv.Atoi(stockStr)
		if err != nil {
			fmt.Println("Stock inválido")
			return
		}
		pm.products[index].Stock = stock
	}

	fmt.Println("Producto actualizado exitosamente")
}

func (pm *ProductManager) Delete() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese ID del producto a eliminar: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID inválido")
		return
	}

	index := -1
	for i, p := range pm.products {
		if p.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Producto no encontrado")
		return
	}

	pm.products = append(pm.products[:index], pm.products[index+1:]...)
	fmt.Println("Producto eliminado exitosamente")
}

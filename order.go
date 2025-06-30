package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type OrderManager struct {
	orders []Order
	PM     *ProductManager
	CM     *CustomerManager
	nextID int
}

func (om *OrderManager) Create() {
	reader := bufio.NewReader(os.Stdin)

	if len(om.CM.customers) == 0 {
		fmt.Println("No hay clientes para hacer un pedido")
		return
	}
	if len(om.PM.products) == 0 {
		fmt.Println("No hay productos para hacer un pedido")
		return
	}

	fmt.Print("Ingrese ID del cliente: ")
	custIDStr, _ := reader.ReadString('\n')
	custIDStr = strings.TrimSpace(custIDStr)
	custID, err := strconv.Atoi(custIDStr)
	if err != nil {
		fmt.Println("ID inválido")
		return
	}

	clientExists := false
	for _, c := range om.CM.customers {
		if c.ID == custID {
			clientExists = true
			break
		}
	}
	if !clientExists {
		fmt.Println("Cliente no encontrado")
		return
	}

	fmt.Println("Ingrese IDs de productos separados por coma (ejemplo: 1,2,3):")
	productsStr, _ := reader.ReadString('\n')
	productsStr = strings.TrimSpace(productsStr)
	prodIDsStr := strings.Split(productsStr, ",")

	var prodIDs []int
	for _, idStr := range prodIDsStr {
		idStr = strings.TrimSpace(idStr)
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Printf("ID de producto inválido: %s\n", idStr)
			return
		}
		prodExists := false
		for _, p := range om.PM.products {
			if p.ID == id {
				prodExists = true
				break
			}
		}
		if !prodExists {
			fmt.Printf("Producto con ID %d no encontrado\n", id)
			return
		}
		prodIDs = append(prodIDs, id)
	}

	order := Order{
		ID:         om.nextID,
		CustomerID: custID,
		ProductIDs: prodIDs,
	}

	om.orders = append(om.orders, order)
	om.nextID++

	fmt.Println("Pedido creado exitosamente")
}

func (om *OrderManager) Read() {
	if len(om.orders) == 0 {
		fmt.Println("No hay pedidos registrados")
		return
	}

	fmt.Println("Lista de pedidos:")
	for _, o := range om.orders {
		fmt.Printf("ID Pedido: %d | Cliente ID: %d | Productos IDs: %v\n", o.ID, o.CustomerID, o.ProductIDs)
	}
}

func (om *OrderManager) Delete() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese ID del pedido a eliminar: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID inválido")
		return
	}

	index := -1
	for i, o := range om.orders {
		if o.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Pedido no encontrado")
		return
	}

	om.orders = append(om.orders[:index], om.orders[index+1:]...)
	fmt.Println("Pedido eliminado exitosamente")
}

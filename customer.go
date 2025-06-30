package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CustomerManager struct {
	customers []Customer
	nextID    int
}

func (cm *CustomerManager) Create() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nombre del cliente: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Teléfono: ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)

	customer := Customer{
		ID:    cm.nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}

	cm.customers = append(cm.customers, customer)
	cm.nextID++

	fmt.Println("Cliente creado exitosamente")
}

func (cm *CustomerManager) Read() {
	if len(cm.customers) == 0 {
		fmt.Println("No hay clientes registrados")
		return
	}
	fmt.Println("Lista de clientes:")
	for _, c := range cm.customers {
		fmt.Printf("ID: %d | Nombre: %s | Email: %s | Teléfono: %s\n", c.ID, c.Name, c.Email, c.Phone)
	}
}

func (cm *CustomerManager) Update() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese ID del cliente a actualizar: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID inválido")
		return
	}

	index := -1
	for i, c := range cm.customers {
		if c.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Cliente no encontrado")
		return
	}

	fmt.Print("Nuevo nombre (dejar vacío para no cambiar): ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name != "" {
		cm.customers[index].Name = name
	}

	fmt.Print("Nuevo email (dejar vacío para no cambiar): ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	if email != "" {
		cm.customers[index].Email = email
	}

	fmt.Print("Nuevo teléfono (dejar vacío para no cambiar): ")
	phone, _ := reader.ReadString('\n')
	phone = strings.TrimSpace(phone)
	if phone != "" {
		cm.customers[index].Phone = phone
	}

	fmt.Println("Cliente actualizado exitosamente")
}

func (cm *CustomerManager) Delete() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese ID del cliente a eliminar: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("ID inválido")
		return
	}

	index := -1
	for i, c := range cm.customers {
		if c.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Cliente no encontrado")
		return
	}

	cm.customers = append(cm.customers[:index], cm.customers[index+1:]...)
	fmt.Println("Cliente eliminado exitosamente")
}

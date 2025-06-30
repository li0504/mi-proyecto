package main

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Stock       int
}

type Customer struct {
	ID    int
	Name  string
	Email string
	Phone string
}

type Order struct {
	ID         int
	CustomerID int
	ProductIDs []int
}

package main

import (
	"database/sql"
	"fmt"
)

// Customer struct
type Customer struct {
	CustomerID   int
	CustomerName string
	SSN          string
}

// GetConnection method which returns sql.DB
func GetConnection() (database *sql.DB) {
	dbDriver := "mysql"
	dbUser := "admindb"
	dbPass := "Test12345"
	dbName := "crm"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(error.Error(err))
	}
	return db
}

// GetCustomers method returns Customer array
func GetCustomers() []Customer {
	var db *sql.DB
	db = GetConnection()

	var error error
	var rows *sql.Rows
	rows, error = db.Query("SELECT * FROM Customer ORDER BY Customerid DESC")
	if error != nil {
		panic(error.Error())
	}
	customer := Customer{}

	customers := []Customer{}
	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		error = rows.Scan(&customerId, &customerName, &ssn)
		if error != nil {
			panic(error.Error())
		}
		customer.CustomerID = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}

	defer db.Close()
	return customers
}

// InsertCustomer method with parameter customer
func InsertCustomer(cust Customer) {
	var database *sql.DB
	database = GetConnection()

	var error error
	var insert *sql.Stmt
	insert, error = database.Prepare("INSERT INTO CUSTOMER(CustomerName, SSN) VALUES(?,?)")
	if error != nil {
		panic(error.Error())
	}
	insert.Exec(cust.CustomerName, cust.SSN)

	defer database.Close()
}

// main method
func main() {
	var customers []Customer

	customers = GetCustomers()
	fmt.Println("Before Insert", customers)
	var customer Customer
	customer.CustomerName = "Arnie Smith"
	customer.SSN = "2345679"
	InsertCustomer(customer)
	customers = GetCustomers()
	fmt.Println("After Insert", customers)

}

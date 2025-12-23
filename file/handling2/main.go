package main

import (
	"fmt"
	"os"
)

type Product struct {
	Name        string
	Price       float64
	Description string
}

func writeProductsToFile(products []Product, filename string) error {
	// Create a new file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write header to file
	header := "PRODUCT NAME, PRICE, DESCRIPTION\n"
	_, err = file.WriteString(header)
	if err != nil {
		return err
	}

	// Loop over products and write each one to file
	for _, product := range products {
		line := fmt.Sprintf("%s, %.2f, %s\n", product.Name, product.Price, product.Description)
		_, err = file.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// Create some sample products
	products := []Product{
		{Name: "Laptop", Price: 849.99, Description: "15-inch screen, 8GB RAM, 512GB SSD"},
		{Name: "Smartphone", Price: 599.99, Description: "6.1-inch screen, 128GB storage, dual camera"},
		{Name: "Headphones", Price: 199.99, Description: "Noise-cancelling, Over-ear, Wireless"},
	}

	// Write products to file
	err := writeProductsToFile(products, "products.txt")
	if err != nil {
		fmt.Println("Error in writing to file:", err)
		return
	}
	fmt.Println("File writing is done successfully.")
}

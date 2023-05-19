package main

import (
	"errors"
	"fmt"
	"log"
	"padaria/src/core/domain"
)

var products []domain.Product

func main() {
	for {
		showMenu()
		option := catchOption()
		treatOption(option)
	}
}

func showMenu() {
	fmt.Println("=========== MAIN MENU ==========")
	fmt.Println("(1) - Register Product")
	fmt.Println("(2) - Edit Product")
	fmt.Println("(3) - Sell Product")
	fmt.Println("(4) - List Of Products")
	fmt.Println("(5) - Leave System")
	fmt.Println("Insert the number of option: ")
}

func catchOption() int {
	var option int
	_, err := fmt.Scanln(&option)
	if err != nil {
		log.Println(errors.New("the option is not valid"))
	}

	return option
}

func treatOption(option int) {
	switch option {
	case 1:
		registerProduct()
	case 2:
		editProduct()
	case 3:
		sellProduct()
	case 4:
		listProducts()
	case 5:
		leaveSystem()
	default:
		fmt.Println("The option is not valid!")
	}
}

func registerProduct() {
	newProduct := formulateProduct()
	products = append(products, *newProduct)
	fmt.Println("Product successfully registered!")
}

func editProduct() {
	var id int
	fmt.Println("Insert the ID of the product: ")
	fmt.Scanln(&id)
	var productEdited domain.Product
	for index := range products {
		if index == id {
			productEdited = *formulateProduct()
		} else {
			fmt.Println("The product does not exist!")
			return
		}
	}

	products = append(products[:id], products[id+1:]...)
	products = append(products, productEdited)

	fmt.Println("Product edited successfully!")
}

func formulateProduct() *domain.Product {
	var (
		id    int
		name  string
		code  string
		price float32
	)
	id = len(products)
	fmt.Print("Insert the name of product: ")
	fmt.Scanln(&name)
	fmt.Print("Insert the code of product: ")
	fmt.Scanln(&code)
	fmt.Print("Insert the price of product: ")
	fmt.Scanln(&price)

	newProduct := domain.NewProduct(id, name, code, price)
	return newProduct
}

func sellProduct() {
	var id int
	fmt.Println("Enter the ID of the product you want to sell: ")
	fmt.Scanln(&id)

	var exists bool
	for index := range products {
		if index == id {
			exists = true
		}
	}
	if exists {
		products = append(products[:id], products[id+1:]...)
		fmt.Println("The product has been successfully sold!")
	} else {
		fmt.Println("The product does not exist in the records!")
	}
}

func listProducts() {
	for _, product := range products {
		fmt.Printf("Product %d: \n", product.ID())
		fmt.Println("Name:", product.Name())
		fmt.Println("Code:", product.Code())
		fmt.Println("Price:$", product.Price())

	}
}

func leaveSystem() {
	panic("Good Bye!")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	//"strings"
)

type user struct {
	username string
	password string
}

type product struct {
	id int
	name string
	stock int
}

type order struct {
	id int
	name string
	quantity int
}

var products []product
var isLogin = false

func main()  {
	var users []user

	for !isLogin {
		fmt.Println("================== Welcome to Ditenun Console Application ==================")
		fmt.Println("Welcome to Ditenun Console Application")
		fmt.Println("List Action")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")

		action := 0

		fmt.Printf("Enter action : ")
		_, _ = fmt.Scanf("%d", &action)

		if action == 0 {
			fmt.Println("Thanks for using our Console Application")
			break
		}else if action == 1 {
			isLogin = login(users)
		}else if action == 2 {
			username := ""
			password := ""
			fmt.Println("Input your personal info")
			fmt.Printf("Input Username : ")
			fmt.Scanf("%s", &username)
			fmt.Printf("Input Password : ")
			fmt.Scanf("%s", &password)
			register(username, password, &users)
		}
	}

	products = append(products, product{id : 1, name : "Ulos Sadum", stock: 25})
	products = append(products, product{id : 2, name : "Ulos Ragihotang", stock: 15})
	products = append(products, product{id : 3, name : "Laptop ROG SCAR", stock: 5})
	products = append(products, product{id : 4, name : "Iphone X S", stock: 10})

	for isLogin {
		menu()
	}

}

// Menu is using for list all menu that using in console application
func menu()  {
	var choise int

	fmt.Println("================== Selamat Datang di Ditenun Console Application==================")
	fmt.Println("List Menu")
	fmt.Println("1. Show List Product")
	fmt.Println("2. Show Product by ID")
	fmt.Println("3. Add Product")
	fmt.Println("4. Update Product")
	fmt.Println("5. Delete Product")
	fmt.Println("6. Order Product")
	fmt.Println("0. Exit")

	fmt.Printf("ENTER ACTION : ")
	_, _ = fmt.Scanf("%d", &choise)

	switch choise {
	case 0:
		fmt.Println("Thank you for using our console application")
		isLogin = false
		break
	case 1:
		showListProduct(products)
		break
	case 2:
		id := 0
		fmt.Println("Enter ID product : ")
		_, _ = fmt.Scanf("%d", &id)
		showProductById(id, products)
		break
	case 3:
		addProduct(&products)
		break
	case 4:
		id := 0
		fmt.Println("Enter ID product : ")
		_, _ = fmt.Scanf("%d", &id)
		updateProductById(id, &products)
		break
	case 5:
		id := 0
		fmt.Println("Enter ID product : ")
		_, _ = fmt.Scanf("%d", &id)
		deleteProductById(id, &products)
		break
	case 6:
		orderProduct(&products)
		break
	default:
		fmt.Println("Your choise is not valid. Please input again!")
		fmt.Println("Input your choise : ")
		break
	}
}

// showListProduct is using for to show all product whi is listed in array
func showListProduct(products []product){
	for _, p := range products {
		fmt.Println("ID		: ", p.id)
		fmt.Println("Name	: ", p.name)
		fmt.Println("Stock	: ", p.stock)
		fmt.Println()
	}
}

// howProductById is using for showing a product by its ID
func showProductById(id int, products []product)  {
	for _, p := range products {
		if p.id == id {
			fmt.Println("ID		: ", p.id)
			fmt.Println("Name	: ", p.name)
			fmt.Println("Stock	: ", p.stock)
			fmt.Println()
		}
	}
}

// addProduct is using for adding a product to array of products
func addProduct(products *[]product)  {
	var id int
	var name string
	var stock int
	fmt.Println()
	fmt.Printf("Input Product's ID : ")
	_, _ = fmt.Scanf("%d", &id)
	fmt.Printf("Input Product's Name : ")
	consoleReader := bufio.NewReader(os.Stdin)
	name, _ = consoleReader.ReadString('\n')
	fmt.Printf("Input Product's Stock : ")
	_, _= fmt.Scanf("%d", &stock)

	product_ := product{id,name,stock}

	*products = append(*products, product_)
}

// updateProductById is using for update a product's stock by its ID
func updateProductById(id int, products *[]product)  {
	var newStock int
	fmt.Println()
	fmt.Print("Enter the New quantity of the Product: ")
	fmt.Scanf("%d", &newStock)

	for i, p := range *products {
		if p.id == id {
			(*products)[i] = product{id, p.name, newStock}
		}
	}
}

// deleteProductById is using for delete a product by ID
func deleteProductById(id int, products *[]product)  {
	for i, p := range *products{
		if p.id == id {
			*products = append((*products)[:i], (*products)[i+1:]...)
		}
	}
}

// orderProduct is ordering products
func orderProduct(products *[]product)  {
	var idorder int
	var namaorder string
	var id int
	var quantityorder int
	var newStock int

	fmt.Printf("Enter ID Order : ")
	_, _ = fmt.Scanf("%d", &idorder)
	fmt.Printf("Enter Order User's Name : ")
	consoleReader := bufio.NewReader(os.Stdin)
	namaorder, _ = consoleReader.ReadString('\n')
	fmt.Printf("Enter ID Product : ")
	_, _ = fmt.Scanf("%d", &id)
	fmt.Printf("Enter Quantity : ")
	_, _ = fmt.Scanf("%d", &quantityorder)

	for i, p := range *products{
		if p.id == id {
			if p.stock >= 1 {
				if p.stock >= quantityorder {
					newStock = p.stock - quantityorder
					(*products)[i] = product{id, p.name, newStock}
				}
			}
		}
	}

	fmt.Println("=======================================")
	fmt.Println("Berhasil di Order")
	fmt.Printf("ID Order : %d\n", idorder)
	fmt.Printf("ID Product yang di order : %d\n", id)
	fmt.Printf("Name : %s\n", namaorder)
	fmt.Printf("Quantity : %d\n", quantityorder)
}

// login is using for login
func login(users []user) bool {
	var username string
	var password string
	//consoleReader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter Username : ")
	_, _ =fmt.Scanf("%s", &username)
	fmt.Printf("Enter Password : ")
	_, _ =fmt.Scanf("%s", &password)

	if checkCredential(username, password, users){
		return true
	}

	return false
}

//checkCredential is using for compare username input with username in array and password input with password in array
func checkCredential(username string, password string, users []user) bool {
	for _, u := range users {
		if strings.Compare(username, u.username) == 0 && strings.Compare(password, u.password) == 0 {
			return true
		}
	}

	return false
}

// register is using for registering username and password to user array
func register(username string, password string, users *[]user)  {
	*users = append(*users, user{username, password})
}
package main

import "fmt"

func main() {
	fmt.Println("Welcom to Go Bank.")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("What is your choice? ")
	fmt.Scan(&choice)

	

	fmt.Println("Your choice is ", choice)
}
package main

import "fmt"

func main() {
	var accountBalance float64 = 1000


	fmt.Println("Welcom to Go Bank.")
	fmt.Println("What would you like to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("What is your choice? ")
	fmt.Scan(&choice)

	//Using conditions
	// checkBalance := choice == 1

	// if checkBalance {

	// }

	if choice == 1 {
		fmt.Println("Your balance is:", accountBalance)
	}
	fmt.Println("Your choice is ", choice)
}
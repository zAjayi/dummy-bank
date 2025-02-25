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
	} else if choice == 2 {
		fmt.Print("Enter amount to deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Account balance has been updated!\nYour new balance is:", accountBalance)
	} else if choice == 3 {
		fmt.Print("Enter amount to withdraw: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		accountBalance -= withdrawAmount
		fmt.Println("Account balance has been updated!\nYour new balance is:", accountBalance)
	} else {
		fmt.Println("Thank you for using Go Bank!")
	}
}
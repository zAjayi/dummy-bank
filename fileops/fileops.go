package fileops

import (
	"os"
	"fmt"
)

func WriteBalancetoFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
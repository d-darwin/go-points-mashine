package main

import (
	"fmt"
	"os"
	points_machine "points_machine/internal/app/points-machine"
)

func main() {
	// оператор defer выполняется после завершения работы функции
	defer func() {
		if v := recover(); v != nil {
			// bad practice to intercept panics
			fmt.Println("Panic happened", v)
			os.Exit(1)
		}
	}()

	// slice
	accountsSlice := make([]int, 0, 1)
	accountsSlice = append(accountsSlice, 100)
	accountsSlice = append(accountsSlice, 200)
	accountsSlice = append(accountsSlice, 300)
	fmt.Println(accountsSlice)
	fmt.Println(cap(accountsSlice))

	// map
	// var accountsMap map[int]int
	// accountsMap[1] = 100
	points_machine.InitAccount(1)
	points_machine.Accrual(1, 100)
	points_machine.Charge(1, 50)
	// accountsMap[2] = 200
	points_machine.InitAccount(2)
	points_machine.Accrual(2, 200)
	points_machine.Charge(2, 100)
	// accountsMap[3] = 300
	// initAccount(2)

	points_machine.PrintAccounts()
	points_machine.PrintTransactions()
}

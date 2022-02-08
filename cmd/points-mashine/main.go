package main

import (
	"fmt"
	"os"
	points_mashine "pointsmashine/internal/app/points-mashine"
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
	points_mashine.InitAccount(1)
	points_mashine.Accrual(1, 100)
	points_mashine.Charge(1, 50)
	// accountsMap[2] = 200
	points_mashine.InitAccount(2)
	points_mashine.Accrual(2, 200)
	points_mashine.Charge(2, 100)
	// accountsMap[3] = 300
	// initAccount(2)

	fmt.Println(accountsMap)
	points_mashine.PrintTransactions()
}

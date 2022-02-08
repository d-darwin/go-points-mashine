package points_mashine

import (
	"fmt"
)

var accountsMap = make(map[int]uint) // bad practice

type transactionType uint8

const (
	AccrualType transactionType = iota + 1
	ChargeType
)

type transaction struct {
	clientId        int
	transactionType transactionType
	amount          uint
}

var transactions []*transaction

func InitAccount(id int) {
	amount, ok := accountsMap[id]
	if ok {
		panic("Account already exists") // bad practice to create panics
	}
	accountsMap[id] = amount
}

func PrintAccounts() {
	fmt.Println(accountsMap)
}

func PrintTransactions() {
	fmt.Println("Transactions:")
	for _, trx := range transactions {
		fmt.Println(*trx)
	}
}

func Accrual(id int, amount uint) {
	if _, ok := accountsMap[id]; !ok {
		panic("Account doesn't exists") // bad practice to create panics
	}

	accountsMap[id] += amount

	trx := transaction{
		clientId:        id,
		transactionType: AccrualType,
		amount:          amount,
	}

	transactions = append(transactions, &trx)
}

func Charge(id int, amount uint) {
	currentAmount, ok := accountsMap[id]

	if !ok {
		panic("Account doesn't exists") // bad practice to create panics
	}

	if currentAmount < amount {
		panic("Not enough points")
	}

	accountsMap[id] -= amount

	trx := transaction{
		clientId:        id,
		transactionType: ChargeType,
		amount:          amount,
	}

	transactions = append(transactions, &trx)
}

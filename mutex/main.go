package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

var wg sync.WaitGroup

func main() {

	//variable for bankbalance
	var bankBalance int
	var balanceMutex sync.Mutex //declare mutex for bankBalance variable
	//print starting balance
	fmt.Printf("Initial bank balance: $%d.00\n", bankBalance)

	//define weekly income for each source
	incomes := []Income{
		{
			Source: "Main Job",
			Amount: 500,
		},
		{
			Source: "Part Time Job",
			Amount: 50,
		},
		{
			Source: "Gifts",
			Amount: 10,
		},
		{
			Source: "Investments",
			Amount: 100,
		},
	}

	wg.Add(len(incomes))
	//calculate income for whole year
	for _, income := range incomes {
		go calculateYearlyBalance(income, &bankBalance, &balanceMutex)
	}

	wg.Wait()

	//print final total balance
	fmt.Printf("Final bank balance: $%d.00\n", bankBalance)
}

func calculateYearlyBalance(income Income, bankBalance *int, balanceMutex *sync.Mutex) {
	defer wg.Done()

	for i := 1; i <= 52; i++ {
		balanceMutex.Lock()

		temp := bankBalance
		*temp += income.Amount
		bankBalance = temp

		balanceMutex.Unlock()

		fmt.Printf("On week %d, you earned $%d.00 from %s\n", i, income.Amount, income.Source)
	}
}

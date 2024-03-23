package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NUMBER_OF_PIZZAS = 10

var pizzasMade, pizzasFailed, totalPizzas int

type PizzaOrder struct {
	PizzaNumber int
	Message     string
	Success     bool
}

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= NUMBER_OF_PIZZAS {
		fmt.Printf("Received order number #%d\n", pizzaNumber)

		random := rand.Intn(12) + 1
		msg := ""
		success := false

		if random < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		totalPizzas++

		pizzaMakingTime := rand.Intn(5) + 1

		fmt.Printf("Making Pizza #%d. It will take %d seconds...\n", pizzaNumber, pizzaMakingTime)
		time.Sleep(time.Duration(pizzaMakingTime) * time.Second)

		if random <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if random <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making the pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("The pizza #%d is ready!", pizzaNumber)
		}

		return &PizzaOrder{
			PizzaNumber: pizzaNumber,
			Message:     msg,
			Success:     success,
		}
	}
	return &PizzaOrder{
		PizzaNumber: pizzaNumber,
		Message:     "*** We are Sold out for today!",
		Success:     false,
	}
}

func processor(pizzaJob *Producer) {
	var i = 0

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.PizzaNumber

			select {
			case pizzaJob.data <- *currentPizza:

			case quitChan := <-pizzaJob.quit:
				close(pizzaJob.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())
	//print starting msg
	color.Cyan("The Pizzeria is open for orders!")
	color.Cyan("----------------------------------")

	//create a producer

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run producer in background
	go processor(pizzaJob)
	//create and run consumer
	for i := range pizzaJob.data {
		if i.PizzaNumber <= NUMBER_OF_PIZZAS {
			if i.Success {
				color.Green(i.Message)
				color.Green("Order %d is out for delivery!", i.PizzaNumber)
			} else {
				color.Red(i.Message)
				color.Red("Order %d got cancelled!", i.PizzaNumber)
			}
		} else {
			color.Cyan("Done making pizzas!")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("*** Error Closing channel!", err)
			}
		}
	}
	//print ending msg

	color.Cyan("-----------------")
	color.Cyan("Done for the Day. Pizzeria is closed now!")
	color.Cyan("We made %d pizzas, but failed to make %d, with %d attempts in total!", pizzasMade, pizzasFailed, totalPizzas)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day!")
	case pizzasFailed >= 6:
		color.Red("It was a bad day!")
	case pizzasFailed >= 4:
		color.Yellow("It was an okay day!")
	case pizzasFailed >= 2:
		color.Yellow("It was a good day!")
	default:
		color.Green("It was a great day!")
	}

}

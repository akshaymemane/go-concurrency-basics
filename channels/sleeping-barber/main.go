package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// declare required variables
var seatingCapacity = 10
var arrivalRate = 100 //milliseconds
var haircutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second

func main() {
	//seed random number generator
	rand.Seed(time.Now().UnixNano())

	//print welcome message
	color.Yellow("The Sleeping Barber Problem!")
	color.Yellow("----------------------------")

	//create required channels
	clientsChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	//create the barbershop
	shop := Barbershop{
		ShopCapacity:    seatingCapacity,
		HaircutDuration: haircutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientsChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}

	//add the barbers
	shop.addBarber("Frank")
	shop.addBarber("Curt")
	shop.addBarber("John")
	shop.addBarber("Mark")
	shop.addBarber("Suzan")

	//start the barbershop as go routine
	shopClosing := make(chan bool)
	shopClosed := make(chan bool)

	go func() {
		color.Green("Shop is open for the day!")

		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForDay()
		shopClosed <- true
	}()

	//add clients
	i := 1

	go func() {

		for {
			random := rand.Int() % (2 * arrivalRate)
			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(random)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	//block until barbershop is closed
	<-shopClosed
}

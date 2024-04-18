package main

import (
	"time"

	"github.com/fatih/color"
)

type Barbershop struct {
	ShopCapacity    int
	HaircutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *Barbershop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			if len(shop.ClientsChan) == 0 {
				color.Yellow("There is nothing to do. so %s takes a nap!", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}
				//cut hair
				shop.cutHair(barber, client)
			} else {
				//shop is closed, send the barber home and close this goroutine
				shop.sendBarberHome(barber)
				return
			}
		}
	}()
}

func (shop *Barbershop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.HaircutDuration)
	color.Green("%s finished cutting %s's hair.", barber, client)
}

func (shop *Barbershop) sendBarberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	shop.BarbersDoneChan <- true
}

func (shop *Barbershop) closeShopForDay() {
	color.Cyan("Closing shop for the day!")

	close(shop.ClientsChan)
	shop.Open = false

	for i := 1; i <= shop.NumberOfBarbers; i++ {
		<-shop.BarbersDoneChan
	}

	close(shop.BarbersDoneChan)

	color.Green("------------------------------------------------------------------")
	color.Green("The barbershop is now closed for the day, every one has gone home!")
}

func (shop *Barbershop) addClient(client string) {
	color.Green("*** %s arrives!", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("%s takes the seat in the waiting room.", client)
		default:
			color.Red("The waiting room is full, so %s leaves.", client)
		}
	} else {
		color.Red("The shop is already closed, so %s leaves.", client)
	}
}

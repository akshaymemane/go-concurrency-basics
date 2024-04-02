package main

import (
	"fmt"
	"strings"
)

func shout(ping, pong chan string) {

	for {
		str := <-ping

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(str))
	}

}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to EXIT!)")

	for {
		fmt.Print("-> ")

		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput

		response := <-pong
		fmt.Println("Response: ", response)
	}

	fmt.Println("All Done. closing channels!")
	close(ping)
	close(pong)
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"time"
)

// type Record struct{
// 	Ticker string
// 	Date string
// 	Time string
// 	LTP float64
// 	BuyPrice float64
// 	BuyQty int
// 	SellPrice float64
// 	SellQty float64
// 	LTQ float64
// 	OpenInterest float64
// }

func produce(ch chan string) {
	for {
		str := <-ch

		fmt.Printf("%s\n", str)

		//pretend to work on some task for a second
		time.Sleep(100 * time.Millisecond)
	}
}

func ReadLines(path string, ch chan string) error {
	file, err := os.Open(path)

	if err != nil {
		return errors.New("failed to open the file!")
	}

	scanner := bufio.NewScanner(file)
	// var lines []string
	for scanner.Scan() {
		// lines = append(lines, scanner.Text())
		// words := strings.Split(scanner.Text(), ",")
		ch <- scanner.Text()
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return errors.New("failed to read the lines in the file!")
	}

	file.Close()
	return nil
}

func main() {

	ch := make(chan string)

	go produce(ch)

	ReadLines("NIFTY BANK.NSE_IDX copy 1316.csv", ch)

}

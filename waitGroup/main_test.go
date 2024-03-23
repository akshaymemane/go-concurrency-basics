package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_printMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("test message", &wg)
	wg.Wait()

	if msg != "test message" {
		t.Error("Expected test message, but not found!")
	}

}

func Test_updatewMessage(t *testing.T) {

	stdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w
	msg = "test message"
	printMessage()

	w.Close()
	result, _ := io.ReadAll(r)
	outout := string(result)

	os.Stdout = stdout

	if !strings.Contains(outout, "test message") {
		t.Error("Expected test message, but not found!")
	}
}

func Test_main(t *testing.T) {
	stdout := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	w.Close()
	result, _ := io.ReadAll(r)
	outout := string(result)

	os.Stdout = stdout

	if !strings.Contains(outout, "Hello, Universe!") {
		t.Error("Expected Hello, Universe!, but not found!")
	}

	if !strings.Contains(outout, "Hello, Cosmos!") {
		t.Error("Expected Hello, Cosmos!, but not found!")
	}

	if !strings.Contains(outout, "Hello, World!") {
		t.Error("Expected Hello, World!, but not found!")
	}
}

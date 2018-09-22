package main

import (
	"awesomeProject/module"
	"github.com/stianeikeland/go-rpio"
	"fmt"
	"os"
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	//logic
	var led module.LedTube8Digits
	led.C595.DIO = rpio.Pin(17)
	led.C595.RCK = rpio.Pin(27)
	led.C595.SCK = rpio.Pin(22)

	led.NumberToShow = 12345678
	for {
		led.Show()
	}
}

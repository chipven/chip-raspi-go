package main

import (
	"github.com/stianeikeland/go-rpio"
	"fmt"
	"os"
	"flag"
	"net/http"
	"awesomeProject/module"
	"strconv"
)

var led module.LedTube8Digits

func ledController(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var number []string = r.Form["number"]
	var num = number[0]
	led.NumberToShow, _ = strconv.Atoi(num)
}

func main() {
	println("use like:\n    showLed -number=12345678 \n")

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()

	//logic
	led.DIO = rpio.Pin(17)
	led.RCK = rpio.Pin(27)
	led.SCK = rpio.Pin(22)

	var number = flag.Int("number", 0, "number to show")
	flag.Parse()

	led.NumberToShow = *number

	http.HandleFunc("/", ledController)
	go http.ListenAndServe(":8080", nil)

	for {
		led.Show()
	}

}

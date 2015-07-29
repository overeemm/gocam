package main

import (
	"fmt"
	"time"
	"github.com/stianeikeland/go-rpio"	
)

func toggle(pinNr int) {
	pin := rpio.Pin(pinNr)
	pin.Output()
	pin.High()
	time.Sleep(500 * time.Millisecond)
	pin.Low()
}

func main() {
	err := rpio.Open()
	if(err != nil) {
		fmt.Println("rpio.Open()")
                fmt.Println(err)
        }      
	// Unmap gpio memory when done
	defer rpio.Close()

	for i := 0; i < 200; i++ {
		toggle(9)
		toggle(10)
		toggle(11)
	}
}

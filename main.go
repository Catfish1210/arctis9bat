package main

import (
	"fmt"

	"github.com/Catfish1210/arctis9bat/core"
)

func main() {
	fmt.Println("connected devices to system USB buses:")
	lsusbData, err := core.RunLSUSB()
	if err != nil {
		fmt.Printf("Error running lsusb: %v", err)
		return
	}
	fmt.Println(lsusbData)

}

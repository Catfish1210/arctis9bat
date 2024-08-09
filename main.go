package main

import (
	"fmt"

	"github.com/Catfish1210/arctis9bat/arctis9"
)

func main() {
	Arctis9 := arctis9.Arctis9
	Arctis9.Init()
	if Arctis9.Error != nil {
		fmt.Printf("error: %v\n", Arctis9.Error)
		return
	}
	Arctis9.GetBattery()
	if Arctis9.Error != nil {
		fmt.Printf("error: %v\n", Arctis9.Error)
		return
	}
	fmt.Println("------------------------------------")
	fmt.Printf("Battery Status: [%v]\n", Arctis9.Battery.Status)
	fmt.Printf("Battery  Level:        [%v]\n", Arctis9.Battery.Level)
	fmt.Println("------------------------------------")
}

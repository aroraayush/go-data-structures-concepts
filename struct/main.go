// Go doesn't have the concept of classes
// It uses structs to create objects and do OOP

package main

import (
	"fmt"
	"unsafe"
)

type Car struct {
	gasPedal      uint16 //  0  to 2^16 - 1 (65,535)
	brakePedal    uint16
	steeringWheel int16   // -32k to 32k - 1
	topSpeedKmph  float64 // for max precision
}

func main() {

	gasPedalVal := uint16(2234)
	brakePedalVal := uint16(0)
	steeringWheelVal := int16(12561)
	topSpeedKmphVal := 225.0

	// Calling like a constructor
	var car1 Car = Car{gasPedal: 2234, brakePedal: 0, steeringWheel: 12561, topSpeedKmph: 225.0}
	fmt.Println(car1) // {2234 0 12561 225}

	// {2234 0 12561 225} => Type: main.Car | Value: {2234 0 12561 225} | Size: 16 bits
	fmt.Printf("%v => Type: %T | Value: %v | Size: %d bits\n\n", car1, car1, car1, unsafe.Sizeof(car1))

	// Without var, with short declaration/walrus operator (:=)
	car2 := Car{gasPedal: 2234, brakePedal: 0, steeringWheel: 12561, topSpeedKmph: 225.0}
	fmt.Println(car2) // {2234 0 12561 225}

	car3 := Car{gasPedalVal, brakePedalVal, steeringWheelVal, topSpeedKmphVal}
	fmt.Println(car3) // {2234 0 12561 225}

	car_formatted := Car{
		gasPedal:      2234,
		brakePedal:    0,
		steeringWheel: 12561,
		topSpeedKmph:  225.0, // extra comma required at last
	}
	fmt.Println(car_formatted) // {2234 0 12561 225}
}

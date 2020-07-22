// Value Receivers
//

package main

import (
	"fmt"
)

const USIXTEENBITMAX = 65535
const KMH_MULTIPLE = 1.60934 // 1 mph = 1.6 kmph

type Car struct {
	gasPedal      uint16 //  0  to 2^16 - 1 (65,535)
	brakePedal    uint16
	steeringWheel int16   // -32k to 32k - 1
	topSpeedKmph  float64 // for max precision
}

// Adding function to struct
func (c Car) kmph() float64 { //  return type is float64
	return float64(c.gasPedal) * (c.topSpeedKmph / USIXTEENBITMAX)
}

func (c Car) mph() float64 { //  return type is float64
	return float64(c.gasPedal) * (c.topSpeedKmph / (USIXTEENBITMAX * KMH_MULTIPLE))
}

func main() {

	// Calling like a constructor
	car := Car{gasPedal: 65000, brakePedal: 0, steeringWheel: 12561, topSpeedKmph: 225.0}
	fmt.Println(car, "\n") // {2234 0 12561 225}

	fmt.Println("Current Postion", car.gasPedal)
	fmt.Println("kmph", car.kmph()) // 7.669947356374457
	fmt.Println("mph", car.mph())   // 7.669947356374457
}

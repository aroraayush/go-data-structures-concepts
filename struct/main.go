// Go doesn't have the concept of classes
// It uses structs to create objects and do OOP

package main

type car struct {
	gasPedal      uint16 //  0  to 2^16 - 1 (65,535)
	brakePedal    uint16
	steeringWheel int16   // -32k to 32k - 1
	topSpeedKmph  float64 // for max precision
}

func main() {
	var car1 car = {
		gasPedal: 2234,
		brakePedal: 1,
		steeringWheel: 1,
		topSpeedKmph: 1
	}

	car2 := car{
		gasPedal: 2234,
		brakePedal: 1,
		steeringWheel: 1,
		topSpeedKmph: 1
	}

	

	
	car3 := car{
		gasPedal: 1,
		brakePedal: 1,
		steeringWheel: 1,
		topSpeedKmph: 1
	}

	
}

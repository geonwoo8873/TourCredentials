package main

import "fmt"

// Object N.1 : time is not over 52 hours
// Object N.2 : percent is not over 100%
// Object N.3 : return the total salary of the $1M employee
func Calculate(weekTime int, ratePercent float64) float64 {
	if (weekTime >= 0 && weekTime <= 52) && (ratePercent >= 0 && ratePercent <= 100) {
		return float64(weekTime) * ((1 + ratePercent/100) * 10000)
	}
	return 0
}

func main() {
	fmt.Println(Calculate(2, 2.5)) // Output : 20500
}

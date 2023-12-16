package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	// var investmentAmount, years float64 = 1000, 10
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Year(s): ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculationFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value is %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value is %.2f", futureRealValue)

	fmt.Printf(formattedFV, formattedRFV)
}

func calculationFutureValues(investmentAmount, expectedReturnRate, years float64) (futureValue, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealValue
}

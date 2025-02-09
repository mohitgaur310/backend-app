package main

import (
	"fmt"
	"math"
)
func main() {
	var investMentAmt float64 
	const inflationRate = 2.5
	expectedReturnRate :=5.5;
	var years float64 =10
	fmt.Print(("please enter investment value :"))
	fmt.Scan(&investMentAmt)

	fmt.Print(("please enter Year : "))
	fmt.Scan(&years)
	fmt.Print(("please enter Expected Return Rate : "))
	fmt.Scan(&expectedReturnRate)
	var futureValue =investMentAmt* math.Pow(1+ expectedReturnRate/100,years)
	futureRealValue := futureValue/math.Pow(1+inflationRate/100,years)
	fmt.Println(futureValue,futureRealValue)
}
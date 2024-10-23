package main

import (
	"fmt"
	"math"
)
func main() {
	var investMentAmt float64 =1000;
	const inflationRate = 2.5
	expectedReturnRate :=5.5;
	var years float64 =10
	var futureValue =investMentAmt* math.Pow(1+ expectedReturnRate/100,years)
	futureRealValue := futureValue/math.Pow(1+inflationRate/100,years)
	fmt.Println(futureValue,futureRealValue)
}
package main

import (
	"fmt"
	"math"
)
func main() {
	var investMentAmt=1000;
	var expectedReturnRate=5.5;
	var years =10
	var futureValue =float64(investMentAmt)* math.Pow(1+ expectedReturnRate/100,float64(years))
	fmt.Println(futureValue)
}
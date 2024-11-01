package main

import (
	"fmt"
)
func main() {


	revenue:= userInput("Revenue :")
	expenses:= userInput("Expenses :")
	texRate:=userInput("Tex rate :")

	ebt,profile:= calculateFutureValue(revenue,expenses,texRate)
	ratioString := fmt.Sprintf("the ratio of the ebt and profit is  %.2f",ebt /profile) 

	fmt.Println(ebt,profile,ratioString)
}

func outputText(text string){
	fmt.Println(text)
}

func calculateFutureValue(revenue float64,expenses float64,texRate float64) (float64,float64) {
	ebt:= revenue-expenses;
	profile := ebt*(1-texRate/100)
	return ebt,profile
}

func userInput(text string)  (value float64 ){
	// outputText("Revenue :") 
	outputText(text)
	fmt.Scan(&value);
	return value
}
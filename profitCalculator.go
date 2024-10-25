package main

import (
	"fmt"
)
func main() {
	var revenue float64 
	var expenses float64 
	var texRate float64 


	fmt.Print(("Revenue :"))
	fmt.Scan(&revenue);
	fmt.Print(("Expenses :"))
	fmt.Scan(&expenses);
	fmt.Print(("Tex rate :"))
	fmt.Scan(&texRate);

	ebt:= revenue-expenses;
	profile := ebt*(1-texRate/100)
	ratioString := fmt.Sprintf("the ratio of the ebt and profit is  %.2f",ebt /profile) 




	fmt.Println(ebt,profile,ratioString)
}
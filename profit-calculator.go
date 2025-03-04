package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := float64(revenue - expenses)
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	newEbt := fmt.Sprintln("Ebt : ", ebt)
	newProfit := fmt.Sprintln("Profit : ", profit)
	newRatio := fmt.Sprintln("Ratio : ", ratio)

	// fmt.Printf("the ebt is: %.0f\nthe profit is : %.0f\nthe ratio is :%.0f\n", ebt, profit, ratio)

	fmt.Println(newEbt, newProfit, newRatio)
}

func outputText(text1 string) {
	fmt.Print(text1)
}

package main

import "fmt"

func CalculateBalance() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	// don't edit above this line

	finalCost = bulkMessageCost
	if isPremiumUser {
		finalCost -= bulkMessageCost * discountRate
	} else {
	}
	if accountBalance >= finalCost {
		accountBalance -= finalCost
		fmt.Println(purchaseSuccessMessage)
	} else {
		fmt.Println(insufficientFundMessage)
	}

	// don't edit below this line

	fmt.Println("Account balance:", accountBalance)
}

func reformat(message string, formatter func(string) string) string {
	// Can be done like this:
	// once := formatter(message)
	// twice := formatter(once)
	// thrice := formatter(twice)
	// prefix := "TEXTIO: "
	// return prefix + thrice
	return "TEXTIO: " + formatter(formatter(formatter(message)))
}

func printReports(intro, body, outro string) {
	printCostReport(func(s string) int {
		return len(s) * 2
	}, intro)

	printCostReport(func(s string) int {
		return len(s) * 3
	}, body)
	
	printCostReport(func(s string) int {
		return len(s) * 4
	}, outro)
}

func printR() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	baseBill := calculateBaseBill(costPerMessage, numMessages)
	discountPercentage := calculateDiscountRate(numMessages)
	discountAmount := baseBill * discountPercentage
	finalBill := baseBill - discountAmount
	return finalBill
}

func calculateDiscountRate(messagesSent int) float64 {
	if messagesSent > 5000 {
		return .2
	}
	if messagesSent > 1000 {
		return .1
	}
	return 0.0
}

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}


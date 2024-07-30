package main

import (
	"fazztrack/golangbasic/calc"
	"fmt"
)
func fazzFood(price int, voucher string, distance int, isTax bool) {
	discount := calc.Discount(price, voucher)
	shippingFee := calc.Shipping(distance)
	taxValue := calc.Tax(price, isTax)
	
	var subtotal int = price - discount + shippingFee + taxValue
	fmt.Printf("Harga		: %d\n", price)
	fmt.Printf("Potongan	: %d\n", discount)
	fmt.Printf("Biaya Antar	: %d\n", shippingFee)
	fmt.Printf("Pajak		: %d\n", taxValue)
	fmt.Printf("Subtotal	: %d", subtotal)
}

func main() {
	fazzFood(75000, "Ditraktir60", 5, true)
}
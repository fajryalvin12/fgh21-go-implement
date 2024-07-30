package mainprocess

import (
	"fazztrack/golangbasic/calc"
	"fmt"
)

func FazzFood(price int, voucher string, distance int, isTax bool) {
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

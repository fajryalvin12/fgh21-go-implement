package main

import "fmt"

func calcDiscount (price int, voucher string) (int) {
	discount := 0 // var discount int = 0
	if voucher == "FazzFood50" {
		if price >= 50000 {
			discount = price * 50/100
			if discount > 50000 {
				discount = 50000
			}
		}
	} else if voucher == "Ditraktir60" {
		if price >= 25000 {
			discount = price * 60/100
			if discount > 30000 {
				discount = 25000
			}
		}
	}
	return discount
}

func calcShipping (distance int) (int) {
	shippingFee := 5000

	if distance > 2 {
		shippingFee = shippingFee + (distance - 2) * 3000 
	}
	return shippingFee
}

func calcTax (price int, isTax bool) (int) {
	taxValue := 0

	if isTax == true {
		taxValue = price * 5/100
	}

	return taxValue
}

func fazzFood(price int, voucher string, distance int, isTax bool) {
	discount := calcDiscount(price, voucher)
	shippingFee := calcShipping(distance)
	taxValue := calcTax(price, isTax)
	
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
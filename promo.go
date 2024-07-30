package main

import "fmt"

func fazzFood(price int, voucher string, distance int, isTax bool) {
	var shippingFee int = 5000
	var discount int = 0
	var taxValue int = 0

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
			if discount > 25000 {
				discount = 25000
			}
		}
	}

	if distance > 2 {
		shippingFee = shippingFee + (distance - 2) * 3000 
	}

	if isTax == true {
		taxValue = price * 5/100
	}

	var subtotal int = price - discount + shippingFee + taxValue
	fmt.Printf("Harga: %d\n", price)
	fmt.Printf("Potongan: %d\n", discount)
	fmt.Printf("Biaya Antar: %d\n", shippingFee)
	fmt.Printf("Pajak: %d\n", taxValue)
	fmt.Printf("Subtotal: %d", subtotal)
}

func main() {
	fazzFood(75000, "FazzFood50", 3, true)
}
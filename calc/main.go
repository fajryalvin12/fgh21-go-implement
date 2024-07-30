package calc

func Discount(price int, voucher string) int {
	discount := 0 // var discount int = 0
	if voucher == "FazzFood50" {
		if price >= 50000 {
			discount = price * 50 / 100
			if discount > 50000 {
				discount = 50000
			}
		}
	} else if voucher == "Ditraktir60" {
		if price >= 25000 {
			discount = price * 60 / 100
			if discount > 30000 {
				discount = 25000
			}
		}
	}
	return discount
}

func Shipping(distance int) int {
	shippingFee := 5000

	if distance > 2 {
		shippingFee = shippingFee + (distance-2)*3000
	}
	return shippingFee
}

func Tax(price int, isTax bool) int {
	taxValue := 0

	if isTax == true {
		taxValue = price * 5 / 100
	}

	return taxValue
}
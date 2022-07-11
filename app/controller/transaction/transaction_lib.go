package transaction

// GetDiscount func
func GetDiscount(amount, disc float64, discType int) *float64 {
	var t float64
	t = amount - disc
	if discType == 1 && disc > 0 {
		t = amount * disc / 100
		t = amount - t
	}
	return &t
}

// GetTax func
func GetTax(amount, tax float64, discType int) *float64 {
	var t float64
	t = tax + amount
	if discType == 1 && tax > 0 {
		t = amount * tax / 100
		t = amount + tax
	}
	return &t
}

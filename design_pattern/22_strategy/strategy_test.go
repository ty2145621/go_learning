package strategy

func ExamplePayByCash() {
	payment := NewPayment("Ada", "", 123, &Cash{})
	payment.Pay()
	// Output:
	// Pay $123 to Ada by cash
}

func ExamplePayByBank() {
	payment := NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()
	// Output:
	// Pay $888 to Bob by bank account 0002
}

func ExamplePay() {
	payment := NewPayment("Bob", "0002", 888, &Bank{})
	payment.Pay()

	payment.setStrategy(&Cash{})
	payment.Pay()
	// Output:
	// Pay $888 to Bob by bank account 0002
	// Pay $888 to Bob by cash
}

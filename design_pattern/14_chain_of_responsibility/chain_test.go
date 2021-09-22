package chain

func ExampleChain() {
	c1 := NewProjectManagerChain()
	c2 := NewDepManagerChain()
	c3 := NewGeneralManagerChain()

	c1.SetSuccessor(c2)
	c2.SetSuccessor(c3)

	var c Manager = c1

	c.HandleFeeRequest("bob", 400)
	c.HandleFeeRequest("tom", 1400)
	c.HandleFeeRequest("ada", 10000)
	c.HandleFeeRequest("floar", 400)
	// Output:
	// Project manager permit bob 400 fee request
	// Dep manager permit tom 1400 fee request
	// General manager permit ada 10000 fee request
	// Project manager don't permit floar 400 fee request

}

func ExampleChain2() {
	c := NewChainManager()

	c1 := NewProjectHandlerChain()
	c2 := NewDepHandlerChain()
	c3 := NewGeneralHandlerChain()

	c.addHandler(c1)
	c.addHandler(c2)
	c.addHandler(c3)

	c.process(chainReq{money: int64(0)})
	// Output:
	// Project Handler change fee request:100
	// Dep Handler change fee request:200
	// General Handler change fee request:300
	// General Handler don't permit 300 fee request
	// General Handler after money:300
	// Dep Handler after money:200
	// Project Handler after money:100
}

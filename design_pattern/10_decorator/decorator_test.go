package decorator

import "fmt"

func ExampleDecorator() {
	var c Component = &ConcreteComponent{}
	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)
	res := c.Calc()

	fmt.Printf("res %d\n", res)
	// Output:
	// res 80
}

func ExampleDecorator2() {
	var c Component = WarpAddDecorator(WarpMulDecorator(&ConcreteComponent{}, 8), 10)
	res := c.Calc()

	fmt.Printf("res %d\n", res)
	// Output:
	// res 10
}

package builder

// Builder 是生成器接口
type BuilderI2 interface {
	Part1() BuilderI2
	Part2() BuilderI2
	Part3() BuilderI2
}

type Director2 struct {
	builder2 BuilderI2
}

// NewDirector ...
func NewDirector2(builder2 BuilderI2) *Director2 {
	return &Director2{
		builder2: builder2,
	}
}

// Construct Product
func (d *Director2) Construct2() {
	d.builder2.Part1().Part2().Part3()
}

// Construct Product
func (d *Director2) setPart1() *Director2 {
	d.builder2.Part1()
	return d
}

// Construct Product
func (d *Director2) setPart2() *Director2 {
	d.builder2.Part2()
	return d
}

// Construct Product
func (d *Director2) setPart3() *Director2 {
	d.builder2.Part3()
	return d
}

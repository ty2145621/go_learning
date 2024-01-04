package main

type allInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type allUint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type ints interface { // 交集
	allInt
	allUint
}

func max[T allInt](x, y T) T {
	if x > y {
		return x
	}
	return y
}

type MyMap[K int | string, V float32 | float64] map[K]V

var aMap MyMap[int, float64] = map[int]float64{1: 1.01}

type MyStruct[T int | string] struct {
	Name string
	Data T
}

var aStruct MyStruct[int] = MyStruct[int]{Name: "ty", Data: 1}

type IPrintData[T int | float32 | float64] interface {
	Print(data T)
}

type MySlice[T int | float64] []T

func (s MySlice[T]) Sum() T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func getSum() int {
	var aSlice MySlice[int] = []int{1, 2, 3, 4}
	return aSlice.Sum()
}

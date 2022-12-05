package tests

type MathInterface interface {
	MockAdder([]int) (int, error)
}

type exampleStruct struct{}

func (*exampleStruct) MockAdder(values []int) (int, error) {
	adder := 0

	for i := range values {
		adder = adder + values[i]
	}

	return adder, nil
}

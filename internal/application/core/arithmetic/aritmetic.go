package arithmetic

type Arith struct{}

func NewAdapter() *Arith {
	return &Arith{}
}

func (arith Arith) Addition(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arith Arith) Substraction(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arith Arith) Multiplication(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arith Arith) Devision(a int32, b int32) (int32, error) {
	return a / b, nil
}

package base

type State int

const (
	Success State = iota
	FailureReadBody
	FailureParseJSON
	FailureCreateJSON
)

type Response struct {
	Status State
	Result int
}

type Expression struct {
	First    int
	Second   int
	Operator rune
}

func (e *Expression) Solve() int {
	switch e.Operator {
	case '*':
		return e.First * e.Second
	case '+':
		return e.First + e.Second
	case '/':
		return e.First / e.Second
	case '-':
		return e.First - e.Second
	default:
		return -1
	}
}

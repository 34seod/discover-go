package func_practice

type MultiSet map[string]int
type SetOp func(m MultiSet, val string)

func Insert(m MultiSet, val string)

func InsertFunc(m MultiSet) func(val string) {
	return func(val string) {
		Insert(m, val)
	}
}

func BindMap(f SetOp, m MultiSet) func(val string) {
	return func(val string) {
		f(m, val)
	}
}

func NewMultiSet() MultiSet {
	return make(MultiSet)
}

// func main() {
// 	m := NewMultiSet()
// 	ReadFrom(r, BindMap(Insert, m))
// }

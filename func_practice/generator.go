package func_practice

func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func NewVertexIDGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func NewEdgeIDGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

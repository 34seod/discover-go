package practice3

type Queue struct {
	q []int
}

func (q *Queue) Add(n int) {
	q.q = append(q.q, n)
}

func (q *Queue) Pop() (result int) {
	if len(q.q) > 0 {
		result = q.q[0]
		q.q = q.q[1:]
	}
	return
}

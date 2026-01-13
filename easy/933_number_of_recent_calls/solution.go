package p933_number_of_recent_calls

type RecentCounter struct {
	q []int
}

func Constructor() RecentCounter {
	return RecentCounter{q: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.q = append(this.q, t)
	winStart := t - 3000
	i := 0
	for i < len(this.q) && this.q[i] < winStart {
		i++
	}
	this.q = this.q[i:]
	return len(this.q)
}

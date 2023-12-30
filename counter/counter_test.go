package counter

import "testing"

func TestCounter(t *testing.T) {
	counter := New[int]()
	counter.Inc(3, 2)
	t.Log(counter.Data())

	var list []int = []int{1, 2, 3}
	t.Log(list)
	for index, value := range list {
		t.Log(index, value)
	}
}

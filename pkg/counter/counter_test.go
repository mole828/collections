package counter

import "testing"

func TestCounter(t *testing.T) {
	counter := NewCounter[int]()
	counter.Inc(3, 2)
	t.Log(counter.Data())
}

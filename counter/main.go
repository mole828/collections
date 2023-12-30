package counter

import (
	"github.com/mole828/collections/defaultdict"
)

type Counter[K comparable] struct {
	defaultdict.DefaultDict[K, int]
}

func (c Counter[K]) Inc(key K, value int) {
	c.Set(key, c.Get(key)+value)
}

func New[K comparable]() Counter[K] {
	dict := defaultdict.New[K, int](func() int {
		return 0
	})
	return Counter[K]{
		dict,
	}
}

func CountAll[K comparable](iter []K) Counter[K] {
	counter := New[K]()
	for _, value := range iter {
		counter.Inc(value, 1)
	}
	return counter
}

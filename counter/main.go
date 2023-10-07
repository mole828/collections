package counter

import "github.com/mole828/collections/defaultdict"

type Counter[K comparable] struct {
	defaultdict.DefaultDict[K, int]
}

func (c Counter[K]) Inc(key K, value int) {
	c.Set(key, c.Get(key)+value)
}

func NewCounter[K comparable]() Counter[K] {
	dict := defaultdict.NewDefaultDict[K, int](func() int {
		return 0
	})
	return Counter[K]{
		dict,
	}
}

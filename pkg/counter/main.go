package counter

import "github.com/mole828/collections/pkg/defaultdict"

type Counter[K comparable] struct {
	defaultdict.DefaultDict[K, int]
}

func NewCounter[K comparable]() Counter[K] {
	dict := defaultdict.NewDefaultDict[K, int](func() int {
		return 0
	})
	return Counter[K]{
		dict,
	}
}

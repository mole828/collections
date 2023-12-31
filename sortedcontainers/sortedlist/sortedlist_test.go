package sortedlist_test

import (
	"testing"

	"github.com/mole828/collections/sortedcontainers/sortedlist"
)

func TestCompare(t *testing.T) {
	var i int
	t.Logf("i: %d", i)
}

func printSortedList[T any](t *testing.T, list *sortedlist.SortedList[T]) {
	p := list.Head()
	arr := []T{}
	for p != nil {
		arr = append(arr, p.Value())
		p = p.Next()
	}
	t.Log(arr)
}

func TestSortedList(t *testing.T) {
	t.Log("begin")
	list := sortedlist.New[int](func(i1, i2 int) bool {
		t.Logf("i1: %d, i2: %d", i1, i2)
		return i1 <= i2
	})
	list.Add(3)
	printSortedList(t, list)
	list.Add(1)
	printSortedList(t, list)
	list.Add(5)
	printSortedList(t, list)
	list.Add(-3)
	list.Add(3)
	list.Add(11)
	printSortedList(t, list)

	list.Remove(func(t int) bool {
		return t == 3
	})
	printSortedList(t, list)
}

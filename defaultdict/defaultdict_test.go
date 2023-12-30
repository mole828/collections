package defaultdict

import "testing"

func TestA(t *testing.T) {
	dict := New[int, int](func() int { return 0 })
	t.Log(dict.Get(1))
	t.Log(dict.Data())
}

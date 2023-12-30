package defaultdict

type DefaultDict[K comparable, V any] struct {
	data         map[K]V
	defaultValue func() V
}

func (dict DefaultDict[K, V]) Get(key K) V {
	value, has := dict.data[key]
	if !has {
		value = dict.defaultValue()
		dict.data[key] = value
	}
	return value
}

func (dict DefaultDict[K, V]) Set(key K, value V) {
	dict.data[key] = value
}

func (dict DefaultDict[K, V]) Data() map[K]V {
	return dict.data
}

func New[K comparable, V any](f func() V) DefaultDict[K, V] {
	return DefaultDict[K, V]{
		data:         map[K]V{},
		defaultValue: f,
	}
}

package syncmap

import "sync"

type Map[K comparable, V any] struct {
	mp *sync.Map
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{mp: &sync.Map{}}
}

func (m *Map[K, V]) CompareAndDelete(key K, value V) bool {
	return m.mp.CompareAndDelete(key, value)
}

func (m *Map[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.mp.CompareAndSwap(key, old, new)
}

func (m *Map[K, V]) Delete(key K) {
	m.mp.Delete(key)
}

func (m *Map[K, V]) Load(key K) (res V, ok bool) {
	value, ok := m.mp.Load(key)
	if !ok {
		return res, false
	}
	return value.(V), true
}

func (m *Map[K, V]) LoadAndDelete(key K) (res V, ok bool) {
	value, ok := m.mp.LoadAndDelete(key)
	if !ok {
		return res, false
	}
	return value.(V), true
}

func (m *Map[K, V]) LoadOrStore(key K, new V) (V, bool) {
	value, ok := m.mp.LoadOrStore(key, new)
	if !ok {
		return value.(V), false
	}
	return value.(V), true
}

func (m *Map[K, V]) Range(run func(key K, value V) bool) {
	m.mp.Range(func(k, v any) bool {
		return run(k.(K), v.(V))
	})
}

func (m *Map[K, V]) Store(key K, value V) {
	m.mp.Store(key, value)
}

func (m *Map[K, V]) Swap(key K, value V) (pre V, ok bool) {
	previous, ok := m.mp.Swap(key, value)
	if !ok {
		return pre, false
	}
	return previous.(V), true
}

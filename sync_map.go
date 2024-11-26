package syncmap

import "sync"

// Map is a generic wrapper for sync.Map that supports type-safe operations.
// Map 是一个面向 sync.Map 的泛型封装，支持精确的类型操作。
type Map[K comparable, V any] struct {
	mp *sync.Map
}

// NewMap creates a new instance of Map.
// NewMap 创建一个新的 Map 实例。
func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{mp: &sync.Map{}}
}

// CompareAndDelete compares the existing value with the provided value and deletes the entry if they match.
// CompareAndDelete 比较存在的值与提供的值，如果匹配，则删除条目。
func (m *Map[K, V]) CompareAndDelete(key K, value V) bool {
	return m.mp.CompareAndDelete(key, value)
}

// CompareAndSwap compares the existing value with the old value and swaps it with the new value if they match.
// CompareAndSwap 将存在的值与旧值比较，如果匹配，则替换为新值。
func (m *Map[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.mp.CompareAndSwap(key, old, new)
}

// Delete removes the value associated with the specified key.
// Delete 删除指定键相关的值。
func (m *Map[K, V]) Delete(key K) {
	m.mp.Delete(key)
}

// Load retrieves the value associated with the specified key.
// Load 获取与指定键相关的值。
func (m *Map[K, V]) Load(key K) (res V, ok bool) {
	value, ok := m.mp.Load(key)
	if !ok {
		return res, false
	}
	return value.(V), true
}

// LoadAndDelete retrieves and removes the value associated with the specified key.
// LoadAndDelete 获取并删除与指定键相关的值。
func (m *Map[K, V]) LoadAndDelete(key K) (res V, ok bool) {
	value, ok := m.mp.LoadAndDelete(key)
	if !ok {
		return res, false
	}
	return value.(V), true
}

// LoadOrStore retrieves the value associated with the key or stores and returns the new value if the key does not exist.
// LoadOrStore 获取指定键相关的值，如果键不存在，则存储并返回新值。
func (m *Map[K, V]) LoadOrStore(key K, new V) (V, bool) {
	value, ok := m.mp.LoadOrStore(key, new)
	if !ok {
		return value.(V), false
	}
	return value.(V), true
}

// Range iterates over all key-value pairs in the map.
// Range 遍历地图中的所有键值对。
func (m *Map[K, V]) Range(run func(key K, value V) bool) {
	m.mp.Range(func(k, v any) bool {
		return run(k.(K), v.(V))
	})
}

// Store sets the value for the specified key.
// Store 为指定的键设置值。
func (m *Map[K, V]) Store(key K, value V) {
	m.mp.Store(key, value)
}

// Swap replaces the value for the specified key and returns the previous value if it existed.
// Swap 替换指定键相关的值，并返回之前的值（如果存在）。
func (m *Map[K, V]) Swap(key K, value V) (pre V, ok bool) {
	previous, ok := m.mp.Swap(key, value)
	if !ok {
		return pre, false
	}
	return previous.(V), true
}

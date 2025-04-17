package syncmap

import (
	"fmt"
)

/*
Offers some utils functions to make usage simple. Support tasks like getting keys, values, and pairs, and copying data between maps.
提供辅助函数，简化使用。这些函数支持获取键、值、键值对或在 map 间复制数据等操作。
*/

// GetMap returns all key-value pairs as a standard Go map.
// GetMap 返回所有键值对，格式为标准的 Go map。
func (m *Map[K, V]) GetMap() map[K]V {
	res := make(map[K]V)
	m.Range(func(k K, v V) bool {
		res[k] = v
		return true
	})
	return res
}

// SetMap adds or updates key-value pairs from a standard Go map.
// SetMap 将标准 Go map 中的键值对添加或更新到 Map。
func (m *Map[K, V]) SetMap(mp map[K]V) {
	for k, v := range mp {
		m.Store(k, v)
	}
}

// SetSyncMap copies all key-value pairs from a Map.
// SetSyncMap 复制一个 Map 中的所有键值对到当前 Map。
func (m *Map[K, V]) SetSyncMap(mp *Map[K, V]) {
	mp.Range(func(k K, v V) bool {
		m.Store(k, v)
		return true
	})
}

// SetSyncMaps copies all key-value pairs from multiple Maps.
// SetSyncMaps 将多个 Map 中的键值对复制到当前 Map。
func (m *Map[K, V]) SetSyncMaps(mps ...*Map[K, V]) {
	for _, mp := range mps {
		m.SetSyncMap(mp)
	}
}

// Debug prints all key-value pairs for debugging purposes.
// Debug 打印所有键值对，主要用于调试逻辑。
func (m *Map[K, V]) Debug() {
	fmt.Println("-----------")
	m.Range(func(k K, v V) bool {
		fmt.Println(k, v)
		return true
	})
	fmt.Println("-----------")
}

// Count returns the count num of key-value pairs in the Map.
// Count 返回 Map 中的键值对数量。
func (m *Map[K, V]) Count() (size int) {
	m.Range(func(k K, v V) bool {
		size++
		return true
	})
	return size
}

// Keys retrieves all keys from the Map.
// Keys 返回 Map 中的所有键。
func (m *Map[K, V]) Keys() (keys []K) {
	m.Range(func(k K, v V) bool {
		keys = append(keys, k)
		return true
	})
	return keys
}

// Values retrieves all values from the Map.
// Values 返回 Map 中的所有值。
func (m *Map[K, V]) Values() (values []V) {
	m.Range(func(k K, v V) bool {
		values = append(values, v)
		return true
	})
	return values
}

package syncmap

import (
	"fmt"
)

/*
以下函数是扩展功能，比如得到 Keys 和 Values 等，基本都是简单的语法糖，没有特别深的逻辑。
*/

func (m *Map[K, V]) GetMap() map[K]V {
	var res = map[K]V{}
	m.Range(func(k K, v V) bool {
		res[k] = v
		return true
	})
	return res
}

func (m *Map[K, V]) SetMap(mp map[K]V) {
	for k, v := range mp {
		m.Store(k, v)
	}
}

func (m *Map[K, V]) SetSyncMap(mp *Map[K, V]) {
	mp.Range(func(k K, v V) bool {
		m.Store(k, v)
		return true
	})
}

func (m *Map[K, V]) SetSyncMaps(mps ...*Map[K, V]) {
	for _, mp := range mps {
		m.SetSyncMap(mp)
	}
}

func (m *Map[K, V]) Debug() {
	fmt.Println("-----------")
	m.Range(func(k K, v V) bool {
		fmt.Println(k, v)
		return true
	})
	fmt.Println("-----------")
}

func (m *Map[K, V]) Count() (size int) {
	m.Range(func(k K, v V) bool {
		size++
		return true
	})
	return size
}

func (m *Map[K, V]) Keys() (keys []K) {
	m.Range(func(k K, v V) bool {
		keys = append(keys, k)
		return true
	})
	return keys
}

func (m *Map[K, V]) Values() (values []V) {
	m.Range(func(k K, v V) bool {
		values = append(values, v)
		return true
	})
	return values
}

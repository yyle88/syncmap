package syncmap

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap_Store(t *testing.T) {
	var mp = NewMap[string, int]()
	mp.Store("a", 1)
	value, ok := mp.Load("a")
	t.Log(ok)
	require.True(t, ok)
	t.Log(value)
	require.Equal(t, 1, value)
	t.Log(value + value) //很明显的这里已经返回 int 结果，就可以按照 int 使用
	t.Log(value * 2)     //也可以使用乘法操作，而无需 interface {} 类型转换
}

/*
以下几个测试分别是使用 sync.Map 和 syncmap.Map 的单元测试，流程相同，结果也相同
*/

func TestSyncMap(t *testing.T) {
	// 创建一个 sync.Map 对象
	var m sync.Map

	// 测试 Store 和 Load 方法
	m.Store("foo", 1)
	if v, ok := m.Load("foo"); !ok || v.(int) != 1 {
		t.Errorf("Load failed: %v, %v", v, ok)
	}

	// 测试 LoadOrStore 方法
	if v, ok := m.LoadOrStore("bar", 2); ok || v.(int) != 2 {
		t.Errorf("LoadOrStore failed: %v, %v", v, ok)
	}
	if v, ok := m.LoadOrStore("foo", 3); !ok || v.(int) != 1 {
		t.Errorf("LoadOrStore failed: %v, %v", v, ok)
	}

	// 测试 Range 方法
	m.Store("baz", 3)
	m.Range(func(k, v interface{}) bool {
		switch k {
		case "foo":
			if v.(int) != 1 {
				t.Errorf("Range failed: %v", v)
			}
		case "bar":
			if v.(int) != 2 {
				t.Errorf("Range failed: %v", v)
			}
		case "baz":
			if v.(int) != 3 {
				t.Errorf("Range failed: %v", v)
			}
		default:
			t.Errorf("Range failed: unexpected key %v", k)
		}
		return true
	})

	// 测试 Delete 方法
	m.Delete("foo")
	if _, ok := m.Load("foo"); ok {
		t.Errorf("Delete failed")
	}
}

func TestSyncMap2(t *testing.T) {
	// 创建一个 sync.Map 对象
	var m sync.Map

	// 测试 Store 和 Load 方法
	m.Store("foo", 1)
	if v, ok := m.Load("foo"); !ok || v.(int) != 1 {
		t.Errorf("Load failed: %v, %v", v, ok)
	}

	// 测试 LoadOrStore 方法
	if v, ok := m.LoadOrStore("bar", 2); ok || v.(int) != 2 {
		t.Errorf("LoadOrStore failed: %v, %v", v, ok)
	}
	if v, ok := m.LoadOrStore("foo", 3); !ok || v.(int) != 1 {
		t.Errorf("LoadOrStore failed: %v, %v", v, ok)
	}

	// 测试 Range 方法
	m.Store("baz", 3)
	var keys []string
	var values []int
	m.Range(func(k, v interface{}) bool {
		keys = append(keys, k.(string))
		values = append(values, v.(int))
		return true
	})

	if len(keys) != 3 || len(values) != 3 {
		t.Errorf("Range failed: %v, %v", keys, values)
	}
	if !containsValue(keys, "foo") || !containsValue(keys, "bar") || !containsValue(keys, "baz") {
		t.Errorf("Range failed: %v", keys)
	}
	if !containsValue(values, 1) || !containsValue(values, 2) || !containsValue(values, 3) {
		t.Errorf("Range failed: %v", values)
	}

	// 测试 Delete 方法
	m.Delete("foo")
	if _, ok := m.Load("foo"); ok {
		t.Errorf("Delete failed")
	}
}

func containsValue(slice interface{}, item interface{}) bool {
	switch s := slice.(type) {
	case []string:
		for _, v := range s {
			if v == item.(string) {
				return true
			}
		}
	case []int:
		for _, v := range s {
			if v == item.(int) {
				return true
			}
		}
	}
	return false
}

/*
使用 syncmap.Map 也能得到同样的效果。
*/

func TestSyncMap3(t *testing.T) {
	// 创建一个 sync.Map 对象
	var m = NewMap[string, int]()

	// 测试 Store 和 Load 方法
	m.Store("foo", 1)
	if v, ok := m.Load("foo"); !ok || v != 1 {
		t.Errorf("Load failed: %v, %v", v, ok)
	}

	// 测试 LoadOrStore 方法
	if v, ok := m.LoadOrStore("bar", 2); ok || v != 2 {
		t.Errorf("LoadOrStore failed: %v, %v", v, ok)
	}
	if v, ok := m.LoadOrStore("foo", 3); !ok || v != 1 {
		t.Errorf("LoadOrStore failed: %v, %v", v, ok)
	}

	// 测试 Range 方法
	m.Store("baz", 3)
	m.Range(func(k string, v int) bool {
		switch k {
		case "foo":
			if v != 1 {
				t.Errorf("Range failed: %v", v)
			}
		case "bar":
			if v != 2 {
				t.Errorf("Range failed: %v", v)
			}
		case "baz":
			if v != 3 {
				t.Errorf("Range failed: %v", v)
			}
		default:
			t.Errorf("Range failed: unexpected key %v", k)
		}
		return true
	})

	// 测试 Delete 方法
	m.Delete("foo")
	if _, ok := m.Load("foo"); ok {
		t.Errorf("Delete failed")
	}
}

func TestSyncMap4(t *testing.T) {
	// 创建一个 sync.Map 对象
	var m = NewMap[string, int64]()

	// 测试 Store 和 Load 方法
	m.Store("foo", 1)
	if v, ok := m.Load("foo"); !ok || v != 1 {
		t.Errorf("Load failed: %v, %v", v, ok)
	}

	// 测试 LoadOrStore 方法
	if v, ok := m.LoadOrStore("bar", 2); ok || v != 2 {
		t.Errorf("LoadOrStore failed: %v, %v", v, ok)
	}
	if v, ok := m.LoadOrStore("foo", 3); !ok || v != 1 {
		t.Errorf("LoadOrStore failed: %v, %v", v, ok)
	}

	// 测试 Range 方法
	m.Store("baz", 3)
	var keys []string
	var values []int64
	m.Range(func(k string, v int64) bool {
		keys = append(keys, k)
		values = append(values, v)
		return true
	})

	if len(keys) != 3 || len(values) != 3 {
		t.Errorf("Range failed: %v, %v", keys, values)
	}
	if !sliceContains(keys, "foo") || !sliceContains(keys, "bar") || !sliceContains(keys, "baz") {
		t.Errorf("Range failed: %v", keys)
	}
	if !sliceContains(values, 1) || !sliceContains(values, 2) || !sliceContains(values, 3) {
		t.Errorf("Range failed: %v", values)
	}
	// 测试 Delete 方法
	m.Delete("foo")
	if _, ok := m.Load("foo"); ok {
		t.Errorf("Delete failed")
	}
}

func sliceContains[V comparable](slice []V, value V) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

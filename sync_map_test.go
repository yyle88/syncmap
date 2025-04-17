package syncmap_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/syncmap"
)

func TestMap_New(t *testing.T) {
	// Test both New and NewMap for consistency
	m1 := syncmap.New[string, int]()
	m2 := syncmap.NewMap[string, int]()
	require.NotNil(t, m1, "New should return a non-nil Map")
	require.NotNil(t, m2, "NewMap should return a non-nil Map")

	// Verify initial state
	_, ok := m1.Load("key")
	require.False(t, ok, "New Map should be empty")
	_, ok = m2.Load("key")
	require.False(t, ok, "NewMap should be empty")
}

func TestMap_StoreAndLoad(t *testing.T) {
	m := syncmap.NewMap[string, float64]()
	m.Store("pi", 3.14)

	value, ok := m.Load("pi")
	require.True(t, ok, "Load should return true for existing key")
	require.Equal(t, 3.14, value, "Load should return stored value")

	_, ok = m.Load("unknown")
	require.False(t, ok, "Load should return false for non-existing key")
}

func TestMap_Store(t *testing.T) {
	var mp = syncmap.NewMap[string, int]()
	mp.Store("a", 1)
	value, ok := mp.Load("a")
	t.Log(ok)
	require.True(t, ok)
	t.Log(value)
	require.Equal(t, 1, value)
	t.Log(value + value) //很明显的这里已经返回 int 结果，就可以按照 int 使用
	t.Log(value * 2)     //也可以使用乘法操作，而无需 interface {} 类型转换
}

func TestMap_Delete(t *testing.T) {
	m := syncmap.NewMap[int, string]()
	m.Store(1, "one")

	m.Delete(1)
	_, ok := m.Load(1)
	require.False(t, ok, "Delete should remove the key")

	// Deleting non-existing key should not panic
	m.Delete(2)
	_, ok = m.Load(2)
	require.False(t, ok, "Deleting non-existing key should have no effect")
}

func TestMap_LoadOrStore(t *testing.T) {
	m := syncmap.NewMap[string, bool]()

	// Store new key
	value, loaded := m.LoadOrStore("active", true)
	require.False(t, loaded, "LoadOrStore should return false for new key")
	require.True(t, value, "LoadOrStore should return stored value")

	// Load existing key
	value, loaded = m.LoadOrStore("active", false)
	require.True(t, loaded, "LoadOrStore should return true for existing key")
	require.True(t, value, "LoadOrStore should return original value")
}

func TestMap_LoadAndDelete(t *testing.T) {
	m := syncmap.NewMap[int, string]()
	m.Store(42, "answer")

	value, ok := m.LoadAndDelete(42)
	require.True(t, ok, "LoadAndDelete should return true for existing key")
	require.Equal(t, "answer", value, "LoadAndDelete should return correct value")
	_, ok = m.Load(42)
	require.False(t, ok, "Key should be deleted")

	value, ok = m.LoadAndDelete(99)
	require.False(t, ok, "LoadAndDelete should return false for non-existing key")
	require.Equal(t, "", value, "LoadAndDelete should return zero value")
}

func TestMap_Swap(t *testing.T) {
	m := syncmap.NewMap[string, int]()
	m.Store("count", 10)

	// Swap existing key
	{
		previous, ok := m.Swap("count", 20)
		require.True(t, ok, "Swap should return true for existing key")
		require.Equal(t, 10, previous, "Swap should return previous value")
		value, _ := m.Load("count")
		require.Equal(t, 20, value, "New value should be set")
	}

	// Swap non-existing key
	{
		previous, ok := m.Swap("new", 30)
		require.False(t, ok, "Swap should return false for non-existing key")
		require.Equal(t, 0, previous)
		value, ok := m.Load("new")
		require.True(t, ok, "Key should now exist")
		require.Equal(t, 30, value, "New value should be set")
	}
}

func TestMap_CompareAndSwap(t *testing.T) {
	m := syncmap.NewMap[string, string]()
	m.Store("color", "blue")

	// Wrong old value
	result := m.CompareAndSwap("color", "red", "green")
	require.False(t, result, "CompareAndSwap should fail with wrong old value")
	value, _ := m.Load("color")
	require.Equal(t, "blue", value, "Value should remain unchanged")

	// Correct old value
	result = m.CompareAndSwap("color", "blue", "green")
	require.True(t, result, "CompareAndSwap should succeed with correct old value")
	value, _ = m.Load("color")
	require.Equal(t, "green", value, "Value should be swapped")
}

func TestMap_CompareAndDelete(t *testing.T) {
	m := syncmap.NewMap[string, int]()
	m.Store("x", 100)

	// Wrong value
	result := m.CompareAndDelete("x", 200)
	require.False(t, result, "CompareAndDelete should fail with wrong value")
	_, ok := m.Load("x")
	require.True(t, ok, "Key should still exist")

	// Correct value
	result = m.CompareAndDelete("x", 100)
	require.True(t, result, "CompareAndDelete should succeed with correct value")
	_, ok = m.Load("x")
	require.False(t, ok, "Key should be deleted")
}

func TestMap_Range(t *testing.T) {
	m := syncmap.NewMap[string, int]()
	m.Store("a", 1)
	m.Store("b", 2)

	count := 0
	m.Range(func(k string, v int) bool {
		count++
		switch k {
		case "a":
			require.Equal(t, 1, v, "Range should yield correct value for key 'a'")
		case "b":
			require.Equal(t, 2, v, "Range should yield correct value for key 'b'")
		default:
			t.Errorf("Unexpected key: %v", k)
		}
		return true
	})
	require.Equal(t, 2, count, "Range should iterate over all entries")
}

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

func TestSyncMap3(t *testing.T) {
	// 创建一个 sync.Map 对象
	var m = syncmap.NewMap[string, int]()

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
	var m = syncmap.NewMap[string, int64]()

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

func sliceContains[V comparable](slice []V, value V) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

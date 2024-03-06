package syncmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap_GetMap(t *testing.T) {
	m := NewMap[int, string]()
	m.Store(1, "a")
	m.Store(2, "b")
	m.Store(3, "c")
	require.Equal(t, map[int]string{
		1: "a", 2: "b", 3: "c",
	}, m.GetMap())
}

func TestMap_SetMap(t *testing.T) {
	m := NewMap[string, int]()
	m.SetMap(map[string]int{"a": 1, "b": 2})
	v, ok := m.Load("a")
	require.True(t, ok)
	require.Equal(t, 1, v)
	v, ok = m.Load("b")
	require.True(t, ok)
	require.Equal(t, 2, v)
	v, ok = m.Load("c")
	require.False(t, ok)
	require.Equal(t, 0, v)
}

func TestMap_SetSyncMap(t *testing.T) {
	m0 := NewMap[string, int]()
	m0.SetMap(map[string]int{"a": 1, "b": 2})
	m1 := NewMap[string, int]()
	m1.SetMap(map[string]int{"c": 3, "d": 4})
	m := NewMap[string, int]()
	m.SetSyncMap(m0)
	m.SetSyncMap(m1)
	require.Equal(t, map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4,
	}, m.GetMap())
}

func TestMap_SetSyncMaps(t *testing.T) {
	m0 := NewMap[string, int]()
	m0.SetMap(map[string]int{"a": 1, "b": 2})
	m1 := NewMap[string, int]()
	m1.SetMap(map[string]int{"c": 3, "d": 4})
	m := NewMap[string, int]()
	m.SetSyncMaps(m0, m1)
	require.Equal(t, map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4,
	}, m.GetMap())
}

func TestMap_Count(t *testing.T) {
	m := NewMap[string, int]()
	m.SetMap(map[string]int{"a": 1, "b": 2})
	require.Equal(t, 2, m.Count())
}

func TestMap_Keys(t *testing.T) {
	m := NewMap[int, string]()
	m.Store(1, "a")
	m.Store(2, "b")
	m.Store(3, "c")
	keys := m.Keys()
	require.Equal(t, 3, len(keys))
	for _, k := range []int{1, 2, 3} {
		require.True(t, sliceContains(keys, k))
	}
}

func TestMap_Values(t *testing.T) {
	m := NewMap[int, string]()
	m.Store(1, "a")
	m.Store(2, "b")
	m.Store(3, "c")
	values := m.Values()
	require.Equal(t, 3, len(values))
	for _, v := range []string{"a", "b", "c"} {
		require.True(t, sliceContains(values, v))
	}
}

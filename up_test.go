package syncmap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/syncmap"
)

func TestMap_GetMap(t *testing.T) {
	{
		m := syncmap.NewMap[int, string]()
		m.Store(1, "a")
		m.Store(2, "b")
		m.Store(3, "c")
		require.Equal(t, map[int]string{
			1: "a", 2: "b", 3: "c",
		}, m.GetMap())
	}
	{
		m := syncmap.NewMap[string, int]()
		result := m.GetMap()
		require.Empty(t, result, "GetMap should return an empty map for an empty Map")
		require.Equal(t, map[string]int{}, result, "GetMap should return an empty map")
	}
}

func TestMap_SetMap(t *testing.T) {
	{
		m := syncmap.NewMap[string, int]()
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
	{
		m := syncmap.NewMap[string, int]()
		m.SetMap(map[string]int{})
		require.Equal(t, 0, m.Count(), "SetMap with empty map should result in empty Map")
		_, ok := m.Load("any")
		require.False(t, ok, "No keys should exist after setting empty map")
	}
}

func TestMap_SetSyncMap(t *testing.T) {
	{
		m0 := syncmap.NewMap[string, int]()
		m0.SetMap(map[string]int{"a": 1, "b": 2})
		m1 := syncmap.NewMap[string, int]()
		m1.SetMap(map[string]int{"c": 3, "d": 4})
		m := syncmap.NewMap[string, int]()
		m.SetSyncMap(m0)
		m.SetSyncMap(m1)
		require.Equal(t, map[string]int{
			"a": 1, "b": 2, "c": 3, "d": 4,
		}, m.GetMap())
	}
	{
		m0 := syncmap.NewMap[string, int]()
		m := syncmap.NewMap[string, int]()
		m.SetSyncMap(m0)
		require.Equal(t, 0, m.Count(), "SetSyncMap with empty Map should result in empty Map")
		_, ok := m.Load("any")
		require.False(t, ok, "No keys should exist after setting empty sync map")
	}
}

func TestMap_SetSyncMaps(t *testing.T) {
	{
		m0 := syncmap.NewMap[string, int]()
		m0.SetMap(map[string]int{"a": 1, "b": 2})
		m1 := syncmap.NewMap[string, int]()
		m1.SetMap(map[string]int{"c": 3, "d": 4})
		m := syncmap.NewMap[string, int]()
		m.SetSyncMaps(m0, m1)
		require.Equal(t, map[string]int{
			"a": 1, "b": 2, "c": 3, "d": 4,
		}, m.GetMap())
	}
	{
		m0 := syncmap.NewMap[string, int]() // Empty map
		m1 := syncmap.NewMap[string, int]()
		m1.SetMap(map[string]int{"x": 10})
		m := syncmap.NewMap[string, int]()
		m.SetSyncMaps(m0, m1)
		require.Equal(t, map[string]int{"x": 10}, m.GetMap(), "SetSyncMaps should copy only non-empty map's contents")
	}
}

func TestMap_Count(t *testing.T) {
	{
		m := syncmap.NewMap[string, int]()
		m.SetMap(map[string]int{"a": 1, "b": 2})
		require.Equal(t, 2, m.Count())
	}
	{
		m := syncmap.NewMap[string, int]()
		require.Equal(t, 0, m.Count(), "Count should return 0 for empty Map")
	}
}

func TestMap_Keys(t *testing.T) {
	{
		m := syncmap.NewMap[int, string]()
		m.Store(1, "a")
		m.Store(2, "b")
		m.Store(3, "c")
		keys := m.Keys()
		require.Equal(t, 3, len(keys))
		for _, k := range []int{1, 2, 3} {
			require.True(t, sliceContains(keys, k))
		}
	}
	{
		m := syncmap.NewMap[string, int]()
		keys := m.Keys()
		require.Empty(t, keys, "Keys should return empty slice for empty Map")
	}
}

func TestMap_Values(t *testing.T) {
	{
		m := syncmap.NewMap[int, string]()
		m.Store(1, "a")
		m.Store(2, "b")
		m.Store(3, "c")
		values := m.Values()
		require.Equal(t, 3, len(values))
		for _, v := range []string{"a", "b", "c"} {
			require.True(t, sliceContains(values, v))
		}
	}
	{
		m := syncmap.NewMap[string, int]()
		values := m.Values()
		require.Empty(t, values, "Values should return empty slice for empty Map")
	}
}

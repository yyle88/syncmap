# syncmap
Generics sync map

泛型的 sync map

Sync map generics

Sync map 泛型版

## 思路
就是把 sync.Map 使用泛型包裹了一层，让你在使用 sync.Map 的 `Load` 和 `Store` 的时候不用使用 `interface {}` 的装箱和解包，而是直接使用确定的类型，这样便于使用。

在绝大多数场景下同一个 sync.Map 存的 "K 类型相同 且 V 类型相同"，很少有在同一个 sync.Map 里 "K 类型不同 或 V 类型不同" 的情况。

当确定 "K 类型相同 且 V 类型相同" 时，就可以用这个工具，能够避免在使用时类型转换/混淆/出错。

当确定 "K 类型不同 或 V 类型不同" 时，就继续用原来的 `sync.Map` 就行。

## 使用
`go get github.com/yyle88/syncmap`

## 用法
[使用样例](/sync_map_test.go)

## 意图
该工具包100%封装sync.Map的方法，而且方法的参数和返回值都保持不变。
This toolkit provides a comprehensive encapsulation of the methods in sync.Map, ensuring that both the parameters and return values remain unchanged.

因此可以直接替换使用（但注意该类需要调用NewMap的初始化函数）。
As a result, these methods can be seamlessly substituted and utilized (note that this class requires the NewMap initialization function to be called).

该工具使用泛型，能够有效解决原来sync.Map 的无类型问题，避免interface{}转换，当然更能让代码上下文更清楚。
Moreover, the presence of generics in this toolkit eliminates the need for any conversions to interface{}, enhancing its efficiency and convenience, and making the code context clean.

## 说明
比如 `var mp = NewMap[string, int]()` 这样就很清楚的知道 k-v 的类型。

而且在使用 `Store` / `Load` 的时候参数就是确定的，比如 k 被限制为 int， 而返回值 v 也直接返回 string。

省去 res.(string) 的转换，因为我已经通过 `value.(V)` 替你转啦。

## 思路
因为通常，我们定义一个 `sync.Map` 的目的都是存储一类 k-v 数据，很少有在同一个 `sync.Map` 里 k-v 的类型还总是变的情况，因此认为这样封装下更有用

## 扩展
由于已经自定义 `Map` 的 `struct` 因此还可以增加些自定义的操作，比如和普通map相互转换，比如提供`Keys` `Values`这些常用的语法糖等，以便于开发者使用。

## 样例
```
go get github.com/yyle88/syncmap
```

demo1:
```
package main

import (
	"fmt"

	"github.com/yyle88/syncmap"
)

func main() {
	mp := syncmap.NewMap[int, string]()

	mp.Store(1, "a")
	mp.Store(2, "b")
	mp.Store(3, "c")

	mp.Range(func(key int, value string) bool {
		fmt.Println(key, value)
		return true
	})
}
```

demo2:
```
package main

import (
	"fmt"

	"github.com/yyle88/syncmap"
)

type Person struct {
	Name     string
	Age      int
	HomePage string
}

func main() {
	mp := syncmap.NewMap[int, *Person]()

	mp.Store(1, &Person{
		Name:     "Kratos",
		HomePage: "https://go-kratos.dev/",
	})
	mp.Store(2, &Person{
		Name: "YangYiLe",
		Age:  18,
	})
	mp.Store(3, &Person{
		Name: "DiLiReBa",
		Age:  18,
	})

	mp.Delete(3)

	mp.Range(func(key int, value *Person) bool {
		fmt.Println(key, value.Name, value.Age, value.HomePage)
		return true
	})
}
```

Give me stars. Thank you!!!

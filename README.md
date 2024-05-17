# syncmap

Generics sync map

泛型的 sync map

Sync map generics

Sync map 范型版

This toolkit provides a comprehensive encapsulation of the methods in sync.Map, ensuring that both the parameters and return values remain unchanged. As a result, these methods can be seamlessly substituted and utilized. Moreover, the presence of generics in this toolkit eliminates the need for any conversions to interface{}, enhancing its efficiency and convenience.

该工具包100%封装sync.Map的方法，而且方法的参数和返回值都保持不变，因此可以直接替换使用，该工具使用泛型，能够有效解决原来sync.Map 的无类型问题，避免interface{}转换，当然更能让代码上下文更清楚。

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

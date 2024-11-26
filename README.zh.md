# SyncMap - `sync.Map` 的类型安全封装

`SyncMap` 是一个 **类型安全** 和 **泛型** 的包装器，基于 Go 的 `sync.Map`。它通过允许你在创建 `sync.Map` 时定义键和值的类型，使得 `sync.Map` 的使用更加简便和安全。

通过在创建时指定键和值的类型，使得使用 `sync.Map` 更加容易和安全。

## 说明文档

[ENGLISH README](README.md)

## 推荐使用 SyncMap

使用 `SyncMap` 有以下几个优点：

- **泛型类型**：通过定义键和值的类型，避免了因类型不匹配导致的运行时错误。
- **更简洁的代码**：不需要类型断言，避免使用 `interface{}` 的复杂操作。
- **简单易用**：与 `sync.Map` 的使用方式相同，无需学习复杂的新概念。

## 安装

```bash
go get github.com/yyle88/syncmap
```

## 如何使用 SyncMap

以下是一个简单的示例，展示了如何使用 `SyncMap` 安全地存储和检索结构化数据。

### 示例 1：存储和检索数据

```go
package main

import (
	"fmt"
	"github.com/yyle88/syncmap"
)

type User struct {
	Name string
	Age  int
}

func main() {
	// 创建一个以 string 为键、User 结构体指针为值的 SyncMap
	users := syncmap.NewMap[string, *User]()

	// 向 Map 中添加一个用户
	users.Store("u1", &User{Name: "Alice", Age: 30})

	// 检索用户信息
	if user, ok := users.Load("u1"); ok {
		fmt.Printf("用户: 姓名: %s, 年龄: %d\n", user.Name, user.Age) // 输出: 用户: 姓名: Alice, 年龄: 30
	}
}
```

#### 说明

1. **明确的类型定义**：键是 `string`，值是 `User` 结构体的指针，避免了使用 `interface{}` 这个模糊的类型来存储值。
2. **简化的数据访问**：不需要手动进行类型断言，直接获取数据时无需 `user = v.(*User)` 的操作。
3. **与 `sync.Map` 使用方式相同**：与传统的 `sync.Map` 使用方式一致，`Store` 和 `Load` 等方法不需要额外学习。

### 示例 2：使用 SyncMap 存储、删除和遍历数据

```go
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
	// 创建一个以 int 为键、*Person 为值的 SyncMap
	mp := syncmap.NewMap[int, *Person]()

	// 向 Map 中添加几个 Person
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

	// 删除键为 3 的元素
	mp.Delete(3)

	// 遍历 Map 中的所有元素
	mp.Range(func(key int, value *Person) bool {
		fmt.Println(key, value.Name, value.Age, value.HomePage)
		return true
	})
}
```

#### 说明

1. **存储和删除**：这个示例展示了如何存储多个对象，并通过 `Delete` 方法删除指定的元素。
2. **遍历操作**：通过 `Range` 方法遍历 `SyncMap` 中所有的键值对，无需手动操作 `sync.Map` 的底层实现。
3. **简化操作**：相比于传统的 `sync.Map`，这段代码中不需要任何类型断言，操作直观且清晰。

---

## SyncMap API

`SyncMap` 提供以下函数：

| 函数                        | 描述                              |
|---------------------------|---------------------------------|
| `Store(key, value)`       | 添加或更新一个键值对。                     |
| `Load(key)`               | 获取指定键对应的值。                      |
| `LoadOrStore(key, value)` | 如果值已存在则返回该值，否则将指定的键值对添加到 Map 中。 |
| `Delete(key)`             | 删除指定的键值对。                       |
| `Range(func)`             | 遍历 Map 中的所有键值对。                 |

完全与 `sync.Map` 相同。

---

## 许可证

项目采用 MIT 许可证开源。详细的许可证内容请见 [LICENSE](LICENSE) 文件。

---

## 贡献

我们欢迎各种形式的贡献！无论是报告问题、建议新功能，还是提交代码改进，都欢迎提交问题或拉取请求。

---

## 感谢

如果你觉得这个包对你有帮助，请在 GitHub 上给我们一个星标！感谢支持！！！

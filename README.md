[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/syncmap/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/syncmap/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/syncmap)](https://pkg.go.dev/github.com/yyle88/syncmap)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/syncmap/master.svg)](https://coveralls.io/github/yyle88/syncmap?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/syncmap.svg)](https://github.com/yyle88/syncmap/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/syncmap)](https://goreportcard.com/report/github.com/yyle88/syncmap)

# SyncMap - A Type-Safe Wrapper for `sync.Map`

`SyncMap` is a **type-safe** and **generic** wrapper around Go's `sync.Map`. It simplifies the use of `sync.Map` by allowing you to define the types for both keys and values.

It makes using `sync.Map` easier and safer by letting you define the types for keys and values when you create the `sync.Map`.

## Why Choose SyncMap?

Using `SyncMap` has several advantages:

- **Generic Type**: When you define the key and value types, you can avoid type mismatches.
- **Clearer Code**: No need for type assertions, making your code easier to write.
- **Simple Usage**: Works just same as `sync.Map`, so there’s no steep learning curve.

## Installation

```bash
go get github.com/yyle88/syncmap
```

## How to Use SyncMap

Here’s a simple example showing how you can use `SyncMap` to safely store and retrieve structured data.

### Example: Storing and Retrieving Data

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
	// Create a SyncMap with string keys and User values
	users := syncmap.NewMap[string, *User]()

	// Add a user to the map
	users.Store("u1", &User{Name: "Alice", Age: 30})

	// Retrieve the user
	if user, ok := users.Load("u1"); ok {
		fmt.Printf("User: Name: %s, Age: %d\n", user.Name, user.Age) // Output: User: Name: Alice, Age: 30
	}
}
```

### How This Example Helps

1. **Defines Clear Types**: The key is a `string`, and the value is a `User`. Without using `interface{}` to store value.
2. **Simplifies Data Access**: No need for manual type assertions when loading data. No need `user = v.(*User)` logic.
3. **Works Same with `sync.Map`**: If you have used `sync.Map` before, the functions like `Store` and `Load` are same.

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

---

## SyncMap API

`SyncMap` provides the following functions:

| Function                  | Description                                                   |
|---------------------------|---------------------------------------------------------------|
| `Store(key, value)`       | Adds or updates a key-value.                             |
| `Load(key)`               | Retrieves the value.                                |
| `LoadOrStore(key, value)` | Returns the value if it exists; otherwise, adds the new key-value. |
| `Delete(key)`             | Removes a key-value pair from the map.                        |
| `Range(func)`             | Iterates over all key-value pairs in the map.                 |

Same with `sync.Map`

---

## License

This project is open-source under the MIT License. You can find the full text of the license in the [LICENSE](LICENSE) file.

---

## Contributing

We welcome contributions of all kinds! Whether it’s reporting a bug, suggesting a feature, or submitting code improvements.

---

## Thank You

If you find this package valuable, give it a star on GitHub! Thank you!!!

# SyncMap - A Generic Wrapper for Go's `sync.Map`

A type-safe wrapper around Go’s `sync.Map` that uses generics to eliminate the need for `interface{}` type assertions. It provides clear, predictable types for key-value pairs, improving code clarity and reducing errors in concurrent programming.

## README
[中文说明](README.zh.md)

## Overview

Go’s `sync.Map` is a powerful tool for concurrent maps, but it lacks type safety, requiring you to use `interface{}` and type assertions when storing and retrieving values. This can lead to bugs and confusion, especially when the map is used with different types of keys or values.

This package solves that problem by wrapping `sync.Map` with generics. You can now define your map with specific key and value types, avoiding type assertions and making the code clearer and more maintainable.

If you know that your `sync.Map` will only store a consistent type of keys and values, this package makes interacting with it simpler and safer.

## Installation

```bash
go get github.com/yyle88/syncmap
```

## Usage

### Example 1: Using with Basic Types

```go
package main

import (
	"fmt"
	"github.com/yyle88/syncmap"
)

func main() {
	// Create a new syncmap with int keys and string values
	mp := syncmap.NewMap[int, string]()

	// Store some key-value pairs
	mp.Store(1, "a")
	mp.Store(2, "b")
	mp.Store(3, "c")

	// Iterate over the map and print each entry
	mp.Range(func(key int, value string) bool {
		fmt.Println(key, value)
		return true
	})
}
```

### Example 2: Using with Structs

```go
package main

import (
	"fmt"
	"github.com/yyle88/syncmap"
)

// Person struct
type Person struct {
	Name     string
	Age      int
	HomePage string
}

func main() {
	// Create a new syncmap with int keys and Person pointers as values
	mp := syncmap.NewMap[int, *Person]()

	// Store some Person objects in the map
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

	// Delete an entry
	mp.Delete(3)

	// Iterate over the map and print details of each Person
	mp.Range(func(key int, value *Person) bool {
		fmt.Println(key, value.Name, value.Age, value.HomePage)
		return true
	})
}
```

## Why Use This Package?

This package wraps `sync.Map` with Go’s generics, enabling type-safe access to the map’s values. Instead of working with `interface{}`, you define the types of the keys and values at compile time, which eliminates type assertions and makes the code easier to understand and safer to use.

By using generics, you can clearly define the types of data the map will hold. This improves code readability and ensures type safety, especially when dealing with concurrent code.

For example, with this package, you can create a map that holds `int` keys and `string` values like this:

```go
var mp = syncmap.NewMap[int, string]()
```

This makes it clear that the map stores `int` keys and `string` values, and you don’t need to deal with type assertions when calling `Store` or `Load`. The correct type is automatically inferred.

## When Should You Use It?

If you have a `sync.Map` that always stores the same types of keys and values, this package is a great choice. It eliminates the need for type assertions, which simplifies your code and prevents common errors.

If your `sync.Map` needs to store different types of keys or values, you can still use the original `sync.Map`. This package is designed for cases where the key and value types are fixed and known.

## Extending the Map

This package provides a simple and efficient wrapper around `sync.Map`, but you can also extend it with additional methods or features. For example, you could add utility functions like `Keys`, `Values`, or a method to convert the map to a regular Go map.

Since we define the `Map` as a custom struct, you have full control to add any extra functionality that fits your use case.

## Key Features

- **Type Safety**: No more `interface{}` type assertions. Use specific types for keys and values.
- **Clear Code**: By defining the types upfront, the map’s contents are clearer, improving the readability of your code.
- **Direct Compatibility with `sync.Map`**: The API is designed to be a drop-in replacement for `sync.Map`, with the same methods and functionality, but with the added benefit of type safety.
- **Simplicity**: With generics, interacting with a concurrent map becomes easier and safer, without the usual boilerplate for type conversions.

## Example Test

Check out the test file for a full working example: [sync_map_test.go](sync_map_test.go).

## Conclusion

This package simplifies working with `sync.Map` by providing a type-safe wrapper around it. It uses Go’s generics to eliminate type assertions, making the code more maintainable and less error-prone, especially in concurrent applications.

If you find this package helpful, please consider giving it a star! Thank you for using it!

## Thank You

If you find this package valuable, give it a star on GitHub! Thank you!!!

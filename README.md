# Lithium 🔋

Lithium is a simple and efficient caching library for Go, supporting multiple caching strategies.

## Installation

To install Lithium, use `go get`:

```sh
go get github.com/notiku/lithium
```

## Usage

To create a new cache, use the `New` function and specify the caching strategy and any required parameters. The following example creates a new LRU cache with a maximum capacity of 100 items:

```go
package main

import (
    "fmt"
    "github.com/notiku/lithium"
    "github.com/notiku/lithium/rules"
)

func main() {
    c := lithium.New(rules.LRU, 100)
    c.Set("key1", "value1")
    value, found := c.Get("key1")
    if found {
        fmt.Println("Found:", value)
    } else {
        fmt.Println("Not found")
    }
}
```

## Available Caching Strategies

- Least Recently Used (LRU)
- More soon...

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you encounter any problems.

## License

Lithium is licensed under the MIT license. See the [LICENSE](https://github.com/notiku/lithium/tree/master/LICENSE) file for more information.
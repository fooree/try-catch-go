# Usage

```go
package main

import try "github.com/fooree/try-catch-go"

func main() {
	try.Try(func() {
		panic(1)
	}).Catch(func(v interface{}) {
		fmt.Println("catch:", v)
	})
}
```
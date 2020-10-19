# Golang Notes

- Outside a function, every statement begins with a keyword (var, func, and so on) and so the `:=` construct is not available.
- The zero value is:
	```go
	0 for numeric types,
	false for the boolean type, and
	"" (the empty string) for strings.
	```
- Constants cannot be declared with `:=` either.
- Unlike C, Go has no pointer arithmetic.
-  The range form of the for loop iterates over a slice or map. When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
	```go
	package main
	import "fmt"
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	func main() {
		for i, v := range pow {
			fmt.Printf("2**%d = %d\n", i, v)
		}
	}
	```




package main
import "fmt"
func main() {
 foo := func(a string) string { return a }
 fmt.Printf("Type %T\nResult %s\n", foo, foo("foo"))
 passit(foo)
}

func passit(bar func(string) string) {
 fmt.Printf("Inside bar: %s", bar("bar"))
}

```go
package main

type Gopher struct {
	Name string
}

func (g Gopher) TryNotToPanic() {
	if me, ok := recover().(Gopher); ok && me.Name == g.Name {
		print("this is fine")
	}
}

func main() {
	me := Gopher{"Marcin Janas"}

	defer me.TryNotToPanic()

	panic(me)
}
```

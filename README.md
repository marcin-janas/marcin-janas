```go
package main

type gopher struct{}

func (g *gopher) selfRecover() (r bool) {
	if b, ok := recover().(*gopher); ok {
		r = g == b
	}

	return
}

func main() {
	var g gopher
	defer g.selfRecover()

	panic(&g)
}
```

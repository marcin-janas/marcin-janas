```go
package main

type Gopher string

func (g Gopher) Save() {
	for _, s := range []string{
		"https://github.com/",
		"https://www.linkedin.com/in/",
	} {
		println(s + string(g))
	}
}

func main() {
	me := Gopher("marcin-janas")
	defer func() {
		recover().(Gopher).Save()
	}()

	panic(me)
}
```

```sh
$ go run main.go
```
[https://github.com/marcin-janas](https://github.com/marcin-janas)
[https://www.linkedin.com/in/marcin-janas](https://www.linkedin.com/in/marcin-janas)

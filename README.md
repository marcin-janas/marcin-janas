```go
package main

type Gopher struct {
	Name     string
	GitHub   string
	LinkedIn string
	Skills   []string
}

var SaveGopher = func() {
	recover()
}

func main() {
	defer SaveGopher()

	me := Gopher{
		Name:     "Marcin Janas",
		GitHub:   "https://github.com/marcin-janas",
		LinkedIn: "https://www.linkedin.com/in/marcin-janas",
		Skills: []string{
			"Go", "Neovim", "Terraform", "AWS", "Kubernetes",
		},
	}

	panic(me)
}
```

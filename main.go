package main

type Gopher struct {
	Name     string
	Editor   string
	Language string
}

var SaveGopher = func() {
	recover()
}

func main() {
	defer SaveGopher()

	me := Gopher{
		Name:     "Marcin Janas",
		Editor:   "Neovim",
		Language: "Go",
	}

	panic(me)
}

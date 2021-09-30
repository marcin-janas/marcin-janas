package main

import "fmt"

type Gopher struct {
	Name     string
	Editor   string
	Language string
}

func main() {
	me := Gopher{
		Name:     "Marcin Janas",
		Editor:   "Neovim",
		Language: "Go",
	}

	fmt.Println(me)
}

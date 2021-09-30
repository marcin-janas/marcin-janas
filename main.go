package main

import "log"

type Gopher struct {
	Name     string
	Editor   string
	Language string
}

var RecoverGopher = func() {
	recover()
}

func main() {
	defer RecoverGopher()

	me := Gopher{
		Name:     "Marcin Janas",
		Editor:   "Neovim",
		Language: "Go",
	}

	log.Panic(me)
}

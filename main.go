package main

type Gopher struct {
	Name string
}

func (g Gopher) TryNotToPanic() string {
	return g.Name
}

func main() {
	me := Gopher{"Marcin Janas"}

	defer func() {
		if g, ok := recover().(Gopher); ok {
			print(g.TryNotToPanic())
		}
	}()

	panic(me)
}

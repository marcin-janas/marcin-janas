package main

type gopher struct{}

func (g *gopher) selfRecover() {
	var self bool

	if me, ok := recover().(*gopher); ok {
		self = me == g
	}

	print(self)
}

func main() {
	me := &gopher{}
	defer me.selfRecover()

	panic(me)
}

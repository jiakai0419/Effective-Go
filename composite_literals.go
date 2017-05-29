package main

import "fmt"

func display(o interface{}) {
	fmt.Printf("%T %s\n", o, o)
}

func main() {
	const Enone = 1
	const Eio = 2
	const Einval = 3
	a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	display(a)
	s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	display(s)
	m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	display(m)
}

package main

import (
	"fmt"
)

type T struct {
	d int
	r int
}

type S struct {
	i string
	t T
}

func main() {
	m := make(map[string]S)

	m["1"] = S{}

	fmt.Println("M", m)
	fmt.Println("Len: ", len(m))
	e := m["1"]
	e.t = T{d: 12, r: 54}
	m["1"] = e
	fmt.Println("M", m)
	fmt.Println("Len: ", len(m))
}

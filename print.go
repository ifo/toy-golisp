package main

import (
	"fmt"
)

func (n *Number) Print() {
	fmt.Print(*n)
}

func (s *String) Print() {
	fmt.Printf("%q", *s)
}

func (p *Pair) Print() {
	fmt.Print("(")
	p.Head.Print()
	if p.Tail != nil {
		fmt.Print(" . ")
		p.Tail.Print()
	}
	fmt.Print(")")
}

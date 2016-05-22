package main

type LispVal interface {
	Print()
	Eval(LispVal) LispVal
}

type Number int

type Pair struct {
	Head LispVal
	Tail LispVal
}

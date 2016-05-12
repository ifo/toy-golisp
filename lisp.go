package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	bts, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}

	l, rest, err := Read(string(bts))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(l)
	fmt.Println(rest)
}

type LispVal interface {
	Print()
	Eval(LispVal) LispVal
}

type Number int

func (n Number) Print() {
	fmt.Print(n)
}

func (n Number) Eval(env LispVal) LispVal {
	return n
}

type Pair struct {
	Head LispVal
	Tail LispVal
}

func (p Pair) Print() {
	fmt.Print("(")
	p.Head.Print()
	if p.Tail != nil {
		fmt.Print(" . ")
		p.Tail.Print()
	}
	fmt.Print(")")
}

func (p Pair) Eval(env LispVal) LispVal {
	return p.Head
}

func Read(s string) (LispVal, string, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, "", fmt.Errorf("No input source")
	}

	if s[0] == '(' {
		head, rest, err := Read(s[1:])
		if err != nil {
			return nil, s, err
		}

		rest = strings.TrimSpace(rest)
		if !strings.HasPrefix(rest, ".") {
			return nil, s, fmt.Errorf("Invalid syntax")
		}

		tail, rem, err := Read(rest[1:])

		rem = strings.TrimSpace(rem)
		if !strings.HasPrefix(rem, ")") {
			return nil, rem, fmt.Errorf("Invalid syntax")
		}

		lVal := Pair{
			Head: head,
			Tail: tail,
		}

		return lVal, rem[1:], nil
	}

	i := len(s)
	for j, c := range s {
		if !unicode.IsDigit(c) {
			i = j
			break
		}
	}

	num, err := strconv.Atoi(s[:i])
	if err != nil {
		return nil, "", err
	}
	var lVal Number = Number(num)

	return lVal, s[i:], nil
}

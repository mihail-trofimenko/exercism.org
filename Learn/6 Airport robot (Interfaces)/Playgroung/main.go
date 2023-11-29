package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type martian struct{}

type laser int

func (m martian) talk() string {
	return "Цви мигель ууов"
}

func (l laser) talk() string {
	return strings.Repeat("пиу ", int(l))
}

func main() {
	var t talker
	t = martian{}
	fmt.Println(t.talk())
	t = laser(6)
	fmt.Println(t.talk())
}

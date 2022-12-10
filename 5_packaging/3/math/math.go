package math

type Math struct {
	A int
	B int
}

func NewMath(a, b int) *Math {
	return &Math{A: a, B: b}
}

func (m Math) Add() int {
	return m.A + m.B
}

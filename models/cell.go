package models

import (
	"fmt"
)

type cell struct {
	posX    int
	poxY    int
	isAlive bool
}

func NewCell(posX int, poxY int, isAlive bool) cell {
	c := cell{posX, poxY, isAlive}
	return c
}

func (c cell) IsAlive() bool {
	fmt.Printf("isAlive=%t\n", c.isAlive)
	return c.isAlive
}

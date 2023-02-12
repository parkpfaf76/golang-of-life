package models

import (  
    "fmt"
)

type cell struct {
	posX      	int
	poxY      	int
	isAlive 	bool
 }

 func New(posX string, poxY string, isAlive int) cell {  
    c := cell {posX, poxY, isAlive}
    return c
}

func (c cell) IsAlive() bool {
	fmt.Printf("isAlive=%b\n", c.isAlive)
	return c.isAlive;
}
package main

import (
    "log"
	"github.com/parkpfaf76/golang-of-life/model/cell"
)

func main() {
    log.Println("Parker's game of life in go!")
    c := cell.New(0, 0, false);
    c.IsAlive()
}
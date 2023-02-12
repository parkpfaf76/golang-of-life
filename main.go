package main

import (
    "fmt"
	"github.com/parkpfaf76/golang-of-life/models"
)

func main() {
    fmt.Println("Parker's game of life in go!")
    c := models.NewCell(0, 0, false);
    c.IsAlive()
}
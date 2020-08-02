package main

import (
	"github.com/copter01252/go-module"
	"github.com/copter01252/go-module/controller"
	"fmt"
)

// Cmd ไม่ต้องเขียน test
func main() {
	result := github.com/copter01252/go-module.SayHi("Krit")
	fmt.Println(result)
	controller.CallController()
}

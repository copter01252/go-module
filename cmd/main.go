package main

import (
	"demo"
	"demo/controller"
	"fmt"
)

// Cmd ไม่ต้องเขียน test
func main() {
	result := demo.SayHi("Krit")
	fmt.Println(result)
	controller.CallController()
}

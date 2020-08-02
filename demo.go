package demo

import "github.com/fatih/color"

// SayHi : ddd
func SayHi(name string) string {
	if name == "" {
		return "Error"
	}
	return say(name)
}

func say(name string) string {
	color.Magenta("Prints text in cyan.")
	return "Say Hi " + name
}

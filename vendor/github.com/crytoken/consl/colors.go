package consl

import (
	"fmt"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func PrintRed(a ...interface{}) {
	fmt.Print(ColorRed)
	fmt.Print(a...)
	fmt.Print(ColorReset)
}

func PrintGreen(a ...interface{}) {
	fmt.Print(ColorGreen)
	fmt.Print(a...)
	fmt.Print(ColorReset)
}

func PrintYellow(a ...interface{}) {
	fmt.Print(ColorYellow)
	fmt.Print(a...)
	fmt.Print(ColorReset)
}

func PrintBlue(a ...interface{}) {
	fmt.Print(ColorBlue)
	fmt.Print(a...)
	fmt.Print(ColorReset)
}

func PrintPurple(a ...interface{}) {
	fmt.Print(ColorPurple)
	fmt.Print(a...)
	fmt.Print(ColorReset)
}

func PrintCyan(a ...interface{}) {
	fmt.Print(ColorCyan)
	fmt.Print(a...)
	fmt.Print(ColorReset)
}

func PrintWhite(a ...interface{}) {
	fmt.Print(ColorWhite)
	fmt.Print(a...)
	fmt.Print(ColorReset)
}

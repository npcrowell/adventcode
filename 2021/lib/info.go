package lib

import (
	"bufio"
	"fmt"
	"os"
)

func Printf(f string, a ...interface{}) {
	fmt.Printf(f, a...)
}

func Print(f string, a ...interface{}) {
	Printf(f, a...)
	fmt.Println()
}

func Perror(f string, a ...interface{}) {
	fmt.Print(" [-] Error: ")
	Printf(f, a...)
	fmt.Println()
}

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

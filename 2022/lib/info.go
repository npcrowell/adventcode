package lib

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
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

func Debug(f string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf(" [*] %v:%v >> ", file, line)
	fmt.Printf(f, a...)
	fmt.Println()
}

package lib

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
)

var Dbg bool = false

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
	if Dbg {
		_, fullfile, line, _ := runtime.Caller(1)
		r, _ := regexp.Compile("2023/(.+[.]go)$")
		file := r.FindAllStringSubmatch(fullfile, 1)[0]

		fmt.Printf(" [*] %v:%v >> ", file, line)
		fmt.Printf(f, a...)
		fmt.Println()
	}
}

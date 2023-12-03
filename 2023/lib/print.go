package lib

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"

	"github.com/fatih/color"
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
	c := color.New(color.FgRed)
	c.Print(" [-] Error: ")
	c.Printf(f, a...)
	c.Println()
}

func Psuccess(f string, a ...interface{}) {
	c := color.New(color.FgGreen)
	c.Print(" [+] ")
	c.Printf(f, a...)
	c.Println()
}

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func Debug(f string, a ...interface{}) {
	if Dbg {
		_, fullfile, line, _ := runtime.Caller(1)
		r, _ := regexp.Compile("2023/(.+go)$")
		file := r.FindAllStringSubmatch(fullfile, 1)[0][1]
		c := color.New(color.FgHiYellow)
		m := color.New(color.FgBlue)

		m.Printf(" [%v:%v] ", file, line)
		c.Printf(f, a...)
		c.Println()
	}
}

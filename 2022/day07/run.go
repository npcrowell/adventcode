package day07

import (
	"adventcode/2022/lib"
	e "adventcode/2022/lib"
	"strconv"
	"strings"
)

func buildFileSystem(data []string) *Filesystem {
	fs := Filesystem{}
	fs.tld = newDirectory(strings.Split(data[0], " ")[2], nil)
	fs.cur = fs.tld

	for _, line := range data[1:] {
		lpart := strings.Split(line, " ")

		if lpart[0] == "$" {
			// handle command

			if lpart[1] == "cd" {
				// handle change directory

				if lpart[2] == ".." {
					// current dir = parent dir
					fs.cur = fs.cur.parent

				} else {
					// curent dir = lpart[2]
					newDir := fs.cur.getDirectory(lpart[2])
					if newDir == nil {
						lib.Perror("Directory not found: %v in %v", lpart[2], fs.cur)
						continue
					}
					fs.cur = newDir
				}
			} else if lpart[1] == "ls" {
				//handle list directory
				// lib.Print("Name: %10v Size: %v", fs.tld.name, fs.tld.size)
				continue

			} else {
				lib.Perror("Unrecognized command: %v", line)
			}
			continue

		} else if lpart[0] == "dir" {
			// add dir to current directory

			newDir := newDirectory(lpart[1], fs.cur)
			fs.cur.addDirectory(*newDir)
			fs.addDirectory(*newDir)
			// fmt.Printf("dir: %v\n", newDir.getPath())
			continue
		}

		// add file to current directory
		size, err := strconv.Atoi(lpart[0])
		if err != nil {
			lib.Perror("Unrecognized input: %v", line)
		}
		file := newFile(lpart[1], size, fs.cur)
		fs.cur.addFile(*file)

	}
	// getSumLTECapInDirectory(100000, fs.tld)
	return &fs
}

func traverseDirectory(dir *Directory, indent string, pfiles bool) int {
	// e.Print("%v%v [%v]", indent, dir.name, dir.size)
	sum := 0
	if dir.size <= 100000 {
		e.Print("%v%v [%v]", indent, dir.name, dir.size)
		sum += dir.size
	}

	indent += "    "

	if pfiles {
		for _, f := range dir.flist {
			e.Print("%v%v [%v]", indent, f.name, f.size)
		}
	}

	for _, d := range dir.dir {
		sum += traverseDirectory(d, indent, pfiles)
	}
	return sum
}

func part1(data []string) string {
	fs := buildFileSystem(data)
	sum := traverseDirectory(fs.tld, "", false)
	return strconv.Itoa(sum)
}

func smallestSortDirectory(dir *Directory, indent string, pfiles bool) int {
	// e.Print("%v%v [%v]", indent, dir.name, dir.size)
	sum := 0
	if dir.size >= 8518336 {
		e.Print("%v%v [%v]", indent, dir.name, dir.size)
		sum = dir.size
	}

	indent += "    "

	if pfiles {
		for _, f := range dir.flist {
			e.Print("%v%v [%v]", indent, f.name, f.size)
		}
	}

	for _, d := range dir.dir {

		s := smallestSortDirectory(d, indent, pfiles)
		if s > 0 && s < sum {
			sum = s
		}
	}
	return sum
}

func part2(data []string) string {
	fs := buildFileSystem(data)
	sremain := 70000000 - fs.tld.size
	lib.Print("Size remaining: %v", sremain)
	diff := 30000000 - sremain
	lib.Print("Difference: %v", diff)
	sum := smallestSortDirectory(fs.tld, "", false)
	return strconv.Itoa(sum)
}

func Run(datafile string, part int) {
	e.Print("Day 07 is driving!")

	data, err := e.ReadInTextFile(datafile)
	if err != nil {
		e.Perror("%v", err)
		return
	}

	switch part {
	case 1:
		e.Print("Part 1: %v", part1(data))
	case 2:
		e.Print("Part 2: %v", part2(data))
	default:
		e.Perror("Unregognized part number: %v", part)
	}
}

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {

	rootDirectory := runCommands(parseInput(input))

	// recursively calculate their size
	rootDirectory.calculateSize()

	// recursively find the "large directories"
	largeDirectories := []*Directory{}
	findSizeLimitDir(rootDirectory, &largeDirectories, MaxSize)

	totalSize := 0
	for _, d := range largeDirectories {
		totalSize += d.size
	}
	return totalSize
}

func part2(input string) int {

	rootDirectory := runCommands(parseInput(input))
	rootDirectory.calculateSize()

	NeededSpace := 3e7 - (7e7 - rootDirectory.size)
	dirToRemove := newDirectory("dummy directory")
	dirToRemove.size = 7e7 // this seems clunky

	dirToRemove = findDeletionCandidate(rootDirectory, dirToRemove, NeededSpace)

	return dirToRemove.size
}

// commands begin with $..
// we will iterate until the next $ and append results to an output
func parseInput(input string) History {
	var history History
	for _, line := range strings.Split(input, "\n") {
		if line[0] == '$' {
			history = append(history, Command{input: strings.Replace(line, "$ ", "", 1)})
		} else {
			history[len(history)-1].output = append(history[len(history)-1].output, line)
		}
	}
	return history
}

type Command struct {
	input  string
	output []string
}

type History []Command

type Directory struct {
	size        int
	name        string
	parent      *Directory
	files       []File
	directories map[string]*Directory
}

type File struct {
	name string
	size int
}

const MaxSize = 100000

func (d *Directory) calculateSize() int {
	var totalSize int
	for _, f := range d.files {
		totalSize += f.size
	}
	for _, d := range d.directories {
		totalSize += d.calculateSize()
	}
	d.size = totalSize
	return totalSize
}

func newDirectory(name string) *Directory {
	return &Directory{name: name, directories: make(map[string]*Directory)}
}

func newFile(line string) File {
	s := strings.Split(line, " ")
	size, err := strconv.Atoi(s[0])
	if err != nil {
		panic(err)
	}
	return File{name: s[1], size: size}
}

func findSizeLimitDir(d *Directory, largeDirectories *[]*Directory, sizeLimit int) {
	if d.size < sizeLimit {
		*largeDirectories = append(*largeDirectories, d)
	}
	for _, d := range d.directories {
		findSizeLimitDir(d, largeDirectories, sizeLimit)
	}
}

func findDeletionCandidate(d *Directory, dirToRemove *Directory, NeededSpace int) *Directory {
	if d.size > NeededSpace && d.size < dirToRemove.size {
		dirToRemove = d
	}
	for _, d := range d.directories {
		dirToRemove = findDeletionCandidate(d, dirToRemove, NeededSpace)
	}
	return dirToRemove
}

func populateDirectory(dir *Directory, fileList []string) {
	for _, line := range fileList {
		if strings.HasPrefix(line, "dir") {
			dirName := strings.Split(line, " ")[1]
			dir.directories[dirName] = newDirectory(dirName)
		} else {
			dir.files = append(dir.files, newFile(line))
		}
	}
}

func runCommands(history History) *Directory {
	// let's assume we always have / as the top most directory
	var rootDirectory = newDirectory("/")
	var currentDirectory *Directory
	for _, command := range history {
		program := strings.Split(command.input, " ")[0]

		switch program {
		case "cd":
			targetDirectoryName := strings.Split(command.input, " ")[1]
			switch targetDirectoryName {
			case "/":
				currentDirectory = rootDirectory
			case "..":
				currentDirectory = currentDirectory.parent
			default:
				oldDirectory := currentDirectory
				currentDirectory = currentDirectory.directories[targetDirectoryName]
				currentDirectory.parent = oldDirectory
			}
		case "ls":
			populateDirectory(currentDirectory, command.output)
		}
	}
	return rootDirectory
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	cli-app $ make << fname >> fext
		crates a file with the name "fname" and the  extension "fext"
		the file created is: fname.fext

	cli-app $ make << fname >> [fext1,fext2]
*/
// function to make a new file with a specific file type or a list of types
func newFile(fname, fext string) {
	make := func(fname, fext string) {
		file := strings.Join([]string{fname, fext}, ".")
		os.Create(file)
		fmt.Printf("cli-app $ made %v\n", file)
	}

	if fext[0] == '[' {
		var exts []string
		fext = strings.TrimPrefix(strings.TrimSuffix(fext, "]"), "[")
		exts = strings.Split(fext, ",")

		for ext := range exts {
			make(fname, exts[ext])
		}
	} else {
		make(fname, fext)
	}
}

// function to exit the cli-app
func cliExt() {
	os.Exit(0)
}

// function that shows the information of a specific file
func show(fname string) {
	
}

func main() {
	goin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("cli-app $ ")
		input, err := goin.ReadString('\n')

		if err != nil {
			fmt.Println("Error. Cannot scan input: ", err)
			continue
		}

		input = strings.TrimSpace(input)
		split := strings.Fields(input)

		switch split[0] {
		case "sysexit":
			cliExt()
		case "make":
			if len(split) >= 4 {
				newFile(split[2], split[4])
			} else {
				fmt.Println("Error. Argument overflow or underflow for make. Command has", len(split), "arguments, expected >= 4.")
			}
		case "show":
			if len(split) >= 2{
				show(split[2])
			} else{
				fmt.Println("Error. Argument overflow or underflow for show. Command has", len(split), "arguments, expected >= 2".)
			}
		default:
			fmt.Println("Error. Unknown input:", input)
		}
	}
}

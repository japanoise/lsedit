package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s file\n", os.Args[0])
		os.Exit(1)
	}
	editor, err := createEditorFromFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	welcome()
	editor.printCurRow()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 1 && line[0] == '.' {
			if line[1] == '.' {
				editor.insert(line[1:])
			} else {
				com := parse(line)
				if editor.exec(com) {
					os.Exit(0)
				}
			}
		} else {
			editor.insert(line)
		}
	}
}

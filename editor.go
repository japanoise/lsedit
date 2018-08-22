package main

import (
	"bufio"
	"fmt"
	"os"
)

type lseditor struct {
	filename string
	rows     []string
	currow   int
}

func (l *lseditor) printCurRow() {
	fmt.Printf("< Insert at line %d >\n", l.currow+1)
}

func createEditorFromFile(filename string) (*lseditor, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	ret := new(lseditor)
	ret.filename = filename
	ret.currow = 0
	ret.rows = make([]string, 0, 100)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ret.rows = append(ret.rows, scanner.Text())
		ret.currow++
	}
	return ret, nil
}

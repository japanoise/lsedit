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
	numrows  int
}

func (l *lseditor) printCurRow() {
	fmt.Printf("< Insert at line %d >\n", l.currow+1)
}

func (l *lseditor) insert(line string) {
	l.numrows++
	l.currow++
	if l.currow == l.numrows {
		l.rows = append(l.rows, line)
	} else {
		i := l.currow - 1
		l.rows = append(l.rows, "")
		copy(l.rows[i+1:], l.rows[i:])
		l.rows[i] = line
	}
}

func (l *lseditor) save() error {
	file, err := os.Create(l.filename)
	if err != nil {
		return err
	}
	lines := 0
	defer file.Close()
	for _, row := range l.rows {
		fmt.Fprintln(file, row)
		lines++
	}
	fmt.Printf("< Wrote %d lines to %s >\n", lines, l.filename)
	return nil
}

func printLinum(num int) {
	fmt.Printf("%4d ", num)
}

// Returns true if exit
func (l *lseditor) exec(com *command) bool {
	linum := false
	switch com.name {
	case ".abort":
		return true
	case ".end":
		err := l.save()
		if err != nil {
			fmt.Fprintf(os.Stderr, "< %s >\n", err.Error())
			return false
		}
		return true
	case ".save":
		err := l.save()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	case ".h", ".help":
		helpScreen()
	case ".i", ".insert":
		if com.index1 == NoIndex || com.index1 == CurLine {
			// Do nothing; just print the current row number
		} else if com.index1 == LastLine || com.index1 > l.numrows {
			l.currow = l.numrows
		} else {
			l.currow = com.index1 - 1
		}
		l.printCurRow()
	case ".p", ".print":
		linum = true
		fallthrough
	case ".l", ".list":
		if com.index1 == NoIndex {
			for i, row := range l.rows {
				if linum {
					printLinum(i + 1)
				}
				fmt.Println(row)
			}
			return false
		} else if com.index1 == LastLine || com.index1 == l.numrows {
			if linum {
				printLinum(l.numrows)
			}
			fmt.Println(l.rows[l.numrows-1])
			return false
		} else if com.index1 > l.numrows {
			return false
		}

		i1 := com.index1
		i2 := com.index2
		if com.index1 == CurLine {
			i1 = l.currow
		}
		if com.index2 == CurLine {
			i2 = l.currow
		}
		if i2 == LastLine || i2 >= l.numrows {
			for i, row := range l.rows[i1-1:] {
				if linum {
					printLinum(i1 + i)
				}
				fmt.Println(row)
			}
		} else if i2 == NoIndex || i2 <= com.index1 {
			if linum {
				printLinum(i1)
			}
			fmt.Println(l.rows[i1-1])
		} else {
			for i, row := range l.rows[i1-1 : i2] {
				if linum {
					printLinum(i1 + i)
				}
				fmt.Println(row)
			}
		}
	}
	return false
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
	ret.numrows = ret.currow
	return ret, nil
}

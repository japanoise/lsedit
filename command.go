package main

import "strings"

// Special indices
const (
	NoIndex int = 0 - iota
	LastLine
	CurLine
)

type command struct {
	name   string
	index1 int
	index2 int
	equals string
}

func parse(line string) *command {
	name := ""
	idx := 0
	eos := len(line)
	for idx < eos && line[idx] != ' ' {
		name += string(line[idx])
		idx++
	}
	idx++
	index1 := NoIndex
	index2 := NoIndex
	cin := &index1

	for idx < eos {
		if line[idx] == '=' {
			return &command{strings.ToLower(strings.TrimSpace(name)), index1, index2,
				strings.TrimPrefix(strings.TrimSpace(line[idx:]), "=")}
		}
		if line[idx] == ' ' {
			cin = &index2
		} else if '0' <= line[idx] && line[idx] <= '9' {
			if *cin < 0 {
				*cin = 0
			}
			*cin *= 10
			*cin += int(line[idx] - 0x30)
		} else if line[idx] == '$' {
			*cin = LastLine
		} else if line[idx] == '.' {
			*cin = CurLine
		} else if line[idx] == '^' {
			*cin = 1
		} else {
			return nil
		}
		idx++
	}
	return &command{strings.ToLower(strings.TrimSpace(name)), index1, index2, ""}
}

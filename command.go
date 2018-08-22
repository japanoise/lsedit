package main

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
	for line[idx] != ' ' && idx < eos {
		name += string(line[idx])
		idx++
	}
	idx++
	index1 := -1
	index2 := -1
	cin := &index1

	for idx < eos {
		if line[idx] == '=' {
			return &command{name, index1, index2, line[idx:]}
		}
		if line[idx] == ' ' {
			cin = &index2
		} else if '0' <= line[idx] && line[idx] <= '9' {
			*cin *= 10
			*cin += int(line[idx] - 0x30)
		} else {
			return nil
		}
	}
	return &command{name, index1, index2, ""}
}

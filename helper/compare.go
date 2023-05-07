package helper

import "strings"

func ListContain(con string, l []string) bool {
	for _, c := range l {
		if c == con {
			return true
		}
	}
	return false
}

func ListSubContain(con string, l []string) bool {
	for _, c := range l {
		if strings.HasSuffix(con, c) {
			return true
		}
	}
	return false
}

package common

import (
	"fmt"
)

type String string

func (str String) Println() {
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	for _, value := range str {
		fmt.Print(rot13(value))
	}
}

// func (str String) String(tmp string) {
//     string
// }

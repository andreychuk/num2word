package num2word

import (
	"fmt"
	"strings"
	"unicode"
)

var mask = []string{",,,", ",,", ",", ",,,,", ",,", ",", ",,,,,", ",,", ",", ",,,,,,", ",,", ","}

// Num2word - гроші прописом
func Num2word(number float64, upperFirstChar bool, repl [][]string) string {

	s := fmt.Sprintf("%.2f", number)
	l := len(s)

	dest := s[l-3:l] + "k"
	s = s[:l-3]
	l = len(s)
	for i := l; i > 0; i-- {
		c := string(s[i-1])
		dest = c + mask[l-i] + dest
	}

	for _, r := range repl {
		dest = strings.Replace(dest, r[0], r[1], -1)
	}
	if upperFirstChar {
		a := []rune(dest)
		a[0] = unicode.ToUpper(a[0])
		dest = string(a)
	}
	return dest
}

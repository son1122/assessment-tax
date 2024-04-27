package util

import (
	"bytes"
	"fmt"
)

// function to format numbers with commas
func commaFormat(n int) string {
	num := fmt.Sprintf("%d", n)
	var buffer bytes.Buffer
	minus := ""
	if num[0] == '-' {
		minus = "-"
		num = num[1:]
	}
	l := len(num)
	for i, digit := range num {
		if i > 0 && (l-i)%3 == 0 {
			buffer.WriteString(",")
		}
		buffer.WriteRune(digit)
	}
	return minus + buffer.String()
}

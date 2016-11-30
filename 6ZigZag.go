package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	length := len(s)
	if numRows == 1 || length == 0 {
		return s
	}
	var buffer bytes.Buffer
	for i := 0; i < numRows; i++ {
		for j := i; j < length; j += 2 * (numRows - 1) {
			buffer.WriteByte(s[j])
			if i != 0 && i != numRows-1 && 2*(numRows-i-1)+j < length {
				buffer.WriteByte(s[2*(numRows-i-1)+j])
			}
		}
	}
	return buffer.String()

}
func main() {
	s := "123456789ab"
	fmt.Println(convert(s, 3))
	fmt.Println(convert(s, 4))
}

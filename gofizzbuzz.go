package gofizzbuzz

import (
	"strconv"
)

func fbconv(i int) (out string) {
	if i%3 == 0 {
		out = out + "fizz"
	}

	if i%5 == 0 {
		out = out + "buzz"
	}

	if out == "" {
		out = strconv.Itoa(i)
	}

	return
}

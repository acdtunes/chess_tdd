package numbernames

import (
	"fmt"
	"strconv"
	"strings"
)

var names = map[int]string {
	0: "zero",
	1:"one",
	2:"two",
	3:"three",
	4:"four",
	5:"five",
	6:"six",
	7:"seven",
	8:"eight",
	9:"nine",
	10:"ten",
	11:"eleven",
	12:"twelve",
	13:"thirteen",
	14:"fourteen",
	15:"fifteen",
	16:"sixteen",
	17:"seventeen",
	18:"eighteen",
	19:"nineteen",
	20:"twenty",
	30:"thirty",
	40:"forty",
	50:"fifty",
	60:"sixty",
	70:"seventy",
	80:"eighty",
	90:"ninety",
	100:"hundred",
	1000:"thousand",
	1000000:"million",
}

func Name(n int) string {
	if val, ok := names[n]; ok && n < 100 {
		return val
	}

	if n < 100 {
		tens := n / 10
		units := n % 10
		return fmt.Sprintf("%s %s", names[10*tens], names[units])
	}

	power := len(strconv.Itoa(n)) - 1
	magnitude, _ := strconv.Atoi("1" + strings.Repeat("0", power))
	main := n / magnitude
	remaining := Name(n % magnitude)

	name := names[magnitude]

	if main > 1 && magnitude > 1000 {
		name += "s"
	}

	if remaining == "zero" {
		return fmt.Sprintf("%s %s", Name(main), name)
	}

	separator := " and"
	if magnitude > 100 {
		separator = ","
	}

	return fmt.Sprintf("%s %s%s %s", Name(main), name, separator, remaining)
}

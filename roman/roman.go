package roman

import (
	"sort"
)

var numerals = map[int]string {
	1 : "I",
	5 : "V",
	10 : "X",
	50 : "L",
	100 : "C",
	500 : "D",
	1000: "M",
	5000: "V2",
	10000: "X2",
	50000: "L2",
	100000: "C2",
	500000: "D2",
	1000000: "M2",
}

func Numeral(n int) string {
	if val, ok := numerals[n]; ok {
		return val
	}

	completeNumerals()

	keys := descendingKeys()
	for _, k := range keys {
		if n >= k {
			return Numeral(k) + Numeral(n-k)
		}
	}

	return ""
}

func completeNumerals() {
	for i := 1; i <= 10000; i=i*10 {
		addPrecedingNumeral(i, i*5)
		addPrecedingNumeral(i, i*10)
	}
}

func addPrecedingNumeral(i int, j int) {
	numerals[j - i] = numerals[i] + numerals[j]
}

func descendingKeys() []int {
	keys := make([]int, 00, len(numerals))
	for k := range numerals {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})

	return keys
}
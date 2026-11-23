package main

import "fmt"

// MRZ check digits as ICAO Doc 9303 defines them. No dependencies.

var weights = [3]int{7, 3, 1}

// charValue is the value of one MRZ character: 0-9 as themselves, A-Z as 10-35, '<' as 0.
func charValue(c byte) (int, error) {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0'), nil
	case c >= 'A' && c <= 'Z':
		return int(c-'A') + 10, nil
	case c == '<':
		return 0, nil
	}
	return 0, fmt.Errorf("not an MRZ character: %q", c)
}

// CheckDigit is the check digit of a field: each value times 7, 3, 1 repeating, summed, mod 10.
func CheckDigit(data string) (int, error) {
	sum := 0
	for i := 0; i < len(data); i++ {
		v, err := charValue(data[i])
		if err != nil {
			return 0, err
		}
		sum += v * weights[i%3]
	}
	return sum % 10, nil
}

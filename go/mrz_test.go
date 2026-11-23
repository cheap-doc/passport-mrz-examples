package main

import "testing"

// Fields of the ICAO Doc 9303 specimen passport line.
const line = "L898902C36UTO7408122F1204159ZE184226B<<<<<10"

func TestCheckDigit(t *testing.T) {
	cases := []struct {
		name, data string
		want       int
	}{
		{"document number", "L898902C3", 6},
		{"birth date", "740812", 2},
		{"expiry date", "120415", 9},
		{"personal number", "ZE184226B<<<<<", 1},
		{"composite", line[0:10] + line[13:20] + line[21:43], 0},
		{"fillers count as zero", "<<<<<<", 0},
	}
	for _, c := range cases {
		got, err := CheckDigit(c.data)
		if err != nil || got != c.want {
			t.Errorf("%s: CheckDigit(%q) = %d, %v; want %d", c.name, c.data, got, err, c.want)
		}
	}
}

func TestOneMisreadCharacterChangesTheDigit(t *testing.T) {
	if got, _ := CheckDigit("L898902C4"); got == 6 {
		t.Errorf("a misread document number still gave 6")
	}
}

func TestRefusesACharacterOutsideTheAlphabet(t *testing.T) {
	if _, err := CheckDigit("l898902c3"); err == nil {
		t.Errorf("lower-case input was accepted")
	}
}

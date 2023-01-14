package romannumerals

import "testing"

func TestIntegerToRoman(t *testing.T) {
	cases := []struct {
		num int
		str string
	}{
		{
			num: 1,
			str: "I",
		},
		{
			num: 2,
			str: "II",
		},
		{
			num: 3,
			str: "III",
		},
		{
			num: 4,
			str: "IV",
		},
		{
			num: 5,
			str: "V",
		},
		{
			num: 6,
			str: "VI",
		},
		{
			num: 7,
			str: "VII",
		},
		{
			num: 8,
			str: "VIII",
		},
		{
			num: 9,
			str: "IX",
		},
		{
			num: 10,
			str: "X",
		},
	}

	for _, c := range cases {
		if want, got := c.str, IntegerToRoman(c.num); want != got {
			t.Errorf("got %s for number %d, want %s", got, c.num, want)
		}
	}
}

package parser

import (
	"testing"
)

func TestFromStrLit(t *testing.T) {
	cases := []struct {
		lit       string
		lastToken int
		expect    int
	}{
		{
			"10",
			0,
			decimal_value,
		},
		{
			"0xff",
			0,
			hex_value,
		},
		// Not hex value
		{
			"ff",
			0,
			0,
		},
	}

	tk := keywordTokenizer{}
	for _, c := range cases {
		actual := tk.FromStrLit(c.lit, c.lastToken)
		if actual != c.expect {
			t.Errorf("Expected: %v, but actual: %v\n", c.expect, actual)
		}
	}
}

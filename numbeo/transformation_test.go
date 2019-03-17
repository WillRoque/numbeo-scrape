package numbeo

import "testing"

func TestGetFloatString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"24.00 Fr.", "24.00"},
		{"8.50 €.", "8.50"},
		{"18.00 R$.", "18.00"},
		{"12.00 $", "12.00"},
	}
	for _, c := range cases {
		got := getFloatString(c.in)
		if got != c.want {
			t.Errorf("getFloatString(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

package strings

import (
	"fmt"
	"testing"
)

func TestReserveWords(t *testing.T) {
	var tests = []struct {
		sentence string
		want     string
	}{
		{"i.like.this.program.very.much", "much.very.program.this.like.i"},
		{"pqr.mno", "mno.pqr"},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf(tt.sentence)
		t.Run(testName, func(t *testing.T) {
			ans := reverseWords(tt.sentence)
			if ans != tt.want {
				t.Errorf("Got %s, want: %s", ans, tt.want)
			}
		})
	}
}

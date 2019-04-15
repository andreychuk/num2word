package num2word

import (
	"testing"
)

var uaSamples = []struct {
	amount   float64
	upper    bool
	expected string
}{
	{1, true, "Одна гривня 00 копійок"},
	{100.21, false, "сто гривень 21 копійка"},
	{8010.00, false, "вісім тисячь десять гривень 00 копійок"},
}

func Test_UaNum2word(t *testing.T) {
	for _, tt := range uaSamples {
		res := UaNum2word(tt.amount, tt.upper)
		if res != tt.expected {
			t.Errorf("UaMoney(%.2f): expected '%s', got '%s'", tt.amount, tt.expected, res)
		}
	}
}

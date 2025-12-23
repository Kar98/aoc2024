package day2

import (
	"testing"
)

func TestExample(t *testing.T) {
	exp := 2
	got := Example()
	if exp != got {
		t.Errorf("expected '%d' got '%d'", exp, got)
	}
}

func TestSameNumber(t *testing.T) {
	report := "44 44 40 38 35 38"
	rep := ToNums(report)
	got := ValidReport(rep)
	if got == true {
		t.Error("should be false")
	}
}

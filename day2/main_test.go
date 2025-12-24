package day2

import (
	"testing"
)

func TestExampleV1(t *testing.T) {
	exp := 2
	got := Example(1)
	if exp != got {
		t.Errorf("expected '%d' got '%d'", exp, got)
	}
}

func TestExampleV2(t *testing.T) {
	exp := 4
	got := Example(2)
	if exp != got {
		t.Errorf("expected '%d' got '%d'", exp, got)
	}
}

func TestSameNumber(t *testing.T) {
	report := "8 6 4 4 1"
	rep := ToNums(report)
	got := ValidReport(rep, 1)
	if got == true {
		t.Error("should be false")
	}
}

func TestSafeReportsV2(t *testing.T) {
	reports := []string{"1 3 2 4 5", "7 6 4 2 1", "8 6 4 4 1", "1 3 6 7 9"} // all should be safe
	for _, report := range reports {
		rep := ToNums(report)
		got := ValidReport(rep, 2)
		if got != true {
			t.Errorf("%s should be true", report)
		}
	}

}

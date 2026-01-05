package day3

import (
	"strings"
	"testing"
)

func TestExampleDay1(t *testing.T) {

	result := Example(1)

	if result != 161 {
		t.Error("not 161")
	}
}

func TestEdgeCasePart1(t *testing.T) {
	case1 := "#mul(886,236}how()mul(674,816)!"
	output := StartBuffering(case1)

	if output[0] != "mul(674,816)" {
		t.Error(output)
	}
}

func TestPart1(t *testing.T) {
	output := Main(1)

	if output != 167090022 {
		t.Error("not 167090022")
	}
}

func TestDoDontPart2(t *testing.T) {
	operations := BufferingWithDoDont(example2)
	if strings.Join(operations, " ") != "mul(2,4) mul(8,5)" {
		t.Logf("got %s", strings.Join(operations, ""))
		t.Error(operations)
	}
}

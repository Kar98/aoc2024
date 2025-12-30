package day3

import "testing"

func TestExampleDay1(t *testing.T) {
	// example to equal 161

	result := Example(1)

	if result != 161 {
		t.Error("not 161")
	}
}

func TestEdgeCaseDay1(t *testing.T) {
	case1 := "#mul(886,236}how()mul(674,816)!"
	output := StartBuffering(case1)

	if output[0] != "mul(674,816)" {
		t.Error(output)
	}

}

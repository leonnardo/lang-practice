package day1_aoc

import "testing"

func TestResultWithZeroSum(t *testing.T) {
	arr := convert_to_int_array("1234")
	sum := calculate(arr)
	if sum != 0 {
		t.Errorf("Incorret sum, got: %d and want %d", sum, 0)
	}
}

func TestResultWithSumThree(t *testing.T) {
	value := 3
	arr := convert_to_int_array("1122")
	sum := calculate(arr)
	if sum != value {
		t.Errorf("Incorret sum, got: %d and want %d", sum, value)
	}
}

func TestResultWithSumFour(t *testing.T) {
	value := 4
	arr := convert_to_int_array("1111")
	sum := calculate(arr)
	if sum != value {
		t.Errorf("Incorret sum, got: %d and want %d", sum, value)
	}
}

func TestResultWithSumNine(t *testing.T) {
	value := 9
	arr := convert_to_int_array("91212129")
	sum := calculate(arr)
	if sum != value {
		t.Errorf("Incorret sum, got: %d and want %d", sum, value)
	}
}

package calculator

import (
	"reflect"
	"testing"
)

//func TestCalculate(t *testing.T){
//	add := Calculate(...string)
//	if add != 16{
//		t.Error("Expected: " ,16 ,"Got: ", add)
//	}
//}

func TestCalculate(t *testing.T) {
	var tests = []struct {
		input []string
		want []float64
	}{
		{[]string{"3*4*5","5-2-1"},[]float64{60,2}},
		{[]string{"9/2/2","8+2+9"},[]float64{2.25,19}},
		{[]string{"98+7+34+6", "9*23*4", "9-19-3","13/2"},[]float64{145,828,-13,6.5}},
	}
	for _, test := range tests {
		if got := Calculate(test.input...); !reflect.DeepEqual(got, test.want){
			t.Errorf("Calculate(%v) = %v", test.input, got)
		}
	}
}


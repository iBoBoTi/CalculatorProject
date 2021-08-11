package operations



import "testing"

func TestAdd(t *testing.T) {
	var tests = []struct {
		input []float64
		want float64
	}{
		{[]float64{3,4,5},float64(12)},
		{[]float64{9,5},float64(14)},
		{[]float64{98,17,34,76},float64(225)},
		{[]float64{1},float64(1)},
		{[]float64{7,9},float64(16)},
		{[]float64{8,9,9},float64(26)},
	}
	for _, test := range tests {
		if got := Add(test.input); got != test.want {
			t.Errorf("Add(%v) = %v", test.input, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	var tests = []struct {
		input []float64
		want float64
	}{
		{[]float64{3,4,5},float64(-6)},
		{[]float64{9,5},float64(4)},
		{[]float64{98,17,34,76},float64(-29)},
		{[]float64{1},float64(1)},
		{[]float64{7,9},float64(-2)},
		{[]float64{99,9,9},float64(81)},
	}
	for _, test := range tests {
		if got := Subtract(test.input); got != test.want {
			t.Errorf("Subtract(%v) = %v", test.input, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	var tests = []struct {
		input []float64
		want float64
	}{
		{[]float64{3,4,5},float64(60)},
		{[]float64{9,5},float64(45)},
		{[]float64{98,17,34,76},float64(4304944)},
		{[]float64{1},float64(1)},
		{[]float64{7,9},float64(63)},
		{[]float64{8,0},float64(0)},
	}
	for _, test := range tests {
		if got := Multiply(test.input); got != test.want {
			t.Errorf("Multiply(%v) = %v", test.input, got)
		}
	}
}

func TestDivide(t *testing.T) {
	var tests = []struct {
		input []float64
		want float64
	}{
		{[]float64{19,4},4.75},
		{[]float64{9,5},1.8},
		{[]float64{1},float64(1)},
		{[]float64{9,3,3},float64(1)},
	}
	for _, test := range tests {
		if got := Divide(test.input); got != test.want {
			t.Errorf("Divide(%v) = %v", test.input, got)
		}
	}
}

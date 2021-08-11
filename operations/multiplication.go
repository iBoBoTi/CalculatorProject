package operations

func Multiply(arr []float64) float64 {
	// initiate a multiple to a zero value
	// loop through the incoming array
	// initiate mul to 1
	// return the multiple which should be a float64 type
	var mul float64
	mul = 1
	for _,j := range arr{
		mul *= j
	}
	return mul
}

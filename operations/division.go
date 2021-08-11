package operations

func Divide(arr []float64) float64 {
	// initiate a dividend to a zero value
	// loop through the incoming array
	// initiate the first value to the first item of the array
	// return the dividend which should be a float64 type
	var div float64
	div = arr[0]
	for i:=1; i<len(arr); i++{
		div/=arr[i]
	}
	return div
}

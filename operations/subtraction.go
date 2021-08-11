package operations

func Subtract(arr []float64) float64 {
	// initiate a difference to a zero value
	// loop through the incoming array
	// initiate diff to the first item
	// return the difference which should be a float64 type
	var diff float64
	diff = arr[0]
	for i:=1; i<len(arr); i++{
		diff-=arr[i]
	}
	return diff
}

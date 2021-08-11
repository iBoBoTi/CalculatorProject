package operations


func Add(arr []float64) float64 {
	// initiate a sum to a zero value
	// loop through an array and pick out the sum
	// initiate the add to zero
	// return the sum
	var sum float64
	for _,j := range arr{
		sum+=j
	}
	return sum

}

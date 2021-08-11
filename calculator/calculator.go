package calculator

// Calculator package has a calculate function that takes in a variable number of strings and returns a collection of float64
// it is responsible for checking the input, converting the input to the needed type for the individual
// operation functions to consume them and return a float which is added to the result of calculate function

import (
	"github.com/iBoBoTi/CalculatorProject/operations"
	"strconv"
	"strings"
)

func Calculate(str ...string) []float64 {
	var finalAnswer []float64
	for _, j := range str {
		// check for addition symbol
		if strings.Contains(j, "+") {
			// split the looped strings based of the symbol
			addStrArr := strings.Split(j, "+")
			var addArr []float64

			// loop through the array and convert to integer
			for _, v := range addStrArr {
				if num, err := strconv.ParseFloat(v, 64); err == nil {
					addArr = append(addArr, num)
				}
			}
			finalAnswer = append(finalAnswer, operations.Add(addArr))
		}
		// check for subtraction symbol
		if strings.Contains(j, "-") {
			// split the looped strings based of the symbol
			diffStrArr := strings.Split(j, "-")
			var diffArr []float64

			// loop through the array and convert to integer
			for _, v := range diffStrArr {
				if num, err := strconv.ParseFloat(v, 64); err == nil {
					diffArr = append(diffArr, num)
				}
			}
			//finalAnswer = append(finalAnswer,subtract(diffArr))
			finalAnswer = append(finalAnswer, operations.Subtract(diffArr))
		}
		// check for multiplication symbol
		if strings.Contains(j, "*") {
			// split the looped strings based of the symbol
			mulStrArr := strings.Split(j, "*")
			var mulArr []float64

			// loop through the array and convert to integer
			for _, v := range mulStrArr {
				if num, err := strconv.ParseFloat(v, 64); err == nil {
					mulArr = append(mulArr, num)
				}
			}
			finalAnswer = append(finalAnswer, operations.Multiply(mulArr))
		}

		// check for division function
		if strings.Contains(j, "/") {
			// split the looped strings based of the symbol
			divStrArr := strings.Split(j, "/")
			var divArr []float64

			// loop through the array and convert to integer
			for _, v := range divStrArr {
				if num, err := strconv.ParseFloat(v, 64); err == nil {
					divArr = append(divArr, num)
				}
			}
			finalAnswer = append(finalAnswer, operations.Divide(divArr))
		}
	}
	return finalAnswer
}
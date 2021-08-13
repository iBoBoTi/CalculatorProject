package main

import (
	"fmt"
	"github.com/iBoBoTi/CalculatorProject/calculator"
)

//	ALGORITHM
// 1. Create main function to collect and print the calculate function
// 2. Create a calculate function (-- variadic function) that takes in any number of strings and returns the results of each strings
// in an array
//--- input = series of strings with no fixed number
//--- output = a slice(float64) of a result from the individual strings
//			i. initialize the resulting slice of float64
//			i. loop through individual strings
//			ii. test for the symbol present
//			iii. split the individual strings based off the symbol
//			iv. loop through the array of strings created
//			v. convert each looped strings to float64 number while putting it in a slice
//			vi. apply the needed function to the float
// 3. Create the individual functions for the operations (multiplication,division,addition, subtraction) each collects
//    the float array and works on to give a final answer of a particular float

func main(){
	fmt.Println(calculator.Calculate( "2*3*4*5","3*2*2*2" , "25/5/2", "2+3+4+5.9+6+7.8", "20.54-7.65-2.897"))
}

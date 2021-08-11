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

//func calculate(str ...string) []float64{
//	var finalAnswer []float64
//	for _,j := range str{
//		// test for symbol
//		if strings.Contains(j,"+"){
//			// split the looped strings based of the symbol
//			addStrArr := strings.Split(j,"+")
//			var addArr []float64
//
//			// loop through the array and convert to integer
//			for _,v := range addStrArr{
//				if num, err := strconv.ParseFloat(v, 64); err == nil {
//					addArr = append(addArr,num)
//				}
//			}
//			finalAnswer = append(finalAnswer,operations.Add(addArr))
//		}
//		if strings.Contains(j,"-"){
//			// split the looped strings based of the symbol
//			diffStrArr := strings.Split(j,"-")
//			var diffArr []float64
//
//			// loop through the array and convert to integer
//			for _,v := range diffStrArr{
//				if num, err := strconv.ParseFloat(v, 64); err == nil {
//					diffArr = append(diffArr,num)
//				}
//			}
//			//finalAnswer = append(finalAnswer,subtract(diffArr))
//			finalAnswer = append(finalAnswer,operations.Subtract(diffArr))
//		}
//		if strings.Contains(j,"*"){
//			// split the looped strings based of the symbol
//			mulStrArr := strings.Split(j,"*")
//			var mulArr []float64
//
//			// loop through the array and convert to integer
//			for _,v := range mulStrArr{
//				if num, err := strconv.ParseFloat(v, 64); err == nil {
//					mulArr = append(mulArr,num)
//				}
//			}
//			finalAnswer = append(finalAnswer,operations.Multiply(mulArr))
//		}
//		if strings.Contains(j,"/"){
//			// split the looped strings based of the symbol
//			divStrArr := strings.Split(j,"/")
//			var divArr []float64
//
//			// loop through the array and convert to integer
//			for _,v := range divStrArr{
//				if num, err := strconv.ParseFloat(v, 64); err == nil {
//					divArr = append(divArr,num)
//				}
//			}
//			finalAnswer = append(finalAnswer,operations.Divide(divArr))
//		}
//	}
//	return finalAnswer
//}

//func add(arr []float64) float64 {
//	// initiate a sum to a zero value
//	// loop through an array and pick out the sum
//	// initiate the add to zero
//	// return the sum
//	var sum float64
//	for _,j := range arr{
//		sum+=j
//	}
//	return sum
//
//}
//
//func subtract(arr []float64) float64 {
//	// initiate a difference to a zero value
//	// loop through the incoming array
//	// initiate diff to the first item
//	// return the difference which should be a float64 type
//	var diff float64
//	diff = arr[0]
//	for i:=1; i<len(arr); i++{
//		diff-=arr[i]
//	}
//	return diff
//}
//
//func multiply(arr []float64) float64 {
//	// initiate a multiple to a zero value
//	// loop through the incoming array
//	// initiate mul to 1
//	// return the multiple which should be a float64 type
//	var mul float64
//	mul = 1
//	for _,j := range arr{
//		mul *= j
//	}
//	return mul
//}
//
//func divide(arr []float64) float64 {
//	// initiate a dividend to a zero value
//	// loop through the incoming array
//	// initiate the first value to the first item of the array
//	// return the dividend which should be a float64 type
//	var div float64
//	div = arr[0]
//	for i:=1; i<len(arr); i++{
//		div/=arr[i]
//	}
//	return div
//}

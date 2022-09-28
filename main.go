package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func add(n1 float64, n2 float64) float64 {
	return n1 + n2
}

func multiply(n1 float64, n2 float64) float64 {
	return n1 * n2
}

func subtract(n1 float64, n2 float64) float64 {
	return n1 - n2
}

func divide(n1 float64, n2 float64) float64 {
	return n1 / n2
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	first_number, _ := getInput("Input the first integer", reader)

	fmt.Printf("Choose any operator, (+,-,*,/)")
	opt, _ := getInput("Choose an operator", reader)

	second_number, _ := getInput("Input the second integer", reader)

	answer := float64(0)

	fn1, err := strconv.ParseFloat(first_number, 64)
	if err != nil {
		fmt.Println("Inputs must be a number")
		main()
	}
	fn2, err := strconv.ParseFloat(second_number, 64)
	if err != nil {
		fmt.Println("Inputs must be a number again")
		main()
	}

	switch opt {
	case "+":
		answer = add(fn1, fn2)
		fmt.Printf("The final answer is : %0.1f \n", answer)
	case "-":
		answer = subtract(fn1, fn2)
		fmt.Printf("The final answer is : %0.1f \n", answer)

	case "*":
		answer = multiply(fn1, fn2)
		fmt.Printf("The final answer is : %0.1f \n", answer)

	case "/":
		answer = divide(fn1, fn2)
		fmt.Printf("The final answer is : %0.1f \n", answer)
	default:
		fmt.Println("Operator not available")
		main()

	}

}

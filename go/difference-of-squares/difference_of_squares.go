package diffsquares

// SquareOfSum takes a list of integer values and sums them
// then returns the result of squaring the sum
func SquareOfSum(input int) int {
	result := 0

	for i := 0; i <= input; i++ {
		result += i
	}

	return result * result

}

// SumOfSquares takes a list of integer values and squares each
// returning the sum of all squared values
func SumOfSquares(input int) (sum int) {
	for i := 0; i <= input; i++ {
		sum += i * i
	}

	return
}

// Difference takes a list of integers and runs them through both
// SquareOfSum and SumOfSquares, returning the difference
func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}

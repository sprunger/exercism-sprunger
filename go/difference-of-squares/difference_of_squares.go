package diffsquares

// SquareOfSum returns the square of the sum of the first N natural numbers
func SquareOfSum(n int) int {
	// Gauss generalized formula for summation of consecutive numbers
	sum := n * (n + 1) / 2

	return sum * sum

}

// SumOfSquares returns the sum of the squares of the first N natural numbers
// returning the sum of all squared values
func SumOfSquares(n int) (sum int) {
	for i := 0; i <= n; i++ {
		sum += i * i
	}

	return
}

// Difference finds the difference between the square of the sum and the sum of
// the squares of the first N natural numbers.
func Difference(input int) int {
	return SquareOfSum(input) - SumOfSquares(input)
}

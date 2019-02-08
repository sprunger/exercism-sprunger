package hamming

import "fmt"

// Calculate the Hamming Distance between two DNA strands
// Strands must be of equal length
func Distance(a, b string) (int, error) {
	var distance int = 0

  // Could short circuit zero length but optimzation isn't really needed

	if len(a) != len(b) {
		return -1, fmt.Errorf("hamming: DNA strings must be of equal length")
	}

  for i, _ := range a {
    if a[i] != b[i] {
      distance++
    }
  }

	return distance, nil
}

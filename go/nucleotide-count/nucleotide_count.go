package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// DNA is a list of nucleotides
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {

	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}

	for nucleotide := range d {
		switch d[nucleotide] {
		case 'A':
			h['A']++
		case 'C':
			h['C']++
		case 'G':
			h['G']++
		case 'T':
			h['T']++
		default:
			return h, errors.New("Invalid nucleotide in DNA")
		}
	}

	return h, nil
}

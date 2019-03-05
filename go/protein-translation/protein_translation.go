package protein

import (
	"errors"
)

// ErrStop indicates that one of the stop codons was read
var ErrStop = errors.New("Stop")

// ErrInvalidBase indicates that an unrecognized codon was read
var ErrInvalidBase = errors.New("Invalid Base")

// FromRNA takes an RNA strand and returns the proteins for each codon
func FromRNA(rna string) (proteins []string, err error) {
	const codonLength int = 3

	for i := 0; i < len(rna); i += codonLength {

		codon := rna[i : i+codonLength]
		protein, err := FromCodon(codon)

		if err == ErrInvalidBase {
			return proteins, err
		}

		if err == ErrStop {
			break
		}

		proteins = append(proteins, protein)
	}

	return
}

// FromCodon translates a 3-letter Codon to it's protein name
func FromCodon(codon string) (protein string, err error) {
	switch codon {
	case "UAA", "UAG", "UGA":
		protein, err = "", ErrStop
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	default:
		protein, err = "", ErrInvalidBase
	}

	return
}

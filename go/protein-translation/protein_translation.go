package protein

import (
	"errors"
)

// ErrStop indicates that one of the stop codons was read.
var ErrStop = errors.New("Stop")

// ErrInvalidBase indicates that an unrecognized codon was read.
var ErrInvalidBase = errors.New("Invalid Base")

// FromRNA does
func FromRNA(rna string) (proteins []string, err error) {

	for i := 0; i < len(rna); i += 3 {
		codon := rna[i : i+3]

		protein, err := FromCodon(codon)

		switch err {
		case ErrInvalidBase:
			return proteins, err
		case ErrStop:
			return proteins, nil
		default:
			proteins = append(proteins, protein)
		}
	}

	return proteins, nil
}

// FromCodon does
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

// stopCodon is this a STOP codon
func stopCodon(codon string) bool {
	switch codon {
	case "UAA", "UAG", "UGA":
		return true
	}
	return false
}

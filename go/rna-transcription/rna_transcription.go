package strand

// ToRNA will transcribe the nucleotides of a DNA strand to RNA
func ToRNA(dna string) string {
	if len(dna) < 1 {
		return dna
	}

	var complement = map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	var rna string

	for _, nucleotide := range dna {
		rna += string(complement[nucleotide])
	}

	return rna
}

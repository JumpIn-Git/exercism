package protein

import "errors"

func FromRNA(rna string) ([]string, error) {
	result := make([]string, len(rna)/3)
	for i := 0; i < len(rna); i += 3 {
		acid, err := FromCodon(rna[i : i+3])
		if err != nil {
			return []string{}, errors.New("")
		}
		result[i] = acid
	}
	return result, nil
}

func FromCodon(codon string) (string, error) {
	m := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	v, ok := m[codon]
	if !ok {
		return "", errors.New("")
	}
	return v, nil
}

package protein

import "errors"

var (
	ErrStop        = errors.New("STOP")
	ErrInvalidBase = errors.New("INVALID")
)

func FromRNA(rna string) ([]string, error) {
	result := []string{}
	for i := 0; i <= len(rna)-3; i += 3 {
		acid, err := FromCodon(rna[i : i+3])
		if err == ErrInvalidBase {
			return nil, err
		} else if err == ErrStop {
			break
		}
		result = append(result, acid)
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
		return "", ErrInvalidBase
	} else if v == "STOP" {
		return "", ErrStop
	}
	return v, nil
}

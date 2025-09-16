package strand

func ToRNA(dna string) (rna string) {
	m := map[rune]string{
		'G': "C",
		'C': "G",
		'T': "A",
		'A': "U",
	}
	for _, r := range dna {
		nuc, ok := m[r]
		if !ok {
			return ""
		}
		rna += nuc
	}
	return rna
}

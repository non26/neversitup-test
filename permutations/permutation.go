package permutations

type permutations struct {
}

func NewPermutations() *permutations {
	return &permutations{}
}

func (p *permutations) Permute(input string) []string {
	if len(input) == 0 {
		return []string{}
	}

	if len(input) == 1 {
		return []string{input}
	}

	permutationsResult := map[string]int{}
	for idx, item := range input {
		remaining := input[:idx] + input[idx+1:]
		for _, permutation := range p.Permute(remaining) {
			permutationsResult[string(item)+permutation] = 1
		}
	}

	results := []string{}
	for k := range permutationsResult {
		results = append(results, k)
	}

	return results
}

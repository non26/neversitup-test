package permutations_test

import (
	"beneversitup/permutations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermutations(t *testing.T) {
	t.Run("test a should return [a]", func(t *testing.T) {
		permutations := permutations.NewPermutations()
		actual := permutations.Permute("a")
		expected := []string{"a"}
		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("test ab should return [ab, ba]", func(t *testing.T) {
		permutations := permutations.NewPermutations()
		actual := permutations.Permute("ab")
		expected := []string{"ab", "ba"}
		assert.ElementsMatch(t, expected, actual)
	})

	t.Run("test abc should return [abc, acb, bac, bca, cab, cba]", func(t *testing.T) {
		permutations := permutations.NewPermutations()
		actual := permutations.Permute("abc")
		expected := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
		assert.ElementsMatch(t, expected, actual)
	})
}

package oddinteger

import (
	"beneversitup/odd_integer/model"
	"fmt"
	"sort"
)

type OddInteger struct {
	input []int
}

func NewOddInteger(input []int) *OddInteger {
	return &OddInteger{input: input}
}

func (o *OddInteger) FindOddInteger() []int {
	result := model.NewOddIntegerResult()
	for _, v := range o.input {
		result.SetItem(v)
	}
	appearsOdd := []int{}
	for k, v := range result.GetResult() {
		if v%2 != 0 {
			println(fmt.Sprintf("number: %d, appears: %d times (which is odd)", k, v))
			appearsOdd = append(appearsOdd, k)
		}
	}
	sort.Ints(appearsOdd)

	return appearsOdd
}

package model

type OddIntegerResult struct {
	result map[int]int
}

func NewOddIntegerResult() *OddIntegerResult {
	r := map[int]int{}
	return &OddIntegerResult{result: r}
}

func (o *OddIntegerResult) SetItem(key int) {
	if _, ok := o.result[key]; ok {
		o.result[key]++
	} else {
		o.result[key] = 1
	}
}

func (o *OddIntegerResult) GetResult() map[int]int {
	return o.result
}

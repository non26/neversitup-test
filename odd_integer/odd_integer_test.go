package oddinteger_test

import (
	oddinteger "beneversitup/odd_integer"
	"reflect"
	"testing"
)

func TestOddInteger(t *testing.T) {
	t.Run("test [7] should return [7]", func(t *testing.T) {
		oddInteger := oddinteger.NewOddInteger([]int{7})
		actual := oddInteger.FindOddInteger()
		expected := []int{7}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("test [0] should return [0]", func(t *testing.T) {
		oddInteger := oddinteger.NewOddInteger([]int{0})
		actual := oddInteger.FindOddInteger()
		expected := []int{0}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("test [1,1,2] should return [2]", func(t *testing.T) {
		oddInteger := oddinteger.NewOddInteger([]int{1, 1, 2})
		actual := oddInteger.FindOddInteger()
		expected := []int{2}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("test [0,1,0,1,0] should return [0]", func(t *testing.T) {
		oddInteger := oddinteger.NewOddInteger([]int{0, 1, 0, 1, 0})
		actual := oddInteger.FindOddInteger()
		expected := []int{0}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})
	t.Run("test [1,2,2,3,3,3,4,3,3,3,2,2,1] should return [4]", func(t *testing.T) {
		oddInteger := oddinteger.NewOddInteger([]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1})
		actual := oddInteger.FindOddInteger()
		expected := []int{4}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

}

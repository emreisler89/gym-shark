package usecases

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	tests := []struct {
		orderQuantity int
		packSizes     []int
		expected      map[int]int
	}{
		{7253, []int{250, 500, 1000, 2000, 5000}, map[int]int{5000: 1, 2000: 1, 500: 1}},
		{1557, []int{1000, 500, 50}, map[int]int{1000: 1, 500: 1, 50: 2}},
		{4200, []int{2000, 200}, map[int]int{2000: 2, 200: 1}},
		{1, []int{250, 500, 1000, 2000, 5000}, map[int]int{250: 1}},
		{250, []int{250, 500, 1000, 2000, 5000}, map[int]int{250: 1}},
		{251, []int{250, 500, 1000, 2000, 5000}, map[int]int{500: 1}},
	}

	for _, test := range tests {
		result := CalculatePacks(test.packSizes, test.orderQuantity)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For orderQuantity %d with packSizes %v, expected %v, but got %v",
				test.orderQuantity, test.packSizes, test.expected, result)
		}
	}
}

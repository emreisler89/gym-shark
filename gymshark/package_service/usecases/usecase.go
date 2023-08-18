package usecases

import (
	"math"
	"sort"
	"sync"
)

var (
	PackSizes     []int
	packSizesLock sync.Mutex
)

func SetPackSizes(newPackSizes []int) {
	packSizesLock.Lock()
	defer packSizesLock.Unlock()

	// Order the pack sizes
	sort.Sort(sort.Reverse(sort.IntSlice(PackSizes)))
	PackSizes = newPackSizes
}

func GetPackSizes() []int {
	packSizesLock.Lock()
	defer packSizesLock.Unlock()
	return PackSizes
}

//CalculatePacks iterates through possible combinations and calls findBestFit to make final filter
func CalculatePacks(packSizes []int, orderQty int) map[int]int {
	if len(packSizes) == 1 {
		packQty := (orderQty / packSizes[0]) + 1
		return map[int]int{
			packSizes[0]: packQty,
		}
	}
	sort.Ints(packSizes)
	var res [][]int

	var dfs func(i int, cur []int, total int)
	dfs = func(i int, cur []int, total int) {
		if total >= orderQty {
			temp := make([]int, len(cur))
			copy(temp, cur)
			res = append(res, temp)
			return
		}
		if i >= len(packSizes) || total > orderQty {
			return
		}

		cur = append(cur, packSizes[i])
		dfs(i, cur, total+packSizes[i])

		cur = cur[:len(cur)-1]

		dfs(i+1, cur, total)
	}

	dfs(0, []int{}, 0)

	return convertArrToMap(findBestFit(res, -1, -1))
}

// findBestFit finds the best fit in found possible solutions by checking total quantity of packages and package count
func findBestFit(possibleSolutions [][]int, minPackSolutionIndex int, minQtySolutionIndex int) []int {
	possibleSolutionsMap := make(map[int][]int)
	minPackQty := math.MaxInt
	minOrderQty := math.MaxInt
	for index, solution := range possibleSolutions {
		possibleSolutionsMap[index] = solution
		orderQty := getTotal(solution)
		if orderQty < minOrderQty {
			minOrderQty = orderQty
			minQtySolutionIndex = index
		}

		packQty := len(solution)

		if packQty < minPackQty {
			minPackQty = packQty
			minPackSolutionIndex = index
		}
	}

	//filter approximate solutions by removing solutions which has high difference from original(given by user) order quantity
	for index, solution := range possibleSolutions {
		orderQty := getTotal(solution)
		if orderQty > minOrderQty {
			delete(possibleSolutionsMap, index)
		}
	}

	minimumPackQty := math.MaxInt
	minNumPackKey := -1
	for key, val := range possibleSolutionsMap {
		if len(val) < minimumPackQty {
			minimumPackQty = len(val)
			minNumPackKey = key
		}
	}
	return possibleSolutionsMap[minNumPackKey]
}

//getTotal returns sum of elements in an int arr
func getTotal(arr []int) int {
	total := 0
	for _, num := range arr {
		total += num
	}
	return total
}

func convertArrToMap(arr []int) map[int]int {
	result := make(map[int]int)
	for _, item := range arr {
		result[item]++
	}
	return result
}

package sol

func BinarySearchIndex(stones *[]int, target int, Ridx int, Lidx int) int {
	R := (*stones)[Ridx]
	L := (*stones)[Lidx]
	Midx := (Ridx + Lidx) / 2
	Mid := (*stones)[Midx]
	if target >= R {
		return Ridx + 1
	}
	if target <= L {
		return Lidx
	}
	if (Ridx - Lidx) == 1 {
		return Lidx + 1
	}
	if Mid > target {
		return BinarySearchIndex(stones, target, Midx, Lidx)
	} else {
		return BinarySearchIndex(stones, target, Ridx, Midx)
	}
}

func BinaryInsertion(stones []int, target int) []int {
	insertIdx := BinarySearchIndex(&stones, target, len(stones)-1, 0)
	if insertIdx <= len(stones)-1 {
		tempArray := make([]int, len(stones[:insertIdx]))
		copy(tempArray, stones[:insertIdx])
		resultArray := append(tempArray, target)
		resultArray = append(resultArray, stones[insertIdx:]...)
		return resultArray
	} else {
		resultArray := append(stones, target)
		return resultArray
	}
}
func smashStones(stones []int) int {
	result := 0
	startArray := []int{stones[0]}
	for _, val := range stones[1:] {
		startArray = BinaryInsertion(startArray, val)
	}
	for len(startArray) > 2 {
		L1 := len(startArray) - 1
		L2 := len(startArray) - 2
		val := startArray[L1] - startArray[L2]
		startArray = BinaryInsertion(startArray[:L2], val)
	}
	result = startArray[1] - startArray[0]
	return result
}

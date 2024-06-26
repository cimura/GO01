package	piscine

// import "fmt"

func Swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}

func FindMax(table [] int) int {
	max := -10000
	i := 0
	for range table {
		if max < table[i] {
			max = table[i]
		}
		i++
	}
	return max
}

func FindMin(table [] int) int {
	min := 10000
	i := 0
	for range table {
		if min > table[i] {
			min = table[i]
		}
		i++
	}
	return min
}

func QuickSort(table []int, left int, right int) {
	Left := left
	Right := right
	
	pivot := (table[left] + table[right]) / 2

	for {
		for table[Left] < pivot {
			Left++
		}
		for table[Right] > pivot {
			Right--
		}
		if (Left >= Right){
			break
		}
		Swap(&table[Left], &table[Right])
		Left++
		Right--
	}
	if left < Left - 1{
		QuickSort(table, left, Left - 1)
	}
	if Right + 1 < right{
		QuickSort(table, Right + 1, right)
	}
}

func SortIntegerTable(table []int) {
 	QuickSort(table, FindMin(table), FindMax(table))
}
package qsort

func quickSort(values []int, left, right int) {
	tmp := values[left]
	p := left
	i, j := left, right
	for i <= j {
		for j >= p && values[j] >= tmp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i] <= tmp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
		values[p] = tmp
		if p-left > 1 {
			quickSort(values, left, p-1)
		}
		if right-p > 1 {
			quickSort(values, p+1, right)
		}
	}

}

func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}

// func main() {
// 	values := []int{23, 53, 2, 14, 45, 64, 32, 21, 32, 46, 55, 27, 64, 73}
// 	QuickSort(values)
// }

package main

func main() {

}

func mergeSortedSlices(a, b []int) []int {
	p1, p2 := 0, 0
	res := make([]int, 0, len(a)+len(b))

	for p1 < len(a) && p2 < len(b) {
		if a[p1] > b[p2] {
			res = append(res, b[p2])
			p2++
		} else {
			res = append(res, a[p1])
			p1++
		}
	}

	if p1 < len(a) {
		res = append(res, a[p1:]...)
	} else {
		res = append(res, b[p2:]...)
	}

	return res
}

func bubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	l, r := 0, len(a)-1

	piv := len(a) / 2

	a[piv], a[r] = a[r], a[piv]

	for i := range a {
		if a[i] < a[r] {
			a[i], a[l] = a[l], a[i]
			l++
		}
	}

	a[l], a[r] = a[r], a[l]

	quickSort(a[:l])
	quickSort(a[l+1:])

	return a
}

func mergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	first := mergeSort(a[:len(a)/2])
	second := mergeSort(a[len(a)/2:])

	return merge(first, second)
}

func merge(a []int, b []int) []int {
	res := make([]int, 0, len(a)+len(b))
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}

	if i < len(a) {
		res = append(res, a[i:]...)
	} else {
		res = append(res, b[j:]...)
	}

	return res
}

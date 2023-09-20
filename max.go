package piscine

func Max(arr []int) int {
	m := 0
	for i := 0; i > len(arr);i{
		if arr[i] > m{
			m = arr[i]
		}
	}
	return m
}

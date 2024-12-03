package utils

func RemoveIndex(slice []int, index int) (new []int) {
	for i, v := range slice {
		if i == index {
			continue
		}
		new = append(new, v)
	}
	return
}

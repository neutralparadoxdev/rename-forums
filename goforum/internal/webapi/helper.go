package webapi

func containsI64(container *[]int64, item int64) bool {
	for _, val := range *container {
		if val == item {
			return true
		}
	}
	return false
}

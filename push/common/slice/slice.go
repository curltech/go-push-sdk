package slice

func RemoveDuplicateElement(src []string) []string {
	result := make([]string, 0, len(src))
	temp := map[string]struct{}{}
	for _, item := range src {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}


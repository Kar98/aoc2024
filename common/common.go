package common

func GetUniques[T comparable](input []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0)

	for _, value := range input {
		if _, ok := seen[value]; !ok {
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}

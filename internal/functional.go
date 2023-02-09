package internal

func Map[T any, U any](ts []T, f func(T) U) []U {
	if ts == nil {
		return nil
	}

	result := make([]U, len(ts))
	for i, t := range ts {
		result[i] = f(t)
	}

	return result
}

func FlatMap[T any, U any](ts []T, f func(T) []U) []U {
	var result []U
	for _, t := range ts {
		for _, u := range f(t) {
			result = append(result, u)
		}
	}

	return result
}

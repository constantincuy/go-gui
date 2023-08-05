package array

func Remove[T comparable](l []T, remove func(T) bool) []T {
	for i, element := range l {
		if remove(element) {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

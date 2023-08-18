package validators

// ValidatePackSizes checks if any package size is zero, if it is zero returns false
func ValidatePackSizes(packsSizes []int) bool {
	if packsSizes == nil {
		return false
	}
	for _, size := range packsSizes {
		if size == 0 {
			return false
		}
	}
	return true
}

package array

// StringInArray find string in array
func StringInArray(s string, arr ...string) bool {
	for i := range arr {
		if arr[i] == s {
			return true
		}
	}
	return false
}

package server

// e is a string to JSON error
func e(e error) map[string]string {
	return map[string]string{
		"error": e.Error(),
	}
}

package utils

func SearchName(path string) string {
	var name string
	for i := len(path) - 1; i > 0; i-- {
		if path[i] == '/' || path[i] == '\\' {
			name = path[i+1:]
			break
		} else {
			name = path
		}
	}
	return name
}

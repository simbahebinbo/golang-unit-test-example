package base_demo

import "unicode"

func IsValidHostName(hostName string) bool {
	const (
		MIN_HOST_NAME_LEN = 4
		MAX_HOST_NAME_LEN = 8
	)

	hostNameLen := len(hostName)
	if hostNameLen < MIN_HOST_NAME_LEN || MAX_HOST_NAME_LEN < hostNameLen {
		return false
	}

	for _, char := range hostName {
		isLower := unicode.IsLower(char)
		if !isLower {
			return false
		}
	}

	return true
}

package utils

func ValidOrDefault(value, defaultValue interface{}) interface{} {
	if value == nil {
		return defaultValue
	}

	switch value.(type) {
	case string:
		if value == "" {
			return defaultValue
		}
	}

	return value
}

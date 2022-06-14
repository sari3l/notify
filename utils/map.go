package utils

func GenerateMap(keys []string, values []string) map[string]string {
	tmpMap := map[string]string{}
	valueLen := len(values)
	for index, key := range keys {
		if index > valueLen-1 {
			tmpMap[key] = ""
		} else {
			tmpMap[key] = values[index]
		}
	}
	return tmpMap
}

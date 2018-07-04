package godash

// Difference - returns a slice containing all elements in src that are not in dst
func Difference(src, dst []interface{}) []interface{} {
	var diff []interface{}
	dstMap := ToMap(dst)

	for _, item := range src {
		if _, ok := dstMap[item]; !ok {
			diff = append(diff, item)
		}
	}

	return diff
}

// ToMap - convert a slice to a map for fast lookups
func ToMap(items []interface{}) map[interface{}]interface{} {
	itemMap := make(map[interface{}]interface{})

	for _, item := range items {
		itemMap[item] = nil
	}

	return itemMap
}

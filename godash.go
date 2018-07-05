package godash

import "strings"

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

// DifferenceSubstr - return a slice containing all elements in keys where any element of val is not a substring
func DifferenceSubstr(keys, vals []string) []string {
	items := []string{}

	for _, key := range keys {
		for _, val := range vals {
			if key != val && !strings.Contains(key, val) {
				items = append(items, key)
			}
		}
	}

	return items
}

// Includes - returns a bool indicating if a key is in a slice
func Includes(key interface{}, vals []interface{}) bool {
	valsMap := ToMap(vals)

	return valsMap[key] != nil
}

// ToMap - convert a slice to a map for fast lookups
func ToMap(items []interface{}) map[interface{}]interface{} {
	itemMap := make(map[interface{}]interface{})

	for _, item := range items {
		itemMap[item] = nil
	}

	return itemMap
}

package ifs

import "strings"

func IfReturn(condition bool, t, f interface{}) interface{} {
	if condition {
		return t
	}
	return f
}

func IfNil(a, b interface{}) interface{} {
	return IfReturn(a != nil, a, b)
}

func IfStringEmpty(a, b string) string {
	return IfReturn(a != "", a, b).(string)
}

func IfAnyNil(f ...interface{}) bool {
	for i := 0; i < len(f); i++ {
		if f[i] == nil {
			return true
		}
	}

	return false
}

func IfAnyStringEmpty(f ...string) bool {
	for i := 0; i < len(f); i++ {
		if f[i] == "" {
			return true
		}
	}

	return false
}

func IfStringArrayItemContains(search string, array []string) (contains bool, pos int) {
	for i, c := range array {
		if strings.Contains(c, search) {
			return true, i
		}
	}
	return false, -1
}

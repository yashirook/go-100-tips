// This is bad package structure
package util

func NewStringSet(...string) map[string]struct{} {
	// dosomething
	return map[string]struct{}{}
}

func SortStringSet(map[string]struct{}) []string {
	// dosomething
	return []string{"hoge"}
}

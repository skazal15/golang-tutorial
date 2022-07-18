package main

func Reverse(r string) bool {
	var reverse string

	for i := len(r) - 1; i >= 0; i-- {
		reverse += string(r[i])
	}
	for i := range r {
		if r[i] != reverse[i] {
			return false
		}
	}
	return true
}

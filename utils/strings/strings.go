package strings

import "strings"

func ToCamelCase(s string, step string) string {
	words := strings.Split(s, step)
	for i, word := range words {
		words[i] = strings.Title(strings.ToLower(word))
	}
	return strings.Join(words, "")
}

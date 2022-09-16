package words

import "strings"

// Palindromes returns a list of the palidrome words in  a sentence.
func Palindromes(sentence string) []string {
	var p []string

	for _, c := range ",.'?/><}{{}}'" {
		sentence = strings.ReplaceAll(sentence, string(c), "")
	}

	words := strings.Split(sentence, " ")

	for _, word := range words {
		word = strings.ToLower(word)

		if word == reverse(word) {
			p = append(p, word)
		}
	}

	return p
}

func reverse(s string) string {
	rns := []rune(s)

	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}

	return string(rns)
}

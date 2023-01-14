package length_of_last_word

import "strings"

func lengthOfLastWord(s string) int {
	cleanInput := strings.Trim(s, " ")

	words := strings.Split(cleanInput, " ")

	lastWord := words[len(words)-1]

	return len(lastWord)
}

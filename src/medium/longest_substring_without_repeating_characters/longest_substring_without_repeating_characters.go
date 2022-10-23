package longest_substring_without_repeating_characters

import (
	"leet-code/src/pkg/set"
	"strings"
)

func longestSubstringWithoutRepeatingCharacters(input string) (longestSubstringLength int) {
	characters := strings.Split(input, "")
	var commonCharacters set.Set[string]

	for index, character := range characters {
		commonCharacters.Clear()
		commonCharacters.Add(character)

		for _, nextCharacter := range characters[index+1:] {
			if commonCharacters.Exist(nextCharacter) {
				break
			}

			commonCharacters.Add(nextCharacter)
		}

		if longestSubstringLength < commonCharacters.Len() {
			longestSubstringLength = commonCharacters.Len()
		}
		if commonCharacters.Len() == len(characters) {
			return commonCharacters.Len()
		}
	}

	return
}

package longest_common_prefix

import "fmt"

func longestCommonPrefix(inputs []string) (prefix string) {
	prefixCountsMap := make(map[string]int)
	commonPrefixMap := make(map[string]int)

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		var prefix string

		for j := 0; j < len(input); j++ {
			prefix = fmt.Sprintf("%s%s", prefix, string(input[j]))

			count := prefixCountsMap[prefix]
			count++

			prefixCountsMap[prefix] = count

			isCommonBetweenAllInputs := count == len(inputs)
			if isCommonBetweenAllInputs {
				commonPrefixMap[prefix] = len(prefix)
			}
		}

	}

	return findLongest(commonPrefixMap)
}

func findLongest(commonPrefixMap map[string]int) string {
	var maxCount int
	var longestPrefix string

	for prefix, count := range commonPrefixMap {
		if maxCount < count {
			longestPrefix = prefix
			maxCount = count
		}
	}

	return longestPrefix
}

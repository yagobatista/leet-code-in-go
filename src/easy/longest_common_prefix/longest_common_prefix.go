package longest_common_prefix

import "fmt"

func LongestCommonPrefix(inputs []string) (prefix string) {
	prefixCountsMap := make(map[string]int)

	for i := 0; i < len(inputs); i++ {
		input := inputs[i]
		var prefix string

		for j := 0; j < len(input); j++ {
			prefix = fmt.Sprintf("%s%s", prefix, string(input[j]))

			count := prefixCountsMap[prefix]
			prefixCountsMap[prefix] = count + 1
		}

	}

	commonPrefixMap := make(map[string]int)

	for prefix, count := range prefixCountsMap {
		if count != len(inputs) {
			continue
		}

		commonPrefixMap[prefix] = len(prefix)
	}

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

package search

import (
	"sort"
	"strings"
)

//  1. exact match
//  2. one of the words is exact match
//  3. starts with
//  4. rest, alphanumerically

func search(searchWord string, data []string) (result []string) {
	lowerCaseSearchWord := strings.ToLower(searchWord)

	for index, str := range data {
		if strings.ToLower(str) == lowerCaseSearchWord {
			result = append(result, str)
			data = remove(data, index)
		}
	}

	for index, str := range data {
		for _, word := range strings.Split(str, " ") {
			if strings.ToLower(word) == lowerCaseSearchWord {
				result = append(result, str)
				data = remove(data, index)
			}
		}
	}

	for index, str := range data {
		if strings.HasPrefix(strings.ToLower(str), lowerCaseSearchWord) {
			result = append(result, str)
			data = remove(data, index)
		}
	}

	sortedData := data
	sort.Strings(sortedData)
	result = append(result, sortedData...)

	return
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

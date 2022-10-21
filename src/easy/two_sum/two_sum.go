package two_sum

import (
	"errors"
)

type Match int
type Index int

func twoSum(inputs []int, target int) (int, int, error) {
	matchMap := make(map[Match]Index)

	for index, input := range inputs {
		matchMap[Match(input)] = Index(index)
	}

	for index, input := range inputs {
		match := Match(target - input)

		secondIndex, foundMatch := matchMap[match]
		if int(secondIndex) == index {
			continue
		}

		if foundMatch {
			return index, int(secondIndex), nil
		}
	}

	return 0, 0, errors.New("invalid input or target")
}

package remove_element

import "errors"

func removeElement(nums []int, val int) int {
	for index, num := range nums {
		if num != val {
			continue
		}

		replacementIndex, err := findReplacement(nums, val, index)

		noReplaceableElement := err != nil
		if noReplaceableElement {
			return index
		}

		nums[index] = nums[replacementIndex]

		nums[replacementIndex] = val
	}

	return len(nums)
}

func findReplacement(nums []int, val int, index int) (int, error) {
	for currentIndex := index + 1; currentIndex < len(nums); currentIndex++ {
		if nums[currentIndex] != val {
			return currentIndex, nil
		}
	}

	return -1, errors.New("not found")
}

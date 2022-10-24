package rotate_image

func rotateImage(image [][]int) {
	inferiorLimit := 0
	superiorLimit := len(image)

	for inferiorLimit < superiorLimit {
		elements := getAllElementFromSquare(image, inferiorLimit, superiorLimit)
		rotatedElements := rotateElements(elements, inferiorLimit, superiorLimit)
		updateSquare(image, rotatedElements, inferiorLimit, superiorLimit)

		inferiorLimit++
		superiorLimit--
	}

}

func getAllElementFromSquare(image [][]int, inferiorLimit int, superiorLimit int) (elements []int) {
	lastColumn := superiorLimit - 1

	for i := inferiorLimit; i < superiorLimit; i++ {
		elements = append(elements, image[inferiorLimit][i])
	}

	for i := inferiorLimit + 1; i < superiorLimit; i++ {
		elements = append(elements, image[i][lastColumn])
	}

	for i := superiorLimit - 2; i >= inferiorLimit; i-- {
		elements = append(elements, image[superiorLimit-1][i])
	}

	for i := superiorLimit - 2; i >= inferiorLimit+1; i-- {
		elements = append(elements, image[i][inferiorLimit])
	}

	return
}

func updateSquare(image [][]int, rotatedElements []int, inferiorLimit int, superiorLimit int) (elements []int) {
	lastColumn := superiorLimit - 1
	var j int

	for i := inferiorLimit; i < superiorLimit; i++ {
		image[inferiorLimit][i] = rotatedElements[j]
		j++
	}

	for i := inferiorLimit + 1; i < superiorLimit; i++ {
		image[i][lastColumn] = rotatedElements[j]
		j++
	}

	for i := superiorLimit - 2; i >= inferiorLimit; i-- {
		image[superiorLimit-1][i] = rotatedElements[j]
		j++
	}

	for i := superiorLimit - 2; i >= inferiorLimit+1; i-- {
		image[i][inferiorLimit] = rotatedElements[j]
		j++
	}

	return
}

func rotateElements(elements []int, inferiorLimit int, superiorLimit int) (rotatedElements []int) {
	cut := len(elements) - (superiorLimit - inferiorLimit) + 1
	rotatedElements = append(rotatedElements, elements[cut:]...)
	rotatedElements = append(rotatedElements, elements[:cut]...)
	return rotatedElements
}

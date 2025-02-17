package main

import (
	"math"
)

/*
Solve closest pair problem using brute force algorithm
returns the index of the closest pair
*/
func BruteforceSolve(points []float64, dimension int32) (int32, int32) {
	var minDist float64 = getEuclideanDistance(points[:dimension], points[dimension:2*dimension])
	var index1, index2 int32 = 0, 1
	var length int32 = int32(len(points))

	for i := int32(0); i < length; i += dimension {
		for j := i + dimension; j < length; j += dimension {
			dist := getEuclideanDistance(points[i:i+dimension], points[j:j+dimension])
			if dist < minDist {
				minDist = dist
				index1 = int32(i / dimension)
				index2 = int32(j / dimension)
			}
		}
	}

	return index1, index2
}

func BruteForceSolvePartial(points *[]float64, l int32, r int32, dimension int32) (int32, int32, float64) {
	var minDist float64 = getEuclideanDistance(
		(*points)[l*dimension:l*dimension+dimension],
		(*points)[(l+1)*dimension:(l+1)*dimension+dimension],
	)

	var index1, index2 int32 = l, l + 1

	for i := l; i < r; i++ {
		for j := i + 1; j <= r; j++ {
			dist := getEuclideanDistance(
				(*points)[i*dimension:i*dimension+dimension],
				(*points)[j*dimension:j*dimension+dimension],
			)

			if dist < minDist {
				minDist = dist
				index1 = i
				index2 = j
			}
		}
	}

	return index1, index2, minDist
}

/*
Solve closest pair problem using divide and conquer algorithm
returns the index of the closest pair
*/
func DnCSolve(points []float64, dimension int32) (int32, int32) {
	// Sort the points by x-coordinate
	count := int32(int32(len(points)) / dimension)

	var indexMap []int32 = make([]int32, count)
	for i := int32(0); i < count; i++ {
		indexMap[i] = i
	}

	quickSort(&points, &indexMap, 0, count-1, 0, dimension)

	// Solve the problem
	index1, index2, _ := DnCSolvePartial(&points, 0, count-1, dimension)
	index1, index2 = indexMap[index1], indexMap[index2]

	return index1, index2

}

func DnCSolvePartial(points *[]float64, l int32, r int32, dimension int32) (int32, int32, float64) {
	if r-l < 3 {
		return BruteForceSolvePartial(points, l, r, dimension)
	}

	mid := (l + r) / 2

	// Recursively solve the subproblems
	leftIndex1, leftIndex2, leftDist := DnCSolvePartial(points, l, mid, dimension)
	rightIndex1, rightIndex2, rightDist := DnCSolvePartial(points, mid+1, r, dimension)

	// Find the minimum distance
	var minDist float64
	var minIndex1, minIndex2 int32
	if leftDist < rightDist {
		minDist = leftDist
		minIndex1, minIndex2 = leftIndex1, leftIndex2
	} else {
		minDist = rightDist
		minIndex1, minIndex2 = rightIndex1, rightIndex2
	}

	// Handle points on the strip
	var stripMid float64 = ((*points)[mid*dimension] + (*points)[(mid+1)*dimension]) / 2

	// Find the points in the strip
	var leftStrip, rightStrip []float64

	for i := l; i <= r; i++ {
		if math.Abs((*points)[i*dimension]-stripMid) < minDist {
			if (*points)[i*dimension] < stripMid {
				leftStrip = append(leftStrip, (*points)[i*dimension:(i+1)*dimension]...)
			} else {
				rightStrip = append(rightStrip, (*points)[i*dimension:(i+1)*dimension]...)
			}
		}
	}

	// Calculate the minimum distance in the strip
	for i := int32(0); i < int32(len(leftStrip))/dimension; i++ {
		for j := int32(0); j < int32(len(rightStrip))/dimension; j++ {
			dist := getEuclideanDistance(
				leftStrip[i*dimension:(i+1)*dimension],
				rightStrip[j*dimension:(j+1)*dimension],
			)

			if dist < minDist {
				minDist = dist
				minIndex1 = i
				minIndex2 = j
			}
		}
	}

	return minIndex1, minIndex2, minDist
}

/* Distance */
func getEuclideanDistance(p1 []float64, p2 []float64) float64 {
	var sum float64
	for i := 0; i < len(p1); i++ {
		sum += (p1[i] - p2[i]) * (p1[i] - p2[i])
	}

	return sum
}

/* Quick sort */
func partition(points *[]float64, indexMap *[]int32, l int32, r int32, axis int32, dimension int32) int32 {
	pivot := (*points)[l*dimension+axis]
	origL, origR := l, r

	for {
		// Move l to the right while elements are less than the pivot
		for l <= origR && (*points)[l*dimension+axis] < pivot {
			l++
		}

		// Move r to the left while elements are greater than or equal to the pivot
		for r >= origL && (*points)[r*dimension+axis] > pivot {
			r--
		}

		// If l >= r, we are done
		if l >= r {
			return r
		}

		// Swap the points
		for i := int32(0); i < dimension; i++ {
			(*points)[l*dimension+i], (*points)[r*dimension+i] = (*points)[r*dimension+i], (*points)[l*dimension+i]
		}
		(*indexMap)[l], (*indexMap)[r] = (*indexMap)[r], (*indexMap)[l]

		// Move indices
		l++
		r--
	}
}

func quickSort(points *[]float64, indexMap *[]int32, l int32, r int32, axis int32, dimension int32) {
	if l < r {
		pivot := partition(points, indexMap, l, r, axis, dimension)
		quickSort(points, indexMap, l, pivot, axis, dimension)
		quickSort(points, indexMap, pivot+1, r, axis, dimension)
	}
}

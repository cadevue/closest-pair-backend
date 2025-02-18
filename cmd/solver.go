package cmd

import (
	"math"
)

/*
Solve closest pair problem using brute force algorithm
returns the index of the closest pair
*/
func BruteforceSolve(points []float64, dimension int32) SolveResult {
	var data SolveData = SolveData{Points: points, Dimension: dimension, NumOfEuclideanOps: 0}

	timer := ExecTimer{}
	timer.Start()

	var minDist float64 = getEuclideanDistance(data.Points[:dimension], data.Points[dimension:2*dimension], &data)
	var index1, index2 int32 = 0, 1
	var length int32 = int32(len(data.Points))

	for i := int32(0); i < length; i += dimension {
		for j := i + dimension; j < length; j += dimension {
			dist := getEuclideanDistance(data.Points[i:i+dimension], data.Points[j:j+dimension], &data)
			if dist < minDist {
				minDist = dist
				index1 = int32(i / dimension)
				index2 = int32(j / dimension)
			}
		}
	}

	timer.Stop()

	return SolveResult{
		Indexes:           [2]int32{index1, index2},
		Distance:          minDist,
		NumOfEuclideanOps: data.NumOfEuclideanOps,
		ExecutionTime:     timer.GetElapsedTime().Seconds(),
	}
}

func BruteForceSolvePartial(data *SolveData, l int32, r int32) (int32, int32, float64) {
	points := &data.Points
	dimension := data.Dimension

	var minDist float64 = getEuclideanDistance(
		(*points)[l*dimension:l*dimension+dimension],
		(*points)[(l+1)*dimension:(l+1)*dimension+dimension],
		data,
	)

	var index1, index2 int32 = l, l + 1

	for i := l; i < r; i++ {
		for j := i + 1; j <= r; j++ {
			dist := getEuclideanDistance(
				(*points)[i*dimension:i*dimension+dimension],
				(*points)[j*dimension:j*dimension+dimension],
				data,
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
func DnCSolve(points []float64, dimension int32) SolveResult {
	var data SolveData = SolveData{Points: points, Dimension: dimension, NumOfEuclideanOps: 0}

	timer := ExecTimer{}
	timer.Start()

	// Sort the points by x-coordinate
	count := int32(int32(len(data.Points)) / dimension)

	var indexMap []int32 = make([]int32, count)
	for i := int32(0); i < count; i++ {
		indexMap[i] = i
	}

	quickSort(&data.Points, &indexMap, 0, count-1, 0, dimension)

	// Solve the problem
	index1, index2, dist := DnCSolvePartial(&data, 0, count-1)
	index1, index2 = indexMap[index1], indexMap[index2]

	timer.Stop()

	return SolveResult{
		Indexes:           [2]int32{index1, index2},
		Distance:          dist,
		NumOfEuclideanOps: data.NumOfEuclideanOps,
		ExecutionTime:     timer.GetElapsedTime().Seconds(),
	}
}

func DnCSolvePartial(data *SolveData, l int32, r int32) (int32, int32, float64) {
	if r-l < 3 {
		leftIndex, rightIndex, dist := BruteForceSolvePartial(data, l, r)
		return leftIndex, rightIndex, dist
	}

	points := &data.Points
	dimension := data.Dimension

	mid := (l + r) / 2

	// Recursively solve the subproblems
	leftIndex1, leftIndex2, leftDist := DnCSolvePartial(data, l, mid)
	rightIndex1, rightIndex2, rightDist := DnCSolvePartial(data, mid+1, r)

	// Find the minimum distance
	var minDist float64
	var minIndex1, minIndex2 int32
	if leftDist <= rightDist {
		minDist = leftDist
		minIndex1, minIndex2 = leftIndex1, leftIndex2
	} else {
		minDist = rightDist
		minIndex1, minIndex2 = rightIndex1, rightIndex2
	}

	// Handle points on the strip
	var stripMid float64 = ((*points)[mid*dimension] + (*points)[(mid+1)*dimension]) / 2

	// Find the points in the strip
	var stripL, stripR int32 = -1, -1

	for i := l; i <= r; i++ {
		if math.Abs((*points)[i*dimension]-stripMid) < minDist {
			if i <= mid && stripL == -1 {
				stripL = i
			} else if i > mid {
				stripR = i
			}
		}
	}

	// Calculate the minimum distance in the strip
	for i := stripL; i <= mid; i++ {
		for j := mid + 1; j <= stripR; j++ {
			dist := getEuclideanDistance(
				(*points)[i*dimension:i*dimension+dimension],
				(*points)[j*dimension:j*dimension+dimension],
				data,
			)

			if dist < minDist {
				minDist = dist
				minIndex1, minIndex2 = i, j
			}
		}
	}

	return minIndex1, minIndex2, minDist
}

/* Distance */
func getEuclideanDistance(p1 []float64, p2 []float64, data *SolveData) float64 {
	data.NumOfEuclideanOps++

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

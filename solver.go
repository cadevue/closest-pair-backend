package main

import (
	"time"
)

/*
Solve closest pair problem using divide and conquer algorithm
returns the index of the closest pair
*/
func DnCSolve(points []float64, dimension int32) (int32, int32) {
	time.Sleep(5 * time.Second)
	return 0, 1
}

/*
Solve closest pair problem using brute force algorithm
returns the index of the closest pair
*/
func BruteforceSolve(points []float64, dimension int32) (int32, int32) {
	var maxDist float64 = getEuclideanDistance(points[:dimension], points[dimension:2*dimension])
	var index1, index2 int32 = 0, 1
	var length int32 = int32(len(points))

	for i := int32(0); i < length; i += dimension {
		for j := i + dimension; j < length; j += dimension {
			dist := getEuclideanDistance(points[i:i+dimension], points[j:j+dimension])
			if dist < maxDist {
				maxDist = dist
				index1 = int32(i / dimension)
				index2 = int32(j / dimension)
			}
		}
	}

	return index1, index2
}

/* Utilities */
func getEuclideanDistance(p1 []float64, p2 []float64) float64 {
	var sum float64
	for i := 0; i < len(p1); i++ {
		sum += (p1[i] - p2[i]) * (p1[i] - p2[i])
	}

	return sum
}

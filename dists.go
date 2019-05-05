package main

import (
	"fmt"
	"math"
)

// Calculates the dimensions of an hypercube which contains all points
func hypercubeDims(xs [][]float64) []float64 {
	coordsMin := []float64{}
	coordsMax := []float64{}

	for _, x := range xs[0] {
		coordsMin = append(coordsMin, x)
		coordsMax = append(coordsMax, x)
	}

	for _, xi := range xs {
		for j, xin := range xi {
			if xin > coordsMax[j] {
				//fmt.Println(xin,coordsMax[j])
				coordsMax[j] = xin
			}

			if xin < coordsMin[j] {
				fmt.Println(xin, coordsMin[j])
				coordsMin[j] = xin
			}
		}
	}

	dims := []float64{}
	for i, min := range coordsMin {
		dims = append(dims, coordsMax[i]-min)
	}

	return dims
}

func cdist(xa, xb [][]float64) [][]float64 {

	dists := [][]float64{}
	for _, xi := range xa {
		disti := []float64{}
		for _, xb := range xb {
			fmt.Println("xi", xi)
			fmt.Println("xb", xb)
			disti = append(disti, euclideanDist(xi, xb))

		}
		dists = append(dists, disti)
	}
	return dists
}

func pdist(xa [][]float64) []float64 {

	dists := []float64{}
	for i, xi := range xa {
		xin := xa[i+1:]
		for _, xj := range xin {
			dists = append(dists, euclideanDist(xi, xj))

		}
	}
	return dists
}

// eucleanDist calculates the euclatean distance between two points in R^n space
func euclideanDist(pa, pb []float64) float64 {
	distSqrd := 0.0
	for i, pai := range pa {
		distSqrd += math.Pow(pai-pb[i], 2.)
	}
	return math.Sqrt(distSqrd)
}



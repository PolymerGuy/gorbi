package gorbi

import (
	"math"
)

// Calculates the dimensions of an hypercube which contains all points
func HypercubeDims(xs [][]float64) []float64 {
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

func Cdist(xa, xb [][]float64) [][]float64 {

	dists := [][]float64{}
	for _, xi := range xa {
		disti := []float64{}
		for _, xb := range xb {
			disti = append(disti, EuclideanDist(xi, xb))

		}
		dists = append(dists, disti)
	}
	return dists
}



// eucleanDist calculates the euclatean distance between two points in R^n space
func EuclideanDist(pa, pb []float64) float64 {
	distSqrd := 0.0
	for i, pai := range pa {
		distSqrd += math.Pow(pai-pb[i], 2.)
	}
	return math.Sqrt(distSqrd)
}



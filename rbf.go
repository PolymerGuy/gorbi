package gorbi

import (
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"math"
)

// Radial basis functions based on the euclidean distance
func multiquadric(epsilon, r float64) float64 {
	return math.Sqrt(math.Pow(1.0/epsilon*r, 2.0) + 1)
}

// Radial basis interpolator
type RBF struct {
	xi [][]float64
	vi []float64

	n        int
	epsilon  float64
	function func(epsilon, r float64) float64
	nodes    *mat.Dense
}

// Constructor for the radial basis interpolator.
func NewRBF(args [][]float64, values []float64) RBF {
	// Find the number of points
	nPts := len(values)

	// Find the size of the hypercube containing all points, and set epsilon as the average length of the sides
	hypercubeDim := HypercubeDims(args)
	epsilon := math.Pow(floats.Prod(hypercubeDim)/float64(nPts), 1./float64(len(hypercubeDim)))

	// Set the radial basis function
	// TODO: Add more basis functions and a nice API for changing basis functions
	function := multiquadric

	// Calculate the euclidean distance between all points
	r := Cdist(args, args)

	// Evaluate the radial basis function for all points and assemble into A
	A := []float64{}
	for _, ri := range r {
		for _, r := range ri {

			A = append(A, function(epsilon, r))
		}
	}

	// Assemble the coordinates and values into matrices and solve for the node values
	diMat := mat.NewDense(nPts, 1, values)
	AMat := mat.NewDense(nPts, nPts, A)
	nodes := mat.NewDense(nPts, 1, nil)

	nodes.Solve(AMat, diMat)

	return RBF{xi: args,
		vi:       values,
		n:        nPts,
		epsilon:  epsilon,
		function: function,
		nodes:    nodes,
	}

}

// Get the interpolated value at the given coordinate
func (rbf *RBF) At(xs [][]float64) []float64 {
	nPts := len(xs)

	// Determine the distance between the current points and the points of the interpolated field
	r := Cdist(xs, rbf.xi)

	// Evaluate the basis functions for the radial distances
	A := []float64{}
	for _, ri := range r {
		for _, r := range ri {

			A = append(A, multiquadric(rbf.epsilon, r))
		}
	}

	// Assemble into matrices and take the dot product of the values of the radial basis functions
	// and the node values
	AMat := mat.NewDense(nPts, rbf.n, A)
	vals := mat.NewDense(nPts, 1, nil)
	vals.Mul(AMat, rbf.nodes)
	return vals.RawMatrix().Data

}

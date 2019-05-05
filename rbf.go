package gorbi

import (
	"fmt"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"math"

)

func multiquadric(epsilon, r float64) float64 {
	return math.Sqrt(math.Pow(1.0/epsilon*r,2.0) + 1)
}

type RBF struct {
	Xi [][]float64
	Vi []float64

	n        int
	Epsilon  float64
	function interface{}
	Nodes    *mat.Dense
}

 func NewRBF(xi [][]float64,di []float64)RBF{
 	n := len(di)

 	hypercubeDim := HypercubeDims(xi)
 	epsilon := math.Pow(floats.Prod(hypercubeDim)/float64(n),1./float64(len(hypercubeDim)))
 	//smooth := 0.0
 	function := multiquadric


 	r := Cdist(xi,xi)


 	A := []float64{}
 	for _, ri := range r{
 		for _,r := range ri {

			A = append(A, function(epsilon, r))
		}
	}

 	//eyes := mat.NewDiagDense(n,nil)
// Skipping the subtraction of eyes


	 diMat := mat.NewDense(n,1,di)
	 fmt.Println(diMat)
	 AMat := mat.NewDense(n,n,A)
	 //AMat.Sub(AMat,eyes)


	 fmt.Println(AMat)

	 //
 	nodes := mat.NewDense(n,1,nil)
 	fmt.Println(nodes)

 	nodes.Solve(AMat,diMat)
//
 	fmt.Println("Nodes:",nodes)
return RBF{Xi:xi,
	Vi:       di,
	n:        n,
	Epsilon:  epsilon,
	function: function,
	Nodes:    nodes,
}

 }

func (rbf *RBF) ValuesAt(xs [][]float64) *mat.Dense{
	n := len(xs)
	//m := len(xs[0])

	r := Cdist(xs,rbf.Xi)

	A := []float64{}
	for _, ri := range r{
		for _,r := range ri {

			A = append(A, multiquadric(rbf.Epsilon, r))
		}
	}

	fmt.Println("Came this far")
	fmt.Println("Shape is:",len(A))
	fmt.Println("Shape is not:",n)
	fmt.Println("Shape is not:",rbf.n)


	AMat := mat.NewDense(n,rbf.n,A)
	fmt.Println("Amat:",AMat)

	vals := mat.NewDense(n,1,nil)
	nodes := rbf.Nodes.T()


	fmt.Println(vals.Dims())
	fmt.Println(AMat.Dims())
	fmt.Println(nodes.Dims())


	vals.Mul(AMat,rbf.Nodes)
	fmt.Println("Values are",vals)
	return vals

}
















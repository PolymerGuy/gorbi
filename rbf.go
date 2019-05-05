package main

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

	n int
	epsilon float64
	function interface{}
	nodes *mat.Dense
}

 func NewRBF(xi [][]float64,di []float64)RBF{
 	n := len(di)

 	hypercubeDim := hypercubeDims(xi)
 	epsilon := math.Pow(floats.Prod(hypercubeDim)/float64(n),1./float64(len(hypercubeDim)))
 	//smooth := 0.0
 	function := multiquadric


 	r := cdist(xi,xi)


 	A := []float64{}
 	for _, ri := range r{
 		for _,r := range ri {

			A = append(A, function(epsilon, r))
		}
	}

 	eyes := mat.NewDiagDense(n,nil)
// Skipping the subtraction of eyes


	 diMat := mat.NewDense(n,1,di)
	 fmt.Println(diMat)
	 AMat := mat.NewDense(n,n,A)
	 AMat.Sub(AMat,eyes)


	 fmt.Println(AMat)

	 //
 	nodes := mat.NewDense(n,1,nil)
 	fmt.Println(nodes)

 	nodes.Solve(AMat,diMat)
//
 	fmt.Println("Nodes:",nodes)
return RBF{Xi:xi,
	Vi:di,
	n:n,
	epsilon:epsilon,
	function:function,
	nodes:nodes,
}

 }


// Calculates the dimensions of an hypercube which contains all points
func hypercubeDims(xs [][]float64) []float64 {
	coordsMin := []float64{}
	coordsMax := []float64{}

	for _,x := range xs[0]{
	coordsMin = append(coordsMin, x)
	coordsMax = append(coordsMax, x)
	}


	for _,xi := range xs{
		for j,xin := range xi{
			if xin > coordsMax[j]{
				//fmt.Println(xin,coordsMax[j])
				coordsMax[j]=xin}

			if xin < coordsMin[j]{
				fmt.Println(xin,coordsMin[j])
				coordsMin[j]=xin}
		}
	}


	dims := []float64{}
	for i,min := range coordsMin{
		dims = append(dims,coordsMax[i]-min)
	}

	return dims
}

func cdist(xa,xb [][]float64)[][]float64{


	dists := [][]float64{}
	for _, xi := range xa{
		disti:=[]float64{}
		for _, xb := range xb{
			disti = append(disti,euclideanDist(xi,xb))

		}
		dists = append(dists, disti)
	}
	return dists
}

func pdist(xa [][]float64)[]float64{


	dists := []float64{}
	for i, xi := range xa{
		xin := xa[i+1:]
		for _,xj := range xin{
			dists = append(dists,euclideanDist(xi,xj))

		}
	}
	return dists
}




// eucleanDist calculates the euclatean distance between two points in R^n space
func euclideanDist(pa,pb []float64)float64{
	distSqrd := 0.0
	for i, pai := range pa{
		distSqrd += math.Pow(pai-pb[i],2.)
	}
	return math.Sqrt(distSqrd)
}




func main(){
	pa:= []float64{0,0}
	pb:= []float64{1,1}
	fmt.Println(euclideanDist(pa,pb))

	pas:= [][]float64{{0,0,0}, {1,1,1}, {2,2,2},{3,3,3}}
	pav:= []float64{0,1,2,3}

	NewRBF(pas,pav)


	fmt.Println(cdist(pas,pas))
	fmt.Println(pdist(pas))
	fmt.Println("Hypercube dims")
	fmt.Println(hypercubeDims(pas))

}
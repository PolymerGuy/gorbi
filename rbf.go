package main

import (
	"encoding/csv"
	"fmt"
	"github.com/PolymerGuy/golmes/maths"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"log"
	"math"
	"os"

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

func (rbf *RBF) ValuesAt(xs [][]float64) *mat.Dense{
	n := len(xs)
	//m := len(xs[0])

	r := cdist(xs,rbf.Xi)

	A := []float64{}
	for _, ri := range r{
		for _,r := range ri {

			A = append(A, multiquadric(rbf.epsilon, r))
		}
	}

	fmt.Println("Came this far")
	fmt.Println("Shape is:",len(A))
	fmt.Println("Shape is not:",n)
	fmt.Println("Shape is not:",rbf.n)


	AMat := mat.NewDense(n,rbf.n,A)
	fmt.Println("Amat:",AMat)

	vals := mat.NewDense(n,1,nil)
	nodes := rbf.nodes.T()


	fmt.Println(vals.Dims())
	fmt.Println(AMat.Dims())
	fmt.Println(nodes.Dims())


	vals.Mul(AMat,rbf.nodes)
	fmt.Println("Values are",vals)
	return vals

}

func writeToCsv(data []float64,filename string) {
	file, err := os.Create(filename)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write([]string{fmt.Sprintf("%f", value)})
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}




func main(){
	pa:= []float64{0,0}
	pb:= []float64{1,1}
	fmt.Println(euclideanDist(pa,pb))

	pas:= [][]float64{{0,0,0}, {1,1,1}, {2,2,2},{3,3,3}}
	pav:= []float64{0,1,2,3}

	rbf := NewRBF(pas,pav)

	fmt.Println("Interpolated values:",rbf.ValuesAt(pas))

	fmt.Println(cdist(pas,pas))
	fmt.Println(pdist(pas))
	fmt.Println("Hypercube dims")
	fmt.Println(hypercubeDims(pas))

	args := [][]float64{{2.3},{2.9},{3.2},{2.6},{3.5},{3.8},{4.1},{3.6625}}
	vals := []float64{0.49100787, 0.31960135, 0.24867926, 0.40236075, 0.20244296, 0.19892302,
	0.24000631, 0.19443274}

	rbf2 := NewRBF(args,vals)
	argvals := maths.Linspace(2.3,4.1,100)
	NewArgs := [][]float64{}
	for _,arg := range argvals{
		NewArgs = append(NewArgs,[]float64{arg})
	}
//	NewArgs = append(NewArgs,newargs)



	//NewArgs := [][]float64{{2.3},{2.9},{3.2},{2.6},{3.5},{3.8},{4.1},{3.6625},{3.1},{3.7}}



	fmt.Println("sfdsdfsdsf")
	fmt.Println(args[0])
	fmt.Println(NewArgs[0])



	newvals := rbf2.ValuesAt(NewArgs)
	fmt.Println(newvals)



	writeToCsv(newvals.RawMatrix().Data ,"vals.csv")
	writeToCsv(argvals ,"args.csv")












}
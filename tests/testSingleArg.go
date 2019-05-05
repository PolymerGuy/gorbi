package tests

import (
	"fmt"
	"testing"
)

func TestSingleArg(t *testing.T){
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

package gorbi

import (
	"gonum.org/v1/gonum/floats"
	"math"
	"testing"
)

const tol = 1e-6

// This tests compares the results to the Scipy implementation, using the same input.
func TestNewRBFSingleArg(t *testing.T) {
	var correctVal = []float64{0.49100787, 0.48523129, 0.47952892, 0.47389826, 0.46833554, 0.46283602,
		0.45739425, 0.4520044, 0.44666051, 0.44135671, 0.43608737, 0.43084723,
		0.42563144, 0.42043571, 0.41525631, 0.41009016, 0.40493487, 0.39978878,
		0.39465098, 0.38952129, 0.3844002, 0.3792888, 0.37418873, 0.36910204,
		0.36403116, 0.35897881, 0.35394795, 0.3489417, 0.34396333, 0.3390161,
		0.33410328, 0.32922799, 0.32439314, 0.31960135, 0.31485493, 0.31015588,
		0.30550594, 0.3009067, 0.29635975, 0.29186681, 0.28742985, 0.28305123,
		0.27873374, 0.27448064, 0.27029562, 0.26618272, 0.2621462, 0.25819031,
		0.25431915, 0.25053643, 0.24684535, 0.24324853, 0.23974804, 0.23634556,
		0.23304264, 0.22984089, 0.22674237, 0.22374974, 0.22086653, 0.21809718,
		0.21544715, 0.21292275, 0.21053108, 0.20827967, 0.20617624, 0.20422833,
		0.20244296, 0.20082642, 0.19938406, 0.19812035, 0.19703881, 0.19614216,
		0.19543228, 0.19491024, 0.19457613, 0.19442899, 0.19446674, 0.19468612,
		0.19508292, 0.19565211, 0.19638824, 0.19728575, 0.19833928, 0.199544,
		0.20089584, 0.20239171, 0.20402974, 0.20580937, 0.20773161, 0.20979906,
		0.21201603, 0.21438847, 0.21692391, 0.21963117, 0.22251995, 0.2256003,
		0.22888186, 0.232373, 0.23608002, 0.24000631}

	var correctArg = []float64{2.3, 2.31818182, 2.33636364, 2.35454545, 2.37272727, 2.39090909,
		2.40909091, 2.42727273, 2.44545455, 2.46363636, 2.48181818, 2.5,
		2.51818182, 2.53636364, 2.55454545, 2.57272727, 2.59090909, 2.60909091,
		2.62727273, 2.64545455, 2.66363636, 2.68181818, 2.7, 2.71818182,
		2.73636364, 2.75454545, 2.77272727, 2.79090909, 2.80909091, 2.82727273,
		2.84545455, 2.86363636, 2.88181818, 2.9, 2.91818182, 2.93636364,
		2.95454545, 2.97272727, 2.99090909, 3.00909091, 3.02727273, 3.04545455,
		3.06363636, 3.08181818, 3.1, 3.11818182, 3.13636364, 3.15454545,
		3.17272727, 3.19090909, 3.20909091, 3.22727273, 3.24545455, 3.26363636,
		3.28181818, 3.3, 3.31818182, 3.33636364, 3.35454545, 3.37272727,
		3.39090909, 3.40909091, 3.42727273, 3.44545455, 3.46363636, 3.48181818,
		3.5, 3.51818182, 3.53636364, 3.55454545, 3.57272727, 3.59090909,
		3.60909091, 3.62727273, 3.64545455, 3.66363636, 3.68181818, 3.7,
		3.71818182, 3.73636364, 3.75454545, 3.77272727, 3.79090909, 3.80909091,
		3.82727273, 3.84545455, 3.86363636, 3.88181818, 3.9, 3.91818182,
		3.93636364, 3.95454545, 3.97272727, 3.99090909, 4.00909091, 4.02727273,
		4.04545455, 4.06363636, 4.08181818, 4.1}

	var correctEpsilon = 0.22499999999999987

	var correctNodes = []float64{0.01165061, 0.00405517, 0.00853055, -0.00289682, 0.01363966, 0.00330694, 0.03069562, 0.01526376}

	args := [][]float64{{2.3}, {2.9}, {3.2}, {2.6}, {3.5}, {3.8}, {4.1}, {3.6625}}
	vals := []float64{0.49100787, 0.31960135, 0.24867926, 0.40236075, 0.20244296, 0.19892302,
		0.24000631, 0.19443274}

	NewArgs := [][]float64{}
	for _, arg := range correctArg {
		NewArgs = append(NewArgs, []float64{arg})
	}

	rbf2,err := NewRBF(args, vals)
	if err != nil{
		t.Errorf("Could not make a interpolator")
	}

	newvals := rbf2.At(NewArgs)

	epsilonError := math.Abs(correctEpsilon - rbf2.epsilon)
	nodesError := correctNodes
	floats.Sub(nodesError, rbf2.nodes.RawMatrix().Data)

	largestError := 0.0
	for i, newval := range newvals[:] {
		error := math.Abs(newval - correctVal[i])
		if error > largestError {
			largestError = error
		}
	}

	if epsilonError > tol {
		t.Errorf("The deviation in epsilon is to large")
	}

	if largestError > tol {
		t.Errorf("The deviation in calculates values is to large")
	}

	if floats.Max(nodesError) > tol {
		t.Errorf("The deviation in nodes value is to large")
	}

}

// This tests compares the results to the Scipy implementation, using the same input.
func TestNewRBFSTwoArgs(t *testing.T) {
	var correctVal = []float64{1.29685888, 1.0849643, 0.88575242, 0.74510041, 0.72742949, 0.9092057,
		1.28164769, 1.75256262, 2.24940739, 2.7319236, 1.24244117, 1.03736418,
		0.84980685, 0.72791679, 0.73161543, 0.92702025, 1.30459246, 1.78049402,
		2.2839795, 2.77055211, 1.1866455, 0.9924245, 0.82383574, 0.73035893,
		0.76444862, 0.97479536, 1.35142446, 1.82627493, 2.33264766, 2.82075493,
		1.12653986, 0.94695682, 0.80389519, 0.74553105, 0.81475881, 1.04106492,
		1.414463, 1.88501429, 2.39147789, 2.87895508, 1.05972554, 0.89844695,
		0.78651245, 0.76637037, 0.87072292, 1.1135352, 1.48529336, 1.95130241,
		2.45605203, 2.94105946, 0.98618895, 0.84660049, 0.7700803, 0.7880233,
		0.92382684, 1.18305082, 1.55711234, 2.02038322, 2.52237703, 3.0033898,
		0.91033097, 0.79475304, 0.75578568, 0.80924723, 0.97108716, 1.24585295,
		1.62620332, 2.08899974, 2.58760598, 3.06348094, 0.84124, 0.74996236,
		0.74756281, 0.83235924, 1.0148486, 1.3034511, 1.69211047, 2.15564543,
		2.6501735, 3.12009252, 0.79056967, 0.72135755, 0.75099023, 0.86192263,
		1.06062311, 1.36056798, 1.75670449, 2.22027607, 2.70945771, 3.17261022,
		0.76893489, 0.71735891, 0.77156144, 0.90287885, 1.11436347, 1.42252908,
		1.82277045, 2.28377275, 2.76550023, 3.22076721}

	var correctArg = [][]float64{{5.1, 2.3}, {5.555555555555555, 2.3}, {6.011111111111111, 2.3}, {6.466666666666666, 2.3}, {6.922222222222222, 2.3}, {7.377777777777777, 2.3}, {7.833333333333332, 2.3}, {8.288888888888888, 2.3}, {8.744444444444444, 2.3}, {9.2, 2.3}, {5.1, 2.5}, {5.555555555555555, 2.5}, {6.011111111111111, 2.5}, {6.466666666666666, 2.5}, {6.922222222222222, 2.5}, {7.377777777777777, 2.5}, {7.833333333333332, 2.5}, {8.288888888888888, 2.5}, {8.744444444444444, 2.5}, {9.2, 2.5}, {5.1, 2.6999999999999997}, {5.555555555555555, 2.6999999999999997}, {6.011111111111111, 2.6999999999999997}, {6.466666666666666, 2.6999999999999997}, {6.922222222222222, 2.6999999999999997}, {7.377777777777777, 2.6999999999999997}, {7.833333333333332, 2.6999999999999997}, {8.288888888888888, 2.6999999999999997}, {8.744444444444444, 2.6999999999999997}, {9.2, 2.6999999999999997}, {5.1, 2.9}, {5.555555555555555, 2.9}, {6.011111111111111, 2.9}, {6.466666666666666, 2.9}, {6.922222222222222, 2.9}, {7.377777777777777, 2.9}, {7.833333333333332, 2.9}, {8.288888888888888, 2.9}, {8.744444444444444, 2.9}, {9.2, 2.9}, {5.1, 3.0999999999999996}, {5.555555555555555, 3.0999999999999996}, {6.011111111111111, 3.0999999999999996}, {6.466666666666666, 3.0999999999999996}, {6.922222222222222, 3.0999999999999996}, {7.377777777777777, 3.0999999999999996}, {7.833333333333332, 3.0999999999999996}, {8.288888888888888, 3.0999999999999996}, {8.744444444444444, 3.0999999999999996}, {9.2, 3.0999999999999996}, {5.1, 3.3}, {5.555555555555555, 3.3}, {6.011111111111111, 3.3}, {6.466666666666666, 3.3}, {6.922222222222222, 3.3}, {7.377777777777777, 3.3}, {7.833333333333332, 3.3}, {8.288888888888888, 3.3}, {8.744444444444444, 3.3}, {9.2, 3.3}, {5.1, 3.5}, {5.555555555555555, 3.5}, {6.011111111111111, 3.5}, {6.466666666666666, 3.5}, {6.922222222222222, 3.5}, {7.377777777777777, 3.5}, {7.833333333333332, 3.5}, {8.288888888888888, 3.5}, {8.744444444444444, 3.5}, {9.2, 3.5}, {5.1, 3.6999999999999997}, {5.555555555555555, 3.6999999999999997}, {6.011111111111111, 3.6999999999999997}, {6.466666666666666, 3.6999999999999997}, {6.922222222222222, 3.6999999999999997}, {7.377777777777777, 3.6999999999999997}, {7.833333333333332, 3.6999999999999997}, {8.288888888888888, 3.6999999999999997}, {8.744444444444444, 3.6999999999999997}, {9.2, 3.6999999999999997}, {5.1, 3.8999999999999995}, {5.555555555555555, 3.8999999999999995}, {6.011111111111111, 3.8999999999999995}, {6.466666666666666, 3.8999999999999995}, {6.922222222222222, 3.8999999999999995}, {7.377777777777777, 3.8999999999999995}, {7.833333333333332, 3.8999999999999995}, {8.288888888888888, 3.8999999999999995}, {8.744444444444444, 3.8999999999999995}, {9.2, 3.8999999999999995}, {5.1, 4.1}, {5.555555555555555, 4.1}, {6.011111111111111, 4.1}, {6.466666666666666, 4.1}, {6.922222222222222, 4.1}, {7.377777777777777, 4.1}, {7.833333333333332, 4.1}, {8.288888888888888, 4.1}, {8.744444444444444, 4.1}, {9.2, 4.1}}

	var correctEpsilon = 0.9055385138137416

	var correctNodes = []float64{1.09301649, -0.12008528, -0.537131, -0.50437023, 0.32338213, 0.05131842,
		-0.24681051, -0.2293633, 0.84371205}

	args := [][]float64{{7.15, 2.3}, {5.1, 2.3}, {5.1, 3.2}, {7.15, 3.2}, {7.15, 4.1}, {9.2, 2.3}, {9.2, 3.2}, {9.2, 4.1}, {5.1, 4.1}}
	vals := []float64{0.79065774, 1.29685888, 1.02364173, 1.00750561, 1.25576945, 2.7319236,
		2.97238382, 3.22076721, 0.76893489}

	rbf2,err := NewRBF(args, vals)
	if err != nil{
		t.Errorf("Could not make a interpolator")
	}
	newvals := rbf2.At(correctArg)

	epsilonError := math.Abs(correctEpsilon - rbf2.epsilon)
	nodesError := correctNodes
	floats.Sub(nodesError, rbf2.nodes.RawMatrix().Data)

	largestError := 0.0
	for i, newval := range newvals[:] {
		error := math.Abs(newval - correctVal[i])
		if error > largestError {
			largestError = error
		}
	}

	if epsilonError > tol {
		t.Errorf("The deviation in epsilon is to large")
	}

	if largestError > tol {
		t.Errorf("The deviation in calculates values is to large")
	}

	if floats.Max(nodesError) > tol {
		t.Errorf("The deviation in nodes value is to large")
	}

}

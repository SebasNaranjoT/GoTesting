package simulator

import (
	"fmt"
	"math/big"
)

type simulatorMock struct {
	Spy bool
	result bool
}

func NewSimulatorMock() *simulatorMock {
	return &simulatorMock{}
}

func (r *simulatorMock) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	r.Spy = true

	return r.result
}

func (r *simulatorMock) GetLinearDistance(position [2]float64) float64 {
	x := big.NewFloat(position[0])
	y := big.NewFloat(position[1])
	z := x.Add(x.Mul(x, x), y.Mul(y, y))
	res, _ := z.Sqrt(z).Float64()
	fmt.Printf("Distance: %.2f meters\n", res)
	return res
}  
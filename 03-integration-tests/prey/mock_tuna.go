package prey

type mockTuna struct {
	speed float64
	spy bool
}

func NewMockTuna () *mockTuna {
	return &mockTuna{}
}

func (tuna *mockTuna) GetSpeed() float64 {
	tuna.spy = true
	return tuna.speed
}
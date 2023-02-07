package prey

type preyStub struct {

}

func NewPreyStub () *preyStub {
	return &preyStub{}
}

func (ps *preyStub) GetSpeed() float64 {
	return 5.0
}
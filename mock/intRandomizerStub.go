package mock

// IntRandomizerStub -
type IntRandomizerStub struct {
	IntnCalled func(n int) (int, error)
}

// Intn -
func (irs *IntRandomizerStub) Intn(n int) (int, error) {
	if irs.IntnCalled != nil {
		return irs.IntnCalled(n)
	}

	return 0, nil
}

// IsInterfaceNil -
func (irs *IntRandomizerStub) IsInterfaceNil() bool {
	return irs == nil
}

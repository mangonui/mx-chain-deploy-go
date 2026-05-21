package generate

// IntRandomizer interface provides functionality over generating integer numbers
type IntRandomizer interface {
	Intn(n int) (int, error)
	IsInterfaceNil() bool
}

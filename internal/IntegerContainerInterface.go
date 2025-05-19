package internal

type IntegerContainer interface {
	Add(value int) int
	Delete(value int) bool
}

type AbstractIntegerContainer struct {
	IntegerContainer
}

func (a *AbstractIntegerContainer) Add(value int) int {
	// default implementation
	return 1 + 1
}

func (a *AbstractIntegerContainer) Delete(value int) bool {
	// default implementation
	return false
}

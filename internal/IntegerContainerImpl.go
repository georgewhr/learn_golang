package internal

type IntegerContainerImpl struct {
	AbstractIntegerContainer
	backingSlice []int
}

func NewIntegerContainerImpl() *IntegerContainerImpl {
	// IntegerContainer := IntegerContainerImpl{AbstractIntegerContainer{}}
	IntegerContainer := IntegerContainerImpl{backingSlice: []int{}}
	return &IntegerContainer
}

func (a *IntegerContainerImpl) Add(value int) int {
	// default implementation
	a.backingSlice = append(a.backingSlice, value)
	return len(a.backingSlice)
}

func (a *IntegerContainerImpl) Delete(value int) bool {
	// default implementation
	for index, val := range a.backingSlice {
		if val == value {
			a.backingSlice = append(a.backingSlice[:index], a.backingSlice[index+1:]...)
			return true
		}
	}
	return false
}

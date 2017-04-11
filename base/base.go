package base

// Structure | base interface all structures implement
type Structure interface {
	Empty() bool
	Size() int
	Values() []interface{}
	Clear()
}

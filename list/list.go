package list

import "github.com/sunny-mittal/data-structures/base"

type node struct {
	val interface{}
}

// List | base List interface
type List interface {
	base.Structure
	Insert(vals ...interface{}) int
}

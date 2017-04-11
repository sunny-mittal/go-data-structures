package list

import (
	"testing"

	"github.com/sunny-mittal/data-structures/utils"
)

func TestInsert(t *testing.T) {
	l := new(SinglyLinked)
	l.Insert('a')
	if l.Size() != 1 {
		t.Errorf("Size() expected to return %d, got %d", 1, l.Size())
	}
	l.Insert('b', 'c')
	if l.Size() != 3 {
		t.Logf("Size() expected to return %d, got %d", 3, l.Size())
	}
}

func TestValues(t *testing.T) {
	l := new(SinglyLinked)
	l.Insert(1, 2, 3)
	if !utils.EqualSlices(l.Values(), []interface{}{3, 2, 1}) {
		t.Errorf("Values() expected to return %v, got %v", []interface{}{3, 2, 1}, l.Values())
	}
}

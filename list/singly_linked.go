package list

type sllNode struct {
	node
	next *sllNode
}

// SinglyLinked | basic singly-linked list implementation
type SinglyLinked struct {
	List
	length int
	head   *sllNode
}

// Insert | inserts a value or values into the list
func (l *SinglyLinked) Insert(values ...interface{}) int {
	for _, v := range values {
		n := new(sllNode)
		n.val, n.next = v, l.head
		l.head = n
		l.length++
	}
	return l.length
}

// Values | returns the values currently in the list
func (l *SinglyLinked) Values() []interface{} {
	var values []interface{}
	head := l.head
	for i := 0; i < l.length; i++ {
		values = append(values, head.val)
		head = head.next
	}
	return values
}

// Empty | returns whether or not the list contains any elements
func (l *SinglyLinked) Empty() bool {
	return l.length == 0
}

// Size | returns the current length of the list
func (l *SinglyLinked) Size() int {
	return l.length
}

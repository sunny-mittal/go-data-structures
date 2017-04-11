package list

type sllNode struct {
	node
	next *sllNode
}

type SinglyLinked struct {
	length int
	head   *sllNode
}

func (l *SinglyLinked) Insert(values ...interface{}) int {
	for _, v := range values {
		n := new(sllNode)
		n.val, n.next = v, l.head
		l.head = n
		l.length++
	}
	return l.length
}

func (l *SinglyLinked) Values() []interface{} {
	var values []interface{}
	head := l.head
	for i := 0; i < l.length; i++ {
		values = append(values, head.val)
		head = head.next
	}
	return values
}

func (l *SinglyLinked) Empty() bool {
	return l.length == 0
}

func (l *SinglyLinked) Size() int {
	return l.length
}

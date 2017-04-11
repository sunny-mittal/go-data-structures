package list

type dllNode struct {
	node
	prev *dllNode
	next *dllNode
}

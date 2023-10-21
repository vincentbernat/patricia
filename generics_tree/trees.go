// Code generated by automation. DO NOT EDIT

package generics_tree

// code common to the IPv4/IPv6 trees

// MatchesFunc[T] is called to check if tag data matches the input value
type MatchesFunc[T any] func(payload T, val T) bool

// FilterFunc[T] is called on each result to see if it belongs in the resulting set
type FilterFunc[T any] func(payload T) bool

// treeIteratorNext is an indicator to know what Next() should return
// for the current node.
type treeIteratorNext int

const (
	nextSelf treeIteratorNext = iota
	nextLeft
	nextRight
	nextUp
)

// deleteNodeResult is the return type for deleteNode() function
type deleteNodeResult int

const (
	notDeleted deleteNodeResult = iota
	deletedNodeReplacedByChild
	deletedNodeParentReplacedBySibling
	deletedNodeJustRemoved
)

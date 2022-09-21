// Code generated by automation. DO NOT EDIT

package int_tree

// code common to the IPv4/IPv6 trees

// MatchesFunc is called to check if tag data matches the input value
type MatchesFunc func(payload int, val int) bool

// FilterFunc is called on each result to see if it belongs in the resulting set
type FilterFunc func(payload int) bool

// UpdatesFunc is called to update the tag value
type UpdatesFunc func(payload int) int

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

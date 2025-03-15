package lists

type Ordered interface {
	~int | ~float64 | ~string // Add other ordered types if needed
}

// singly-linked list
type Node[T Ordered] struct {
	Val  T
	Next *Node[T]
}

func MergeKLists[T Ordered](lists []*Node[T]) *Node[T] {
	if lists == nil {
		return nil
	}

	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	mid := len(lists) / 2
	return merge(MergeKLists(lists[0:mid]), MergeKLists(lists[mid:]))
}

func merge[T Ordered](a *Node[T], b *Node[T]) *Node[T] {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	if a.Val <= b.Val {
		a.Next = merge(a.Next, b)
		return a
	} else {
		b.Next = merge(a, b.Next)
		return b
	}
}

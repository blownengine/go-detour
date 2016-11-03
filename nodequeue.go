package detour

import (
	"unsafe"

	"github.com/aurelien-rainone/assertgo"
)

type dtNodeQueue struct {
	m_heap     []*DtNode
	m_capacity int32
	m_size     int32
}

func newDtNodeQueue(n int32) *dtNodeQueue {
	q := &dtNodeQueue{}

	q.m_capacity = n
	assert.True(q.m_capacity > 0, "dtNodeQueue capacity must be > 0")

	q.m_heap = make([]*DtNode, q.m_capacity+1)
	assert.True(len(q.m_heap) > 0, "allocation error")

	return q
}

func (q *dtNodeQueue) bubbleUp(i int32, node *DtNode) {
	parent := (i - 1) / 2
	// note: (index > 0) means there is a parent
	for (i > 0) && (q.m_heap[parent].Total > node.Total) {
		q.m_heap[i] = q.m_heap[parent]
		i = parent
		parent = (i - 1) / 2
	}
	q.m_heap[i] = node
}

func (q *dtNodeQueue) trickleDown(i int32, node *DtNode) {
	child := (i * 2) + 1
	for child < q.m_size {
		if ((child + 1) < q.m_size) &&
			(q.m_heap[child].Total > q.m_heap[child+1].Total) {
			child++
		}
		q.m_heap[i] = q.m_heap[child]
		i = child
		child = (i * 2) + 1
	}
	q.bubbleUp(i, node)
}

func (q *dtNodeQueue) clear() {
	q.m_size = 0
}

func (q *dtNodeQueue) top() *DtNode {
	return q.m_heap[0]
}

func (q *dtNodeQueue) pop() *DtNode {
	result := q.m_heap[0]
	q.m_size--
	q.trickleDown(0, q.m_heap[q.m_size])
	return result
}

func (q *dtNodeQueue) push(node *DtNode) {
	q.m_size++
	q.bubbleUp(q.m_size-1, node)
}

func (q *dtNodeQueue) modify(node *DtNode) {
	for i := int32(0); i < q.m_size; i++ {
		if q.m_heap[i] == node {
			q.bubbleUp(i, node)
			return
		}
	}
}

func (q *dtNodeQueue) empty() bool {
	return q.m_size == 0
}

func (q *dtNodeQueue) getMemUsed() int32 {
	return int32(unsafe.Sizeof(*q)) +
		int32(unsafe.Sizeof(DtNode{}))*(q.m_capacity+1)
}

func (q *dtNodeQueue) getCapacity() int32 {
	return q.m_capacity
}
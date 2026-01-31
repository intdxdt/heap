package heap

import "github.com/intdxdt/cmp"

// Heap - heap data structure
// `reference` : http://eloquentjavascript.net/appendix2.html
// eloquent JS (Appendix 2:Binary Heaps)
// Type parameter specifies what type of heap max or min
type Heap struct {
	htype   *HeapType
	content []interface{}
	cmp     cmp.Compare
}

// NewHeap - Creates a new binary heap
func NewHeap(comparator cmp.Compare, heapType ...*HeapType) *Heap {
	var htype = NewHeapType().AsMin()
	if len(heapType) > 0 {
		htype = heapType[0]
	}

	if !(htype.IsMax() || htype.IsMin()) {
		panic("invalid heap type")
	}
	return &Heap{
		htype:   htype,
		content: make([]interface{}, 0, 32),
		cmp:     comparator,
	}
}

// Clone heap
func (hp *Heap) Clone() *Heap {
	content := make([]interface{}, 0, len(hp.content))
	for i := range hp.content {
		content = append(content, hp.content[i])
	}
	return &Heap{
		htype:   hp.htype.Clone(),
		content: content,
	}
}

// Size of heap
// returns {Number} - heap length
func (hp *Heap) Size() int {
	return len(hp.content)
}

// IsEmpty - checks heap emptiness
func (hp *Heap) IsEmpty() bool {
	return hp.Size() == 0
}

// data view
func (hp *Heap) View() []interface{} {
	return hp.content
}

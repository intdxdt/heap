package heap

// HeapType - Type
type HeapType struct {
	htype int
}

// NewHeapType - create a new heap type
func NewHeapType() *HeapType {
	return &HeapType{htype: -1}
}

// Clone - clone heap type
func (ht *HeapType) Clone() *HeapType {
	return &HeapType{htype: ht.htype}
}

// IsMax - is max type
func (ht *HeapType) IsMax() bool {
	return ht.htype > 0
}

// IsMin - is min type
func (ht *HeapType) IsMin() bool {
	return ht.htype < 0
}

// AsMin - make heap type min type
func (ht *HeapType) AsMin() *HeapType {
	ht.htype = -1
	return ht
}

// AsMax - make heap type max type
func (ht *HeapType) AsMax() *HeapType {
	ht.htype = 1
	return ht
}

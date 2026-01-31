package heap

// Push item on to the heap.
func (hp *Heap) Push(element interface{}) *Heap {
	hp.content = append(hp.content, element)
	// bubble up.
	hp.BubbleUp(len(hp.content) - 1)
	return hp
}

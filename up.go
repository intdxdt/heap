package heap

// BubbleUp - swaps element at index with parent upward the heap
// param index {Number} - integer index
func (hp *Heap) BubbleUp(index int) {
	var element = hp.content[index]
	var found = false

	for !found && index > 0 {
		// When at 0, an element can not go up any further.
		var parentIndex = ((index + 1) / 2) - 1
		var parent = hp.content[parentIndex]
		if hp.htype.IsMin() {
			if hp.cmp(element, parent) >= 0 {
				found = true
			}
		} else {
			if hp.cmp(element, parent) <= 0 {
				found = true
			}
		}
		if !found {
			//swap the parent with element
			hp.content[parentIndex] = element
			hp.content[index] = parent
			index = parentIndex
		}
	}
}

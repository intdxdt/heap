package heap

// SinkDown - swaps element at index with child downward the heap
// param  index {Number} - integer index
func (hp *Heap) SinkDown(index int) {
	var child1 interface{}
	// Look up the target element and its score.
	var length = len(hp.content)
	var element = hp.content[index]
	var found = false

	for !found {
		// Compute the indices of the child elements.
		var indexChild2 = (index + 1) * 2
		var indexChild1 = indexChild2 - 1
		var swapIndex = nilIndex
		// if first child exists (is inside the array)
		if indexChild1 < length {
			// Look it up and compute its score.
			child1 = hp.content[indexChild1]
			// If the score is less than our element"s, we need to swap_index.
			if hp.htype.IsMin() {
				if hp.cmp(child1, element) < 0 {
					swapIndex = indexChild1
				}
			} else {
				if hp.cmp(child1, element) > 0 {
					swapIndex = indexChild1
				}
			}
		}
		// Do the same checks for the other child.
		if indexChild2 < length {
			var child2 = hp.content[indexChild2]

			var tempElement interface{}
			//update swap_index candidate, with elemscore || ch1score
			if isNil(swapIndex) {
				tempElement = element
			} else {
				tempElement = child1
			}

			if hp.htype.IsMin() {
				if hp.cmp(child2, tempElement) < 0 {
					swapIndex = indexChild2
				}
			} else {
				//max heap
				if hp.cmp(child2, tempElement) > 0 {
					swapIndex = indexChild2
				}
			}
		}
		// No need to swap_index further
		if isNil(swapIndex) {
			found = true
		} else {
			// swap_index and continue.
			hp.content[index] = hp.content[swapIndex]
			hp.content[swapIndex] = element
			index = swapIndex
		}
	}
}

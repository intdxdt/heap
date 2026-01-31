package heap

// Pop pops min/max score element from the heap
// <br> <i>returns type used to build heap</i>
// returns {*}
func (hp *Heap) Pop() interface{} {
	if hp.Size() == 0 {
		return nil
	}
	var end interface{}
	// store for return later
	var result = hp.content[0]
	end, hp.content = pop(hp.content)
	// if size > 0 , insert end at index 0, let it sink
	if !hp.IsEmpty() {
		hp.content[0] = end
		hp.SinkDown(0)
	}
	return result
}

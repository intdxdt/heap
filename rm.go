package heap

// Remove node from the heap<br>
// hp does O(N) search
// param  node {*} - type used to create heap
func (hp *Heap) Remove(node interface{}) {
	var length = len(hp.content)
	var i = 0
	var found = false
	var end interface{}

	for !found && i < length {
		if node == hp.content[i] {
			found = true
			end, hp.content = pop(hp.content)
			if i != (length - 1) {
				hp.content[i] = end
				hp.BubbleUp(i)
				hp.SinkDown(i)
			}
		} else {
			i += 1
		}
	}
}

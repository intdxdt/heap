package heap

import (
	"fmt"
	"sort"
	"testing"

	"github.com/franela/goblin"
	"github.com/intdxdt/cmp"
)

func TestHeap(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Heap", func() {
		var _data = []float64{
			0, 1, 0, 2, 3, 4, 3, 3.3, 9, 29, 3.1, 0.1, 1.1,
			1.81, 0.91, 0.81, 0.71, 0.88, 0.82, 0.81,
		}
		var data = make([]float64, 0)
		for _, d := range _data {
			data = append(data, d)
		}

		var minHeap = NewHeap(cmp.F64)

		var maxHeap = NewHeap(cmp.F64, NewHeapType().AsMax())
		//populate min and max heap
		for _, d := range data {
			minHeap.Push(d)
			maxHeap.Push(d)
		}

		g.It("should test items in the heap as queue", func() {
			var minHeapClone = minHeap.Clone()
			g.Assert(len(minHeap.View())).Equal(len(data))
			g.Assert(len(maxHeap.View())).Equal(len(data))
			g.Assert(len(minHeapClone.View())).Equal(len(data))

			//, 'max heap desc sort'
			g.Assert(minHeap.Peek()).Eql(0.)
			g.Assert(minHeapClone.Peek()).Eql(0.)

			g.Assert(minHeap.htype).Eql(minHeapClone.htype)
			g.Assert(minHeap.content).Eql(minHeapClone.content)

			//, 'max heap desc sort'
			g.Assert(maxHeap.Peek()).Eql(29.0)

			fmt.Println(data)

			var minQueue []float64
			var maxQueue []float64

			for minHeap.Size() > 0 {
				minQueue = append(minQueue, float64(minHeap.Pop().(float64)))
			}
			for maxHeap.Size() > 0 {
				maxQueue = append(
					maxQueue,
					float64(maxHeap.Pop().(float64)),
				)
			}

			//min heap queue sort
			sort.Float64s(_data)
			g.Assert(minQueue).Eql(_data)

			//max heap desc sort
			sort.Sort(sort.Reverse(sort.Float64Slice(_data)))
			g.Assert(maxQueue).Eql(_data)

			//rebuild heap
			for _, d := range data {
				minHeap.Push(d)
				maxHeap.Push(d)
			}
			//remove
			minHeap.Remove(29.0)

			//, 'removing element at the end does not affect heap'
			g.Assert(minHeap.content[8]).Equal(0.82)

			minHeap.Remove(1.)

			g.Assert(minHeap.content[len(minHeap.content)-1]).Equal(9.0)

			//empty heap
			for minHeap.Size() > 0 {
				minHeap.Pop()
			}
			for maxHeap.Size() > 0 {
				maxHeap.Pop()
			}
			g.Assert(minHeap.Peek() == nil).IsTrue()
			g.Assert(minHeap.Pop() == nil).IsTrue()
			g.Assert(maxHeap.Pop() == nil).IsTrue()
		})

		g.It("should  panic with invalid heaptype", func() {
			defer func() {
				r := recover()
				g.Assert(r != nil).IsTrue()
			}()
			NewHeap(cmp.F64, &HeapType{})
		})

		g.It("should test util", func() {
			N := 500000
			content := make([]interface{}, 0)
			for i := 0; i < N; i++ {
				content = append(content, i)
			}
			var v interface{}
			for i := 0; i < N; i++ {
				v, content = pop(content)
				g.Assert(v).Eql(N - i - 1)
			}
			g.Assert(len(content)).Equal(0)
			v, content = pop(content)
			g.Assert(v == nil).IsTrue()
		})
	})
}

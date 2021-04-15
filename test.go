package test

import "container/heap"

func Good1(someHeap heap.Interface) {

  someHeap.Push(0)
  someHeap.Pop()

}

func Good2(someHeap heap.Interface) {

  heap.Push(someHeap, 0)
  heap.Pop(someHeap)

}

func Bad (someHeap heap.Interface) {

  someHeap.Push(0)
  heap.Pop(someHeap)

}

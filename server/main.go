package main

import "fmt"

type Element interface {
}

type Queue interface {
	Add(e Element)
	Poll() Element
	Clear() bool
	Size() int
	IsEmpty() bool
}

type QueueElement struct {
	element []Element
}

func New() *QueueElement {
	return &QueueElement{}
}

func (entity *QueueElement) Add(e Element) {
	entity.element = append(entity.element, e)
}

func (entity *QueueElement) IsEmpty() bool {
	return len(entity.element) == 0
}

func (entity *QueueElement) Poll() (e Element) {
	if entity.IsEmpty() {
		return nil
	}
	e = entity.element[0]
	entity.element = entity.element[1:]
	return e
}

func (entity *QueueElement) Clear() bool {
	if entity.IsEmpty() {
		return false
	}
	for i := 0; i < entity.Size(); i++ {
		entity.element[i] = nil
	}
	entity.element = nil
	return true
}

func (entity *QueueElement) Size() int {
	return len(entity.element)
}

func main() {
	queue := New()
	fmt.Printf("address: %p\n", &queue)
	fmt.Printf("queue element: %p\n", &queue.element)
	for i := 1; i <= 50; i++ {
		queue.Add(i)
		fmt.Printf("element %d: %p\n", i-1, &queue.element[i-1])
	}

	fmt.Println("size:", queue.Size())
	fmt.Println("poll first element:", queue.Poll())
	fmt.Println("size:", queue.Size())
	fmt.Println("clear:", queue.Clear())
	fmt.Println("is empty:", queue.IsEmpty())
	for i := 1; i < 50; i++ {
		queue.Add(i)
	}

	fmt.Println("poll first element:", queue.Poll())
	fmt.Println("size:", queue.Size())
	fmt.Println("clear:", queue.Clear())
	fmt.Println("is empty:", queue.IsEmpty())

	var test [5] int
	fmt.Printf("fisrt address: %p\n", &test)

	for i := 0; i < len(test); i++ {
		fmt.Printf("test[%d]: %p\n", i, &test[i])
	}
}

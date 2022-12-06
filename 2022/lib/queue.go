package lib

import "fmt"

type Queue struct {
	queue []string
}

// Add element to queue
func (q Queue) Push(element string) Queue {
	q.queue = append(q.queue, element)
	return q
}

func (q Queue) PushPlus(elements []string) Queue {
	q.queue = append(q.queue, elements...)
	return q
}

// Returns elment 0 of the queue
// First in first out function
func (q *Queue) Get() (string, error) {
	if len(q.queue) < 1 {
		return "", fmt.Errorf("no items in queue to get")
	}

	element := q.queue[0]
	q.queue = q.queue[1:]
	return element, nil
}

// Pop element from queue
// First in last out function
func (q Queue) Pop() (string, Queue, error) {
	if len(q.queue) < 1 {
		return "", q, fmt.Errorf("no items in queue to get")
	}

	element := q.queue[len(q.queue)-1]
	q.queue = q.queue[:len(q.queue)-1]
	return element, q, nil
}

func (q Queue) PopPlus(count int) ([]string, Queue, error) {
	if len(q.queue) < count {
		return nil, q,
			fmt.Errorf("count (%v) exceeds number of items (%v) in queue", count, len(q.queue))
	}

	elements := q.queue[len(q.queue)-count:]
	q.queue = q.queue[:len(q.queue)-count]
	return elements, q, nil

}

func (q Queue) PreInsert(element string) Queue {
	var temp []string
	temp = append(temp, element)
	temp = append(temp, q.queue...)
	q.queue = temp
	return q
}

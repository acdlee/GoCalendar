package main

import "fmt"

const SUCCESS = 0
const FAILURE = -1

type event struct {
	name                         string
	start_time, duration_minutes int
	next                         *event
}

type calendar struct {
	name   string
	events **event
}

func main() {
	var event1 *event = &event{
		name:             "Event 1",
		start_time:       10,
		duration_minutes: 20,
		next:             nil,
	}
	event2 := &event{
		name:             "Event 2",
		start_time:       5,
		duration_minutes: 30,
		next:             event1,
	}

	cal := &calendar{
		events: &event2,
		name:   "My Calendar",
	}

	fmt.Printf("Events in %v:\n", cal.name)
	for curr_event := *cal.events; curr_event != nil; curr_event = curr_event.next {
		fmt.Println(curr_event.name)
	}
}

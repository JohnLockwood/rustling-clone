package main

import "fmt"
import tb "github.com/nsf/termbox-go"

func example8() {
	fmt.Println("Press a key?")
	arr := make([]byte, 5)
	event := tb.PollRawEvent(arr)

	fmt.Println("An event: %v", event)
}

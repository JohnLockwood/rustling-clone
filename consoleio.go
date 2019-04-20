
package main

import (
"bufio"
"fmt"
"os"
"time"

tm "github.com/buger/goterm"
)

func example2() {
	tm.Clear()

	// Create Box with 30% width of current screen, and height of 20 lines
	box := tm.NewBox(30|tm.PCT, 20, 0)

	// Add some content to the box
	// Note that you can add ANY content, even tables
	fmt.Fprint(box, "Some box content")

	// Move Box to approx center of the screen
	tm.Print(tm.MoveTo(box.String(), 40|tm.PCT, 40|tm.PCT))

	tm.Flush()
}

func example1() {
	tm.Clear() // Clear current screen

	for {
		// By moving cursor to top-left position we ensure that console output
		// will be overwritten each time, instead of adding new.
		tm.MoveCursor(1, 1)

		tm.Println("Current Time:", time.Now().Format(time.RFC1123))

		tm.Flush() // Call it every time at the end of rendering

		time.Sleep(time.Second)
	}
}

func example3_bufio_only() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	// print out the unicode value i.e. A -> 65, a -> 97
	fmt.Println(char)

	switch char {
	case 'A':
		fmt.Println("A Key Pressed")
		break
	case 'a':
		fmt.Println("a Key Pressed")
		break
	}
}

func printMenu() {
	// By moving cursor to top-left position we ensure that console output
	// will be overwritten each time, instead of adding new.
	tm.MoveCursor(1, 1)
	tm.Print("Next : Go to next Screen")
	tm.MoveCursor(1, 2)
	tm.Print("Quit : End Program")

	tm.Flush()
}

func getInput() string {
	tm.MoveCursor(1, 4)
	tm.Flush()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()

}

func printInput(r string) {
	tm.MoveCursor(1, 4)
	tm.Print("    ")
	tm.Flush()
	tm.MoveCursor(1, 3)
	tm.Print(fmt.Sprintf("You selected %v", r))
	tm.Flush()
}

func example4() {

	tm.Clear() // Clear current screen

	for {

		printMenu()
		rune := getInput()
		printInput(rune)
		if rune == "q" || rune == "Q" {
			tm.MoveCursor(1, 8)
			tm.Print("")
			tm.Flush()
			break
		}

		tm.Flush() // Call it every time at the end of rendering

	}
	fmt.Println("Thanks for learning Go!")
}

func example5() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}


package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"os"
	"strings"
)

type Handler func() error

type MenuItem struct {
	short string
	long string
	description string
	handler Handler
}


func quitHandler() error {
	os.Exit(0)
	return nil
}

func helpHandler() error {
	PrintMenu()
	return nil
}

func nextHandler() error {
	tm.Println("Next")

	tm.Println(tm.Color("./configuration/configuration.go", tm.YELLOW))
	return nil
}

func prevHandler() error {
	tm.Println("Prev")
	return nil
}

var menu []MenuItem
func init() {
	menu = []MenuItem{
		MenuItem{"q", "Quit", "Exit the program", quitHandler},
		MenuItem{"h", "Help", "Display this menu", helpHandler},
		MenuItem{"n", "Next", "Go to next exercise", nextHandler},
		MenuItem{"p", "Previous", "Go to previous exercise", prevHandler},
	}
}


func getShortMenu() string {
	s := ""
	for _, item := range(menu) {
		if len(s) > 0 {
			s += " - "
		}
		s = s + fmt.Sprintf("%v (%v)", item.short, item.long)
	}
	return s
}

func MenuLoop() {
	cmd := ""
	for {


		for _, item:= range(menu) {
			if strings.ToLower(cmd) == strings.ToLower(item.short) || strings.ToLower(cmd) == strings.ToLower(item.long) {
				tm.Clear()
				tm.MoveCursor(1,3)
				item.handler()
				tm.Flush()
			}
		}


		tm.MoveCursor(1, tm.Height()-2)
		tm.Println(getShortMenu())
		tm.Flush()
		fmt.Scanln(&cmd)


	}

}

func PrintMenu() {
	tm.Clear()
	tm.MoveCursor(1,1)
	tm.Println("Enter a shortcut or command, and press enter.")
	tm.Println(" ")
	tm.Println("  Shortcut : Command")
	tm.Println("  --------   -------")
	for _, item := range(menu) {
		tm.Printf("         %v : %v\t\t%v\n", item.short, item.long, item.description)
	}
	tm.Flush()
}
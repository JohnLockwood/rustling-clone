package main

import (
	"fmt"
	"os"
	"strings"
	// tm "github.com/buger/goterm"
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
	fmt.Println("Next")
	return nil
}

func prevHandler() error {
	fmt.Println("Prev")
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



func MenuLoop() {
	for {
		cmd := ""
		fmt.Scanln(&cmd)
		for _, item:= range(menu) {
			if strings.ToLower(cmd) == strings.ToLower(item.short) || strings.ToLower(cmd) == strings.ToLower(item.long) {
				item.handler()
			}
		}
	}

}

func PrintMenu() {

	fmt.Println("Enter a shortcut or command, and press enter.")
	fmt.Println(" ")
	fmt.Println("  Shortcut : Command")
	fmt.Println("  --------   -------")
	for _, item := range(menu) {

		fmt.Printf("         %v : %v\t\t%v\n", item.short, item.long, item.description)
	}
}
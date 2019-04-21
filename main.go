package main

import (
	"fmt"
	"os"
	"rustling-clone/configuration"
)

func loadCourse() {
	loader := configuration.NewCourseLoader(nil)
	course, err := loader.GetCourse()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(course.Language)
}

func main() {
	PrintMenu()
	MenuLoop()
}

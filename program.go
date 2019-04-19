package main

import (
	"fmt"
	"rustling-clone/configuration"
)

func main() {
	fmt.Println("Hello")
	loader := configuration.NewCourseLoader(nil)
	course := loader.GetCourse()
	fmt.Println(course.Language)
}

package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type CourseModule struct {
	File string `json:"file"`
	Mode string `json:"mode"`
}

type CourseSection struct {
	Name string `json:"name"`
	Modules []CourseModule `json:"modules"`
}

// Course is the main data structure that holds information about our course lessons.
// It is read from config.json and should be considered read-only
type Course struct {
	Language string `json:"language"`
	Sections []CourseSection `json:"sections"`
}

type CourseLoader struct {
	dir cwd
}

func NewCourseLoader(dir cwd) *CourseLoader {
	dirLoader := dir
	if dir == nil {
		dirLoader = DirectoryServer{}
	}
	return &CourseLoader{
		dir: dirLoader,
	}
}

func (self *CourseLoader) getCourceConfigurationFile() string {
	return self.dir.getCWD() + "/" + "config.json"
}

// GetCourse loads the course object from the root directory
// TODO more error handling
func (self *CourseLoader) GetCourse() Course {

	// Open our jsonFile
	jsonFile, err := os.Open(self.getCourceConfigurationFile())
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer jsonFile.Close()

	// read our opened json file as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	course := Course{}

	err = json.Unmarshal(byteValue, &course)
	// defer the closing of our jsonFile so that we can parse it later on

	return course

}




package configuration

import (
	"encoding/json"
	"errors"
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

type cwd interface {
	getCWD() string
}


type CourseLoader struct {
	dir cwd
}

// NewCourseLoader returns a CourseLoader with a default cwd interface if cwd is nil
// (which returns the current working directory).  This enables us to inject
// different
func NewCourseLoader(dir cwd) *CourseLoader {
	dirLoader := dir
	if dir == nil {
		dirLoader = DirectoryServer{}
	}
	return &CourseLoader{
		dir: dirLoader,
	}
}

func (self *CourseLoader) getCourseConfigurationFile() string {
	return self.dir.getCWD() + "/" + "config.json"
}

func (self * CourseLoader) configFileExists() bool {
	if fileInfo, err := os.Stat(self.getCourseConfigurationFile()); err == nil {
		return ! fileInfo.IsDir()
	}
	return false
}

// GetCourse loads the course object from the root directory
func (self *CourseLoader) GetCourse() (*Course, error) {

	if(!self.configFileExists()) {
		// TODO use real program name.
		return nil, errors.New("Configuration file not found.  This program must be run from the program directory.")
	}

	// Open our jsonFile
	jsonFile, err := os.Open(self.getCourseConfigurationFile())
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer jsonFile.Close()

	// read our opened json file as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	course := Course{}

	err = json.Unmarshal(byteValue, &course)
	if err != nil {
		return nil, err
	}

	return &course, nil
}




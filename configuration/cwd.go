package configuration

import (
	"log"
	"os"
)

type DirectoryServer struct{}

type cwd interface {
	getCWD() string
}


func (DirectoryServer) getCWD() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

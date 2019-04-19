package configuration

import (
	"log"
	"os"
)

type DirectoryServer struct{}


func (DirectoryServer) getCWD() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

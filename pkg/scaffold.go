package scaffold

import (
	"errors"
	"fmt"
	"os"
)

func Create(name, slug string) error {
	if collisionErr := filenameCollisionCheck(slug, name); collisionErr != nil {
		return collisionErr
	}

	fmt.Println("Create success!")
	return nil
}

func filenameCollisionCheck(slug string, name string) error {
	//TODO improve this filename collision check
	_, err := os.Stat(slug)
	if os.IsNotExist(err) == false {
		msg := fmt.Sprintf("Project name: %s conflicts with existing file", name)
		fmt.Println(msg)
		return errors.New(msg)
	}
	return nil
}

package scaffold

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func Create(name, slug string, log *logrus.Logger) error {
	if collisionErr := filenameCollisionCheck(slug, name); collisionErr != nil {
		log.WithFields(logrus.Fields{
			"Error":   collisionErr,
			"Context": "Checking for filename collisions",
		}).Debug("Filename collision detected")
		return collisionErr
	}

	if mkdirErr := makeDirectory(slug); mkdirErr != nil {
		log.WithFields(logrus.Fields{
			"Error":   mkdirErr,
			"Context": "Create local directory",
		}).Debug("makeDirectory failed")
		return mkdirErr
	}

	if gitInitErr := initializeGitRepo(slug); gitInitErr != nil {
		log.WithFields(logrus.Fields{
			"Error":   gitInitErr,
			"Context": "Cd'ing into target directory and running git init",
		}).Debug("initializeGitRepo failed")
		return gitInitErr
	}

	if _, readmeErr := copyFile("./data/README.md", fmt.Sprintf("./%s/README.md", slug)); readmeErr != nil {
		log.WithFields(logrus.Fields{
			"Error":   readmeErr,
			"Context": "Copying readme file template into target repository",
		}).Debug("copying README file failed")
		return readmeErr
	}

	if _, licenseErr := copyFile("./data/LICENSE", fmt.Sprintf("./%s/LICENSE", slug)); licenseErr != nil {
		log.WithFields(logrus.Fields{
			"Error":   licenseErr,
			"Context": "Copying LICENSE file template into target repository",
		}).Debug("copying LICENSE file failed")
		return licenseErr
	}

	if addFilesErr := gitAddFiles(slug); addFilesErr != nil {
		log.WithFields(logrus.Fields{
			"Error":   addFilesErr,
			"Context": "Adding initial files to git",
		}).Debug("adding files to git failed")
		return addFilesErr
	}

	log.Debug(fmt.Sprintf("Successfully created %s", name))
	return nil
}

func copyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func gitAddFiles(fp string) error {
	cmdString := []string{"add", "README.md", "LICENSE"}
	// Execute git add from within the target directory
	cmd := exec.Command("git", cmdString...)
	cmd.Dir = fp
	fmt.Println(cmd.String())

	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git output: %s, stderr: %s", out, err.Error())
	}
	return nil
}

func initializeGitRepo(fp string) error {
	cmd := exec.Command("git", "init", fp)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func makeDirectory(fp string) error {
	if err := os.Mkdir(fp, 0755); err != nil {
		return err
	}
	return nil
}

func filenameCollisionCheck(slug string, name string) error {
	//TODO improve this filename collision check
	_, err := os.Stat(slug)
	if os.IsNotExist(err) == false {
		msg := fmt.Sprintf("Project name: %s conflicts with existing file", name)
		return errors.New(msg)
	}
	return nil
}

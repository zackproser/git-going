package scaffold

import (
	"errors"
	"fmt"
	"os"

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

	if readmeErr := renderReadmeFile(name, slug); readmeErr != nil {
		log.WithFields(logrus.Fields{
			"Error":   readmeErr,
			"Context": "Rendering README template file to repo",
		}).Debug("renderReadmeFile failed")
	}

	if licenseErr := renderLicenseFile(slug); licenseErr != nil {
		log.WithFields(logrus.Fields{
			"Error":   licenseErr,
			"Context": "Rendering LICENSE template file to repo",
		}).Debug("renderLicensefile failed")
		return licenseErr
	}

	if addFilesErr := gitAddFiles(slug); addFilesErr != nil {
		log.WithFields(logrus.Fields{
			"Error":   addFilesErr,
			"Context": "Adding initial files to git",
		}).Debug("adding files to git failed")
		return addFilesErr
	}

	if commitErr := gitCommit(slug, "Init to win it"); commitErr != nil {
		log.WithFields(logrus.Fields{
			"Error":   commitErr,
			"Context": "Committing new files to git",
		}).Debug("committing files to git failed")
		return commitErr
	}

	log.Debug(fmt.Sprintf("Successfully created %s", name))
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

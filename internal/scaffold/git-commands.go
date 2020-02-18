package scaffold

import (
	"fmt"
	"os/exec"
)

// PushRepo executes a git push command to push up
// all local files that were scaffolded
func PushRepo(fp string) error {
	cmdString := []string{"push", "origin", "master"}
	cmd := exec.Command("git", cmdString...)
	cmd.Dir = fp

	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git output %v, stderr: %v", out, err.Error())
	}
	return nil
}

// AddRemoteOrigin sets the URL for the 'origin' remote
func AddRemoteOrigin(fp, url string) error {
	cmdString := []string{"remote", "add", "origin", url}
	cmd := exec.Command("git", cmdString...)
	cmd.Dir = fp

	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git output: %v, stderr: %v", out, err.Error())
	}
	return nil
}

func gitCommit(fp, commitMsg string) error {
	cmdString := []string{"commit", "-m", commitMsg}
	cmd := exec.Command("git", cmdString...)
	cmd.Dir = fp

	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git output: %v, stderr: %s", out, err.Error())
	}
	return nil
}

func gitAddFiles(fp string) error {
	cmdString := []string{"add", "README.md", "LICENSE"}
	// Execute git add from within the target directory
	cmd := exec.Command("git", cmdString...)
	cmd.Dir = fp

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

package scaffold

import (
	"fmt"
	"html/template"
	"os"
	"strconv"
	"time"
)

func renderLicenseFile(fp string) error {
	paths := []string{
		"./data/LICENSE.tpl",
	}

	t := template.Must(template.New("LICENSE.tpl").ParseFiles(paths...))
	destination, err := os.Create(fmt.Sprintf("./%s/LICENSE", fp))
	if err != nil {
		return err
	}
	defer destination.Close()

	ld := &LicenseData{
		CopyrightYear: strconv.Itoa(time.Now().Year()),
		Author:        "Zack Proser",
	}

	if renderErr := t.Execute(destination, ld); renderErr != nil {
		return renderErr
	}

	return nil
}

func renderReadmeFile(name, fp string) error {
	paths := []string{
		"./data/README.tpl",
	}

	t := template.Must(template.New("README.tpl").ParseFiles(paths...))
	destination, err := os.Create(fmt.Sprintf("./%s/README.md", fp))
	if err != nil {
		return err
	}
	defer destination.Close()

	rd := &ReadmeData{
		ProjectName: name,
	}

	if renderErr := t.Execute(destination, rd); renderErr != nil {
		return renderErr
	}

	return nil
}
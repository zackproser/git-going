package scaffold

import (
	"fmt"
	"html/template"
	"os"
	"strconv"
	"time"
)

func renderLicenseFile(fp, author string) error {

	t := template.Must(template.New("LICENSE.tpl").Parse(LicenseTemplate))
	destination, err := os.Create(fmt.Sprintf("./%s/LICENSE", fp))
	if err != nil {
		return err
	}
	defer destination.Close()

	ld := &LicenseData{
		CopyrightYear: strconv.Itoa(time.Now().Year()),
		Author:        author,
	}

	if renderErr := t.Execute(destination, ld); renderErr != nil {
		return renderErr
	}

	return nil
}

func renderReadmeFile(name, fp string) error {

	t := template.Must(template.New("README.tpl").Parse(ReadmeTemplate))
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

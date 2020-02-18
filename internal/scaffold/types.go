package scaffold

// ReadmeData is the information required to
// render the README.md file in the new repo
type ReadmeData struct {
	ProjectName string
}

// LicenseData is the information required to
// render the LICENSE file in the new repo
type LicenseData struct {
	CopyrightYear string
	Author        string
}

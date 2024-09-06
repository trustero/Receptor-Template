package receptorPackage

import (
	"embed"
)

// following lines have been added to embed resources
// in the receptor executable

//go:embed resources/*
var embedFS embed.FS

func readEmbedFile(path string) (string, error) {
	// helper function to read files that will be embedded
	// with the receptor executable
	data, err := embedFS.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), err
}

func GetInstructionsImpl() (string, error) {
	// replace this with the name of the file in the
	// resources directory that contains the instructions
	return readEmbedFile("resources/trr-receptorpackage.md")
}

func GetLogoImpl() (string, error) {
	// replace this with the name of the file in the
	// resources directory that contains the logo
	// must be in svg format
	return readEmbedFile("resources/trustero.svg")
}

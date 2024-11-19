package utils

import (
	"errors"
	fp "path/filepath"
	"strings"
)

func GetRepoDir() (string, error) {
	// get current directory
	dir, err := fp.Abs(".")
	if err != nil {
		return "", err
	}

	// walk the directory until we find tipJar
	parents := strings.Split(dir, "/")
	for i := len(parents) - 1; i >= 0; i-- {
		if parents[i] == "tipJar" {
			abs := fp.Join("/", fp.Join(parents[:i+1]...))
			return abs, nil
		}
	}

	return "", errors.New("tipJar directory not found")
}

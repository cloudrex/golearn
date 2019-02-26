package main

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// GetExePath : Retrieve the executable's absolute path.
func GetExePath() string {
	ex, err := os.Executable()

	if err != nil {
		panic(err)
	}

	return filepath.Dir(ex)
}

// MergeExePath : Append input path to executable's absolute path.
func MergeExePath(input ...string) string {
	return path.Join(append(input, GetExePath())...)
}

// DoesFilePathExist : Determine if a file path exists.
func DoesFilePathExist(filePath ...string) bool {
	if _, err := os.Stat(path.Join(filePath...)); os.IsNotExist(err) {
		return false
	}

	return true
}

// Read : Read the contents of a file into a string. Returns an empty string and an error if file does not exist.
func Read(filePath ...string) string {
	result, err := ioutil.ReadFile(path.Join(filePath...))

	if err != nil {
		panic(err)
	}

	return string(result)
}

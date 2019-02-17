package link

import (
	"errors"
	"golearn/lex"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

// ResolveSourceTokens : Reads target file path and returns corresponding tokens.
func ResolveSourceTokens(filePath ...string) ([]lex.Token, error) {
	if !DoesFilePathExist(filePath...) {
		return nil, errors.New("Provided file path does not exist")
	}

	source, err := ReadFile(filePath...)

	return lex.Tokenize(source), err
}

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

// ReadFile : Read the contents of a file into a string. Returns an empty string and an error if file does not exist.
func ReadFile(filePath ...string) (string, error) {
	result, err := ioutil.ReadFile(path.Join(filePath...))

	if err != nil {
		return "", err
	}

	return string(result), nil
}

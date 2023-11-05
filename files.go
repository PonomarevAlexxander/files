package files

import (
	"bytes"
	"os"
)

func Contains(filepath string, substring string) (bool, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return false, err
	}

	return bytes.Contains(file, []byte(substring)), nil
}

package ollama

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ReadTestData(fname string) ([]byte, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	rootPath := filepath.Dir(filepath.Dir(pwd))
	fPath := filepath.Join(rootPath, "testdata", fname)
	data, err := os.ReadFile(fPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func TestEncodeBase64(t *testing.T) {
	testData, err := ReadTestData("test.txt")
	if err != nil {
		assert.Fail(t, "Unable to read test data: %v", err)
	}

	got := EncodeBase64(testData)
	want := "VGhpcyBpcyBhIHRlc3Q="
	assert.Equal(t, want, got, fmt.Sprintf("Want: %v Got: %v", want, got))
}

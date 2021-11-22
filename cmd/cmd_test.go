package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/garciaolais/gownloader/cmd"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	expectedBytes := []byte{24, 108, 90, 204, 81, 189, 102, 126}
	resultBytes := cmd.Hash([]byte{12})

	assert.Equal(t, expectedBytes, resultBytes)
}

func TestCmd(t *testing.T) {
	filePath := filepath.Join(os.TempDir(), "file.dat")
	err := cmd.Run("https://raw.githubusercontent.com/garciaolais/gownload/main/cmd/file12.dat", false, filePath)
	assert.NoError(t, err)

	str, err := cmd.PrintHexFile(filePath)
	assert.Equal(t, "186c5acc51bd667e", str)
}

package cmd

import (
	"log"
	"net"
	"net/http"
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
	http.HandleFunc("/file12.dat", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "file12.dat")
	})

	l, err := net.Listen("tcp", "localhost:8080")
	assert.NoError(t, err)

	go func() {
		log.Fatal(http.Serve(l, nil))
	}()

	filePath := filepath.Join(os.TempDir(), "file.dat")
	err = cmd.Run("http://localhost:8080/file12.dat", false, filePath)
	assert.NoError(t, err)

	str, err := cmd.PrintHexFile(filePath)
	assert.Equal(t, "186c5acc51bd667e", str)
}

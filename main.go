package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/garciaolais/gownloader/cmd"
	"github.com/garciaolais/gownloader/util"
)

func main() {
	urlPtr := flag.String("url", "https://raw.githubusercontent.com/garciaolais/gownload/main/cmd/file12.dat", "")
	pathPtr := flag.String("path", filepath.Join(os.TempDir(), "file.dat"), "a destination path for a file containing a hash")
	throttlingPtr := flag.Bool("throttling", false, "throttling download")
	flag.Parse()

	if err := util.IsURL(*urlPtr); err != nil {
		log.Fatal(err)
	}

	dir := filepath.Dir(*pathPtr)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		log.Fatal(err)
	}

	fmt.Println("GOWNLOAD")

	if err := cmd.Run(*urlPtr, *throttlingPtr, *pathPtr); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

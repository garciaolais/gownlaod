package cmd

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var datachunk int64 = 100
var timelapse time.Duration = 1

func Run(url string, throttling bool, path string) error {
	b, err := downloadFile(url, throttling)
	if err != nil {
		return err
	}

	hashBytes := Hash(b)
	err = CreateFile(hashBytes, path)
	if err != nil {
		return err
	}

	str, err := PrintHexFile(path)
	if err != nil {
		return err
	}
	fmt.Printf("file hex - %s", str)
	return nil
}

func PrintHexFile(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b), nil
}

func downloadFile(url string, throttling bool) ([]byte, error) {
	fmt.Printf("download - %s\n", url)

	var temp bytes.Buffer
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if r.StatusCode != 200 {
		return nil, errors.New(r.Status)
	}

	if throttling {
		for range time.Tick(timelapse * time.Second) {
			fmt.Printf(".")
			_, err := io.CopyN(&temp, r.Body, datachunk)
			if err != nil {
				break
			}
		}
		fmt.Println()
	} else {
		_, err := io.Copy(&temp, r.Body)
		if err != nil {
			return nil, err
		}
	}

	return temp.Bytes(), nil
}

func Hash(data []byte) []byte {
	fmt.Printf("hash - data %v\n", data)
	coefficients := getCoefficients()
	h := make([]byte, 8)
	for _, ib := range data {
		for i := range h {
			if i-1 == -1 {
				h[i] = byte((int(ib) * coefficients[i]) % 255)
			} else {
				h[i] = byte((int(h[i-1]+ib) * coefficients[i]) % 255)
			}
		}
	}

	fmt.Printf("hash - data result %v\n", h)
	return h
}

func CreateFile(hashBytes []byte, path string) error {
	fmt.Printf("create file %s\n", path)
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write(hashBytes)
	if err != nil {
		return err
	}

	return nil
}

func getCoefficients() []int {
	return []int{2, 3, 5, 7, 11, 13, 17, 19}
}

package input_reader

import (
	"fmt"
	"os"
    "path/filepath"
)

func ReadInput(fileLocation string) *os.File {

	readFile, err := os.Open(fileLocation)

	if err != nil {
		fmt.Println(err)
	}

	return readFile
}

func ShowCurrentExeDir() {
    ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    exPath := filepath.Dir(ex)
    fmt.Println(exPath)
}

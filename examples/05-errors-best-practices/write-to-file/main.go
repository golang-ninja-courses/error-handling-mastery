package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var accessLogFilename = filepath.Join(os.TempDir(), "access.log")

func main() {
	fmt.Println(accessLogFilename)

	if err := writeToAccessLog("GET /"); err != nil {
		log.Fatal(err)
	}
}

func writeToAccessLog(data string) error {
	f, err := os.OpenFile(accessLogFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("cannot open file: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString(data + "\n"); err != nil {
		return fmt.Errorf("cannot write data: %v", err)
	}

	// return f.Close()
	// return nil
	return f.Sync()
}

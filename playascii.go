package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		PrintHelpAndExit(1)
	}
	if err := CheckCastFile(os.Args[1]); err != nil {
		os.Exit(1)
	}
	CreateAndBrowse(os.Args[1])
}

func PrintHelpAndExit(exitCode int) {
	Help(os.Stdout)
	os.Exit(exitCode)
}

func Help(out io.Writer) {
	fmt.Fprintf(out, "call with: playascii <file.cast>")
}

func CheckCastFile(file string) error {
	fileinfo, err := os.Stat(file)
	if err != nil {
		fmt.Println("Can not access cast file.")
		return err
	}
	if fileinfo.IsDir() {
		fmt.Println("cast file is a directory.")
		return errors.New("file is directory")
	}
	return nil
}

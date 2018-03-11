package main

import (
	"fmt"
	"github.com/pkg/browser"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func CreateAndBrowse(file string) error {
	dir, path, err := CreateWebPage(file)
	if err != nil {
		return err
	}
	if err := browser.OpenFile(path); err != nil {
		os.Remove(dir)
		return err
	}
	return nil
}

func CreateWebPage(castFile string) (string, string, error) {
	dir, err := ioutil.TempDir(os.TempDir(), "playascii")
	if err != nil {
		log.Fatal(err)
	}

	AddHelperFiles(castFile, dir)

	filepath, file, err := CreateIndexFile(dir)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	template, err := CreateIndexTemplate()
	if err != nil {
		return "", "", fmt.Errorf("creating index template: %s", err.Error())
	}
	_, err = file.WriteString(template)
	return dir, filepath, err
}

func CreateIndexFile(dir string) (string, *os.File, error) {
	file, err := ioutil.TempFile(dir, "index")
	return file.Name(), file, err
}

func addFileAndContent(file, content string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(content)
	f.Close()
}

func AddHelperFiles(castFile, dir string) {
	addFileAndContent(filepath.Join(dir, "asciinema-player.css"), playerCSS)
	addFileAndContent(filepath.Join(dir, "asciinema-player.js"), playerJS())
	addFileAndContent(filepath.Join(dir, "demo.cast"), cast(castFile))
}

func cast(filename string) string {
	cast, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(cast)
}

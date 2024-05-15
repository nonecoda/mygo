package filerename

import (
	"io/fs"
	"log"
	"os"
	"path"
	"strings"
)

const (
	TAG string = ""
	SEP string = ""
)

func Do(dirname string) {
	fileInfos, err := listFilesByDir(dirname)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range fileInfos {
		if strings.Contains(r.Name(), TAG) {
			renameFileName(dirname, r.Name())
		}
	}
}

func listFilesByDir(dirname string) ([]fs.DirEntry, error) {
	fileInfos, err := os.ReadDir(dirname)

	return fileInfos, err
}

func renameFileName(dirname, fileName string) {
	newFileName := strings.Split(fileName, TAG)[0]
	newFileName = strings.TrimSpace(newFileName)
	extList := strings.Split(fileName, ".")
	ext := extList[len(extList)-1]
	oldPath := path.Join(dirname, fileName)
	newPath := path.Join(dirname, newFileName+"."+ext)
	_ = os.Rename(oldPath, newPath)
}

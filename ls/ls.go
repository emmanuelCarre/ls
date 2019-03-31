package ls

import (
	"fmt"
	"os"

	"github.com/emmanuelCarre/ls/io/ioutil"
)

type LsData struct {
	Files      []os.FileInfo
	ShowHidden bool
}

func (d *LsData) listFiles(util ioutil.IoUtil, path string) error {
	var err error
	d.Files, err = util.ReadDir(path)
	if err != nil {
		return err
	}
	return nil
}

func (d *LsData) display() {
	for _, aFile := range d.Files {
		if aFile.Name()[0] == '.' && !d.ShowHidden {
			continue
		}
		fmt.Println(aFile.Name())
	}
}

func Ls(util ioutil.IoUtil, data LsData, path string) error {
	err := data.listFiles(util, path)
	data.display()
	return err
}

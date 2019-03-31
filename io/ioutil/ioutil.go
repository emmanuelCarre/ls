package ioutil

import (
	"io/ioutil"
	"os"
)

type IoUtil interface {
	ReadDir(dirname string) ([]os.FileInfo, error)
}

type IoUtilImpl struct{}

func (IoUtilImpl) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

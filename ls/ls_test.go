package ls

import (
	"errors"
	"os"
	"testing"

	"github.com/emmanuelCarre/ls/mocks"
	"github.com/emmanuelCarre/ls/mocks/os"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestIoIsCall(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var data LsData
	path := "./"
	m := mock_ioutil.NewMockIoUtil(ctrl)

	m.EXPECT().ReadDir(gomock.Eq(path))
	err := Ls(m, data, path)
	assert.Nil(t, err)
}

func TestListFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var data LsData
	path := "./"
	m := mock_ioutil.NewMockIoUtil(ctrl)

	afile := mock_os.NewMockFileInfo(ctrl)
	files := []os.FileInfo{afile}
	afile.EXPECT().Name().Return("file_name.txt").AnyTimes()
	m.EXPECT().ReadDir(gomock.Eq(path)).Return(files, nil)
	err := Ls(m, data, path)
	assert.Nil(t, err)
}

func TestIoCallFail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var data LsData
	path := "./"
	m := mock_ioutil.NewMockIoUtil(ctrl)

	errorResult := errors.New("io error")
	m.EXPECT().ReadDir(gomock.Eq(path)).Return(nil, errorResult)
	result := Ls(m, data, path)
	assert.Equal(t, result, errorResult, "they should be equal")
}

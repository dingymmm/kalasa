// MIT License

// Copyright (c) 2022 Leon Ding

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package bigmap

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
)

var (
	perm = os.FileMode(0755)

	dataFileSuffix  = ".bm"
	indexFileSuffix = ".idx"
	hintFileSuffix  = ".hint"

	secret = []byte("1234567890123456")

	FileOnlyReadANDWrite = os.O_RDWR | os.O_APPEND | os.O_CREATE

	ErrPathIsExists         = errors.New("big map error 101: an empty path is illegal")
	ErrPathNotAvailable     = errors.New("big map error 102: the current directory path is unavailable")
	ErrCreateDirectoryFail  = errors.New("big map error 103: failed to create a data store directory")
	ErrCreateActiveFileFail = errors.New("big map error 104: failed to create a writable and readable active file")
	ErrSourceDataEncodeFail = errors.New("big map error 201: source data fails to be encrypted by the encoder")
	ErrSourceDataDecodeFail = errors.New("big map error 202: source data failed to be decrypted by encoder")
	ErrEntityDataBufToFile  = errors.New("big map error 203: error writing entity record data from buffer to file")
)

type ActiveFile struct {
	fid string
	*os.File
}

// Identifier return active file identifier
func (f ActiveFile) Identifier() string {
	return f.fid
}

func createActiveFile(path string) error {
	// 创建文件
	currentFile = new(ActiveFile)
	currentFile.fid = uuid.NewString()
	filePath := dataFilePath(path)

	if file, err := os.OpenFile(filePath, FileOnlyReadANDWrite, perm); err != nil {
		return err
	} else {
		// 可能需要加锁
		currentFile.File = file
	}
	return nil
}

func dataFilePath(path string) string {
	return fmt.Sprintf("%s%s%s", path, currentFile.fid, dataFileSuffix)
}

func indexFilePath(path string) string {
	return fmt.Sprintf("%s%s%s", path, currentFile.fid, indexFileSuffix)
}

func hintFilePath(path string) string {
	return fmt.Sprintf("%s%s%s", path, currentFile.fid, hintFileSuffix)
}

func recoveryIndex() {

}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
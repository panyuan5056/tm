package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"tm/pkg/logging"
	"tm/pkg/setting"

	"os"
	"time"

	"github.com/go-basic/uuid"
)

func ParseFile(ext string) string {
	day := time.Now().Format("20060102")
	path := fmt.Sprintf("%s%s/", setting.UPLOADPATH, day)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			logging.Error(err.Error())
		}
	}
	dst := fmt.Sprintf("%s%s/%s.%s", setting.UPLOADPATH, day, uuid.New(), ext)
	return dst
}

func ReadFile(path string) []byte {
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		logging.Error(err.Error())
	} else {
		if contentByte, err := ioutil.ReadAll(f); err == nil {
			return contentByte
		} else {
			logging.Error(err.Error())
		}
	}
	return bytes.NewBuffer(nil).Bytes()
}

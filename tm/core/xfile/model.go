package xfile

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"strings"
	"tm/pkg/logging"
	"tm/pkg/setting"
)

type UploadForm struct {
	Header string                `form:"header"`
	Upload *multipart.FileHeader `form:"upload"`
	Name   string                `form:"name"`
}

func (u *UploadForm) valid() (string, bool) {
	filenames := strings.Split(u.Upload.Filename, ".")
	if len(filenames) > 1 {
		ext := strings.ToLower(filenames[len(filenames)-1])
		for _, ext2 := range setting.EXT {
			if ext == ext2 {
				return ext, true
			}
		}
	}
	return "", false
}
func (u *UploadForm) str(dst string) (string, bool) {
	c := map[string]string{"name": u.Name, "header": u.Header, "dst": dst}
	b, err := json.Marshal(c)
	if err != nil {
		logging.Error(err.Error())
		return "", false
	}
	return string(b), true
}

func (u *UploadForm) content() []byte {
	content, _ := u.Upload.Open()
	defer content.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, content); err != nil {
		return nil
	}
	return buf.Bytes()
}

type Tables struct {
	Name string
}

type Schemas struct {
	Name string
}

type Datas map[string][]map[string]string

func (d *Datas) Tables() []string {
	tables := []string{}
	for table, _ := range *d {
		tables = append(tables, table)
	}
	return tables
}

func (d *Datas) Schemas() map[string][]string {
	results := map[string][]string{}
	for name, v := range *d {
		keys := []string{}
		for _, m := range v {
			for key, _ := range m {
				keys = append(keys, key)
			}
			break
		}
		results[name] = keys
		break
	}
	return results
}

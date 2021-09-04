package xdb

import (
	"encoding/json"
	"tm/pkg/logging"
)

type Config struct {
	Category string `form:"category" json:"category" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Network  string `form:"network" json:"network" binding:"required"`
	Server   string `form:"server" json:"server" binding:"required"`
	Port     string `form:"Port" json:"port" binding:"required"`
	Database string `form:"database" json:"database" binding:"required"`
	Charset  string `form:"charset" json:"charset" binding:"required"`
}

func (c *Config) str() (string, bool) {
	b, err := json.Marshal(c)
	if err != nil {
		logging.Error(err.Error())
		return "", false
	}
	return string(b), true
}

type Tables struct {
	Name string `db:"Name"`
}

type Schemas struct {
	Name string `db:"Name"`
}

type Datas map[string]interface{}

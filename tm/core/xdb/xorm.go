package xdb

import (
	"github.com/spf13/cast"
	//"database/sql"

	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Xorm struct {
	Db       *sqlx.DB
	Category string
	Username string
	Password string
	Network  string
	Server   string
	Port     string
	Database string
}

func (xm *Xorm) close() {
	xm.Db.Close()
}

func (xm *Xorm) parseTable() (string, bool) {
	switch xm.Category {
	case "mysql":
		return fmt.Sprintf("SELECT table_name as Name FROM information_schema.tables WHERE table_schema='%s'", xm.Database), true
	case "postgres":
		return "SELECT table_name as Name FROM information_schema.tables WHERE table_schema = 'public'", true
	case "kingbase":
		return "SELECT table_name as Name FROM information_schema.tables WHERE table_schema = 'public'", true
	case "sqlite":
		return "SELECT table_name as Name FROM information_schema.tables WHERE table_schema = 'public'", true
	case "oracle":
		return "select table_name from sys.dba_tables where owner='schema名'", true
	case "sqlserver":
		return "SELECT table_name FROM information_schema.tables WHERE table_schema = 'mydatabasename' AND table_type = 'base table' ", true
	default:
		return "SELECT table_name as Name FROM information_schema.tables WHERE table_schema = 'public'", true
	}
}

func (xm *Xorm) parseSchemas(table string) (string, bool) {

	switch xm.Category {
	case "mysql":
		return fmt.Sprintf("SELECT COLUMN_NAME as Name FROM INFORMATION_SCHEMA.COLUMNS where table_schema='%s' AND table_name ='%s'", xm.Database, table), true
	case "postgres":
		return fmt.Sprintf("SELECT column_name as Name FROM INFORMATION_SCHEMA.COLUMNS where table_name ='%s'", table), true
	case "sqlite":
		return "", true
	case "oracle":
		return "", true
	case "sqlserver":
		return "", true
	default:
		return "", false
	}
}

func (xm *Xorm) parseData(table string) (string, bool) {
	switch xm.Category {
	case "mysql":
		return fmt.Sprintf("SELECT * FROM %s", table), true
	case "postgres":
		return fmt.Sprintf("SELECT * FROM %s", table), true
	case "sqlite":
		return fmt.Sprintf("SELECT * FROM %s", table), true
	case "oracle":
		return fmt.Sprintf("SELECT * FROM %s", table), true
	case "sqlserver":
		return fmt.Sprintf("SELECT * FROM %s", table), true
	default:
		return fmt.Sprintf("SELECT * FROM %s", table), true
	}
}

func (xm *Xorm) Tables() []string {
	data := []Tables{}
	results := []string{}
	if xsql, ok := xm.parseTable(); ok {
		if err := xm.Db.Select(&data, xsql); err == nil {
			for _, table := range data {
				results = append(results, table.Name)
			}
		}
	}
	return results
}

func (xm *Xorm) Xschemas(table string) []string {
	data := []Schemas{}
	results := []string{}
	if xsql, ok := xm.parseSchemas(table); ok {
		if err := xm.Db.Select(&data, xsql); err == nil {
			for _, row := range data {
				results = append(results, row.Name)
			}
		}
	}
	return results
}

func (xm *Xorm) Data(table string) []map[string]string {
	results := []map[string]string{}
	if xsql, ok := xm.parseData(table); ok {
		if rows, err := xm.Db.Queryx(xsql); err == nil {
			for rows.Next() {
				//下面演示如何将数据保存到struct、map和数组中
				//定义map类型
				m := map[string]interface{}{}
				//保存到map
				if err = rows.MapScan(m); err != nil {
					fmt.Println(err)
				} else {
					n := map[string]string{}
					for k, v := range m {
						n[k] = cast.ToString(v)
					}
					results = append(results, n)
				}
			}
		}
	}
	return results
}

func Run(config map[string]string) *Xorm {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8", config["username"], config["password"], config["network"], config["server"], config["port"], config["database"])

	DB, err := sqlx.Connect(config["category"], dsn)
	if err != nil {
		fmt.Printf("Open failed,err:%v\n", err)
		return &Xorm{}
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数
	return &Xorm{Db: DB, Category: config["category"], Username: config["username"], Password: config["password"], Network: config["network"], Server: config["server"], Port: config["port"], Database: config["database"]}
}

package db

import (
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/mint-leaf/gin-tpl/utils"
	"gopkg.in/yaml.v2"
)

var (
	db *xorm.Engine
)

const (
	configFile = "db.config.yaml"
)

// Config database config
type Config struct {
	Name  string                 `yaml:"name"`
	DSN   string                 `yaml:"dsn"`
	Extra map[string]interface{} `yaml:"extra"`
}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		panic("[-]get current path error, msg: " + err.Error())
	}
	c, path := new(Config), filepath.Join(pwd, configFile)
	err = utils.LoadDataFromFile(path, yaml.Unmarshal, c)
	if err != nil {
		log.Fatalf("[-]load config from file error, path: %s, err: %s", path, err.Error())
	}
	db, err = xorm.NewEngine(c.Name, c.DSN)
	if err != nil {
		log.Fatalf("[-]open db driver error! msg: %s", err.Error())
	}
	// use `db` log for i/o log, because for me, database info is the most import
	// maybe use gin log as default logger
	// todo: use gin log as default logger and use other log plugin for detail log info
	logger, ok := c.Extra["logfile"]
	if ok {
		if filename, ok := logger.(string); ok {
			file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0777)
			if err == nil {
				xorm.NewSimpleLogger(file)
			}
		}
	}
}

// GetCli get database client
func GetCli() *xorm.Engine {
	return db
}

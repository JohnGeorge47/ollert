package config

import (
	"flag"
	"fmt"
	"golang.org/x/tools/go/analysis/passes/copylock/testdata/src/a"
	"os"
)

type Config struct {
	RedisConfig RedisConf
	MySqlConfig MySqlConf
}

type RedisConf struct {
}

type MySqlConf struct {
	DbUser string
	DbPwd  string
	DbPort string
	DbHost string
	DbName string
}

func Get() *Config {
	config := &Config{}
	flag.StringVar(&config.MySqlConfig.DbUser, "dbuser", os.Getenv("MYSQL_HOST"), "mysql db host")
	flag.StringVar(&config.MySqlConfig.DbPwd, "dbpwd", os.Getenv("MYSQL_PWD"), "mysql db pwd")
	flag.StringVar(&config.MySqlConfig.DbPwd, "dbpwd", os.Getenv("MYSQL_PORT"), "mysql port")
	flag.StringVar(&config.MySqlConfig.DbHost, "dbhost", os.Getenv("MYSQL_HOST"), "mysql host")
	flag.StringVar(&config.MySqlConfig.DbName, "dbname", os.Getenv("MYSQL_DB"), "mysql db")
	flag.Parse()
	return config
}

func (c *Config) Connectionstring() string {
	return fmt.Sprintf("%s:%s@/%s", c.MySqlConfig.DbUser, c.MySqlConfig.DbPwd, c.MySqlConfig.DbPwd)
}

package config

import (
	"fmt"
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

var config Config

func Get() *Config {
	val := os.Getenv("MYSQL_USER")
	fmt.Println(val)
	sqlconf := MySqlConf{
		DbUser: os.Getenv("MYSQL_USER"),
		DbPwd:  os.Getenv("MYSQL_PWD"),
		DbPort: os.Getenv("MYSQL_PORT"),
		DbHost: os.Getenv("MYSQL_HOST"),
		DbName: os.Getenv("MYSQL_DB"),
	}
	config.MySqlConfig = sqlconf
	return &config
}

func (c *Config) Connectionstring() string {
	return fmt.Sprintf("%s:%s@/%s", c.MySqlConfig.DbUser, c.MySqlConfig.DbPwd, c.MySqlConfig.DbName)
}

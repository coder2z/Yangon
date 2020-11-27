/**
* @Author: myxy99 <myxy99@foxmail.com>
* @Date: 2020/11/4 11:18
 */
package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type Options struct {
	Host                  string        `json:"host,omitempty" yaml:"host" description:"db service host address"`
	Username              string        `json:"username,omitempty" yaml:"username"`
	Password              string        `json:"-" yaml:"password"`
	Type                  string        `json:"type" yaml:"type"`
	DBName                string        `json:"dbName" yaml:"dbName"`
	Debug                 bool          `json:"debug" yaml:"debug"`
	Port                  string        `json:"port" yaml:"port"`
	MaxIdleConnections    int           `json:"maxIdleConnections,omitempty" yaml:"maxIdleConnections"`
	MaxOpenConnections    int           `json:"maxOpenConnections,omitempty" yaml:"maxOpenConnections"`
	MaxConnectionLifeTime time.Duration `json:"maxConnectionLifeTime,omitempty" yaml:"maxConnectionLifeTime"`
}

func NewDatabaseOptions() *Options {
	return &Options{
		Host:                  "127.0.0.1",
		Username:              "root",
		Password:              "root",
		DBName:                "reminder",
		Type:                  "mysql",
		Port:                  "3306",
		Debug:                 true,
		MaxIdleConnections:    100,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Duration(10) * time.Second,
	}
}

func (m *Options) GetDSN() string {
	switch m.Type {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&allowNativePasswords=true&parseTime=true", m.Username, m.Password, m.Host, m.Port, m.DBName)
	case "sqlite3":
		return m.DBName
	case "postgres":
		fallthrough
	default:
		return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", m.Host, m.Port, m.Username, m.DBName, m.Password)
	}
}
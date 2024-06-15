package conf

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

// config for the backend
func DefaultConfig() *Config {
	return &Config{
		MySQL: &MySQL{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "myblog",
			Username: "root",
			Password: "87654321",
		},
	}
}

type Config struct {
	MySQL *MySQL `json:"mysql"`
}

type MySQL struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DB       string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`

	lock sync.Mutex
	conn *gorm.DB
}

func (m *MySQL) GetDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.DB)
	return dsn
}

func (m *MySQL) GetConn() *gorm.DB {
	if m.conn == nil {
		// return m.conn
		m.lock.Lock()
		defer m.lock.Unlock()
		conn, err := gorm.Open(mysql.Open(m.GetDSN()), &gorm.Config{})

		if err != nil {
			panic(err)
		}
		m.conn = conn

	} // if db is nil, create a new connection otherwise return the existing connection

	return m.conn
}

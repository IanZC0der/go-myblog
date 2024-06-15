package conf

import "gorm.io/gorm"

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
}

func (m *MySQL) GetConn() *gorm.DB {
	return nil
}

package conf

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"

	"encoding/json"

	"os"
	"strconv"
	// "github.com/joho/godotenv"
)

// config for the backend
func DefaultConfig() *Config {
	return &Config{
		MySQL: &MySQL{
			Host:     "127.0.0.1",
			Port:     3306,
			DB:       "myblog",
			Username: "root",
			Password: "12345678",
		},

		App: &App{
			HttpHost: "127.0.0.1",
			HttpPort: 7080,
		},
	}
}

func ConfigFromEnv() *Config {
	// pwd, _ := os.Getwd()
	// fmt.Println(pwd)
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	dbPortNumber, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	httpPortNumber, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	fmt.Println(os.Getenv("MYSQL_HOST"))
	fmt.Println(httpPortNumber)
	return &Config{
		MySQL: &MySQL{
			Host:     os.Getenv("MYSQL_HOST"),
			Port:     dbPortNumber,
			DB:       os.Getenv("MYSQL_DB"),
			Username: os.Getenv("MYSQL_USERNAME"),
			Password: os.Getenv("MYSQL_PASSWORD"),
		},
		App: &App{
			HttpHost: os.Getenv("HTTP_HOST"),
			HttpPort: int64(httpPortNumber),
		},
	}
}

type Config struct {
	MySQL *MySQL `json:"mysql"`

	App *App `json:"app"`
}

type App struct {
	HttpHost string `json:"http_host" env:"HTTP_HOST"`
	HttpPort int64  `json:"http_port" env:"HTTP_PORT"`
}

func (a *App) HttpAddress() string {
	return fmt.Sprintf("%s:%d", a.HttpHost, a.HttpPort)
}

func (c *Config) String() string {

	jsonConfig, _ := json.Marshal(c)
	return string(jsonConfig)
}

type MySQL struct {
	Host     string `json:"host" toml:"host" env:"MYSQL_HOST"`
	Port     int    `json:"port" toml:"port" env:"MYSQL_PORT"`
	DB       string `json:"database" toml:"database" env:"MYSQL_DB"`
	Username string `json:"username" toml:"username" env:"MYSQL_USERNAME"`
	Password string `json:"password" toml:"password" env:"MYSQL_PASSWORD"`

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

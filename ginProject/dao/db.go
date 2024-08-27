package dao

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

var db *gorm.DB

// Config 定义了数据库的配置结构
type Config struct {
	Database struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		User      string `yaml:"user"`
		Password  string `yaml:"password"`
		Dbname    string `yaml:"dbname"`
		Charset   string `yaml:"charset"`
		ParseTime bool   `yaml:"parseTime"`
		Loc       string `yaml:"loc"`
	} `yaml:"database"`
}

// 读取配置文件
func readConfig() (*Config, error) {
	var config Config
	data, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// 初始化数据库连接
func initDB() {
	config, err := readConfig()
	if err != nil {
		log.Fatal("Failed to read config file:", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Dbname,
		config.Database.Charset,
		config.Database.ParseTime,
		config.Database.Loc,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	//db.AutoMigrate(&User{})
}

package dao

import (
	"context"
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var (
	db  *gorm.DB
	rdb *redis.Client
	ctx = context.Background()
)

// Config 定义了数据库和 Redis 的配置结构
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
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
}

// 读取配置文件
func readConfig() (*Config, error) {
	var config Config

	// 从环境变量中获取配置文件路径，如果未设置，则使用默认路径
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./config/config.yaml"
	}

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}
	return &config, nil
}

// 初始化MySQL数据库连接
func initDBMySQL() {
	config, err := readConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
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
		log.Fatalf("Error connecting to database: %v", err)
	}
	log.Println("Successfully connected to MySQL database.")

	// 自动迁移表
	// db.AutoMigrate(&User{})
}

// 初始化Redis连接
func initRedis() {
	config, err := readConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}
	log.Println("Successfully connected to Redis.")
}

// GetDB 返回GORM的数据库连接
func GetDB() *gorm.DB {
	if db == nil {
		initDBMySQL()
	}
	return db
}

// GetRedis 返回Redis的客户端连接
func GetRedis() *redis.Client {
	if rdb == nil {
		initRedis()
	}
	return rdb
}

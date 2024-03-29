package db

import (
	"fmt"
	"github.com/tsopia/gokit/conf"
	"github.com/tsopia/gokit/xerrors"
	"gorm.io/gorm/schema"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLConfig represents MySQL configuration parameters
type MySQLConfig struct {
	Host              string
	Port              int
	Username          string
	Password          string
	Database          string
	MaxIdleConn       int
	MaxOpenConn       int
	IdleTimeout       time.Duration
	MaxLifetime       time.Duration
	HealthCheckPeriod time.Duration
	MaxIdleClosed     bool
}

// Client represents a MySQL client
type Client struct {
	db *gorm.DB
}

// 定义一个全局的Client变量，以便在整个程序中使用
var DbClient *Client

// 初始化DbClient 通过conf 下的MysqlConf 获取相关配置，复制给MySQLConfig
func InitDbClient() *xerrors.Error {
	config := MySQLConfig{
		Host:              *conf.DefaultConf.Mysql.Host,
		Port:              *conf.DefaultConf.Mysql.Port,
		Username:          *conf.DefaultConf.Mysql.Username,
		Password:          *conf.DefaultConf.Mysql.Password,
		Database:          *conf.DefaultConf.Mysql.Database,
		MaxIdleConn:       *conf.DefaultConf.Mysql.MaxIdleConn,
		MaxOpenConn:       *conf.DefaultConf.Mysql.MaxOpenConn,
		IdleTimeout:       *conf.DefaultConf.Mysql.IdleTimeout,
		MaxLifetime:       *conf.DefaultConf.Mysql.MaxLifetime,
		HealthCheckPeriod: *conf.DefaultConf.Mysql.HealthCheckPeriod,
	}
	client, err := NewClient(config)
	if err != nil {
		return err
	}
	DbClient = client
	return nil
}

// NewClient creates a new MySQL client
func NewClient(config MySQLConfig) (*Client, *xerrors.Error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username, config.Password, config.Host, config.Port, config.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		// 在这里设置你的自定义命名策略
		SingularTable: true, // 禁用复数形式转换
	}})
	if err != nil {
		return nil, xerrors.Wrap(err, xerrors.ErrDatabaseConnectionFailed, "failed to connect to MySQL database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, xerrors.Wrap(err, xerrors.ErrDatabaseConnectionFailed, "failed to get database connection")
	}

	sqlDB.SetMaxIdleConns(config.MaxIdleConn)
	sqlDB.SetMaxOpenConns(config.MaxOpenConn)
	sqlDB.SetConnMaxIdleTime(config.IdleTimeout * time.Second)
	sqlDB.SetConnMaxLifetime(config.MaxLifetime * time.Second)
	sqlDB.SetConnMaxIdleTime(config.HealthCheckPeriod * time.Second)

	return &Client{db: db}, nil
}

// Close closes the MySQL client connection
func (c *Client) Close() *xerrors.Error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return xerrors.Wrap(err, xerrors.ErrDatabaseConnectionFailed, "failed to get database connection")
	}

	if err := sqlDB.Close(); err != nil {
		return xerrors.Wrap(err, xerrors.ErrDatabaseConnectionFailed, "failed to close MySQL database connection")
	}

	return nil
}

// Ping pings the MySQL database
func (c *Client) Ping() *xerrors.Error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return xerrors.Wrap(err, xerrors.ErrDatabaseConnectionFailed, "failed to get database connection")
	}

	if err := sqlDB.Ping(); err != nil {
		return xerrors.Wrap(err, xerrors.ErrDatabaseConnectionFailed, "failed to ping MySQL database")
	}

	return nil
}

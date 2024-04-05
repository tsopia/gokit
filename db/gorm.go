package db

import (
	"fmt"
	"github.com/tsopia/gokit/model"
	"github.com/tsopia/gokit/xerrors"
	"gorm.io/gorm/schema"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Client represents a MySQL client
type Client struct {
	db *gorm.DB
}

// DbClient 定义一个全局的Client变量，以便在整个程序中使用
var DbClient *Client

// InitDbClient 初始化DbClient 通过conf 下的MysqlConf 获取相关配置，复制给MySQLConfig
func InitDbClient(conf model.MysqlConf) error {

	client, err := NewClient(conf)
	if err != nil {
		return err
	}
	DbClient = client

	return nil
}

// NewClient creates a new MySQL client
func NewClient(config model.MysqlConf) (*Client, error) {
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
func (c *Client) DB() *gorm.DB {
	return c.db
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

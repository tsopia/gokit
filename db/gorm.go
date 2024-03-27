package db

import (
	"fmt"
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
	IdleTimeout       int
	MaxLifetime       int
	HealthCheckPeriod int
	MaxIdleClosed     bool
}

// Client represents a MySQL client
type Client struct {
	db *gorm.DB
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
	sqlDB.SetConnMaxIdleTime(time.Duration(config.IdleTimeout) * time.Second)
	sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifetime) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.HealthCheckPeriod) * time.Second)

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

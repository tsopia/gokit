package db

import (
	"github.com/stretchr/testify/assert"
	"github.com/tsopia/gokit/model"
	"github.com/tsopia/gokit/xerrors"
	"testing"
	"time"
)

func TestMySQLClient(t *testing.T) {
	// 设置 MySQL 配置信息
	config := model.MysqlConf{
		Host:              "localhost",
		Port:              3306,
		Username:          "root",
		Password:          "123456",
		Database:          "pdrland",
		MaxIdleConn:       10,
		MaxOpenConn:       100,
		IdleTimeout:       30,
		MaxLifetime:       0,
		HealthCheckPeriod: 10,
		MaxIdleClosed:     true,
	}

	// 创建 MySQL 客户端
	err := InitDbClient(config)
	defer func() {
		if err := DbClient.Close(); err != nil {
			t.Errorf("Failed to close MySQL client: %v", err)
		}
	}()

	// 检查是否创建 MySQL 客户端时发生错误
	assert.Nil(t, err)

	// 测试 MySQL 客户端的 Ping 方法
	err = DbClient.Ping()
	assert.Nil(t, err)

	// 在此添加其他需要测试的操作，例如执行查询、插入等
	// 测试创建表、插入数据和删除表
	err = testCreateInsertDelete(DbClient)
	assert.Nil(t, err)
	// 休眠一段时间，以便在关闭连接之前观察连接池的行为
	time.Sleep(5 * time.Second)
}
func testCreateInsertDelete(client *Client) *xerrors.Error {
	// 创建表如果存在就先删除
	client.db.Exec("DROP TABLE IF EXISTS test_table")

	if err := client.db.Exec("CREATE TABLE IF NOT EXISTS test_table (id INT PRIMARY KEY, name VARCHAR(100))").Error; err != nil {
		return xerrors.Wrap(err, xerrors.ErrDatabaseQueryFailed, "failed to create table")
	}

	// 插入数据
	if err := client.db.Exec("INSERT INTO test_table (id, name) VALUES (?, ?), (?, ?)", 1, "John", 2, "Doe").Error; err != nil {
		return xerrors.Wrap(err, xerrors.ErrDatabaseQueryFailed, "failed to insert data into table")
	}

	// 从表中查询数据
	var count int64
	if err := client.db.Model(&TestTable{}).Count(&count).Error; err != nil {
		return xerrors.Wrap(err, xerrors.ErrDatabaseQueryFailed, "failed to count rows in table")
	}

	//删除表
	if err := client.db.Exec("DROP TABLE IF EXISTS test_table").Error; err != nil {
		return xerrors.Wrap(err, xerrors.ErrDatabaseQueryFailed, "failed to drop table")
	}

	return nil
}

type TestTable struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

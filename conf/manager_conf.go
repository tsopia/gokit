package conf

type ManagerConf struct {
	Mysql          *MysqlConf
	LogLevel       *LogLevelConf
	Port           *int
	GrpcPort       *int
	GrpcServerName *string
	SeverName      *string
}
type LogLevelConf struct {
	GoResetDebug *bool
	LogLevel     *string
}
type MysqlConf struct {
	Host              *string
	Port              *string
	Username          *string
	Password          *string
	Database          *string
	MaxIdleConn       *int  // 最大空闲连接数
	MaxOpenConn       *int  // 最大打开连接数
	IdleTimeout       *int  // 空闲连接超时时间（秒）
	MaxLifetime       *int  // 连接的最大生命周期，0表示无限制
	HealthCheckPeriod *int  // 连接健康检查周期（秒）
	MaxIdleClosed     *bool // 是否关闭最大生命周期超过MaxLifetime的空闲连接
}

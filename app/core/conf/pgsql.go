package conf

type PGSqlConf struct {
	Port     string
	Host     string
	Database string
	Username string
	Password string
}

type PGSqlLog struct {
	Enable bool
	Level  string
	Format string
	Type   string
	Path   string
}

type PGSql struct {
	Log   PGSqlLog
	Write PGSqlConf
	Read  PGSqlConf
}

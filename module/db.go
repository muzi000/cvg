package module

type DatabaseSetting struct {
	DBType   string
	UserName string
	Password string
	Host     string
	Port     int
	DBName   string
	Charset  string
}

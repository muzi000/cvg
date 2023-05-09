package module

type LdapSetting struct {
	Method   string
	Host     string
	Port     int
	Base     string
	UserName string
	Password string
}

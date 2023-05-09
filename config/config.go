package config

import (
	"github.com/muzi000/vuln-go/module"
	"github.com/spf13/viper"
)

var (
	ServerSetting   module.ServerSetting
	DatabaseSetting module.DatabaseSetting
	LdapSetting     module.LdapSetting
)

func InitSetting() error {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("config/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return err
	}
	err = vp.UnmarshalKey("Server", &ServerSetting)
	if err != nil {
		return err
	}
	err = vp.UnmarshalKey("Database", &DatabaseSetting)
	if err != nil {
		return err
	}
	err = vp.UnmarshalKey("Ldap", &LdapSetting)
	if err != nil {
		return err
	}
	return nil
}

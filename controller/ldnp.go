package controller

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
	"github.com/muzi000/vuln-go/config"
)

var LdapCon *ldap.Conn

func InitLdap() error {
	var err error
	LdapCon, err = ldap.Dial(config.LdapSetting.Method, fmt.Sprintf("%s:%d", config.LdapSetting.Host, config.LdapSetting.Port))
	if err != nil {
		return err
	}
	err = LdapCon.Bind(fmt.Sprintf("%s,%s", config.LdapSetting.UserName, config.LdapSetting.Base), config.LdapSetting.Password)
	if err != nil {
		return err
	}
	return nil
}

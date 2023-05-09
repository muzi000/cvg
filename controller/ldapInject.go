package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ldap/ldap/v3"
	"github.com/muzi000/vuln-go/logging"
)

func LdapVuln1(ctx *gin.Context) {
	username := ctx.Query("user")

	searchRequest := ldap.NewSearchRequest("dc=libaigo,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", ldap.EscapeFilter(username)),
		[]string{"dn"},
		nil)

	sr, err := LdapCon.Search(searchRequest)
	if err != nil {
		logging.LogPrint("err", err.Error())
		ctx.String(http.StatusBadRequest, "search err")
	}
	for _, entry := range sr.Entries {
		//fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
		ctx.Writer.WriteString(fmt.Sprintf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn")))
	}

}

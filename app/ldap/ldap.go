package ldap

import (
	"fmt"

	ldapAPI "gopkg.in/ldap.v2"
)

type ADConn struct {
	conn     *ldapAPI.Conn
	number   string
	userIsIn bool
	mail     string
}

func NewLdapConnection(number string) *ADConn {

	ldapConn, err := ldapAPI.Dial("tcp", fmt.Sprintf("%s:%d", "0.0.0.0", 389))

	defer func() {
		if panicValue := recover(); panicValue != nil {
			fmt.Printf("LDAP connection faild, err - %s", err)
		}
	}()

	ldapConn.Debug = true

	ldapConn.Bind("", "")

	return &ADConn{
		conn:     ldapConn,
		userIsIn: false,
		number:   number,
	}

}

func (ad *ADConn) userSearchRequest(number string) {

	serchReq := ldapAPI.SearchRequest{
		BaseDN:       "dc=test,dc=local",
		Scope:        ldapAPI.FilterSubstringsFinal,
		DerefAliases: ldapAPI.NeverDerefAliases,
		SizeLimit:    0,
		TimeLimit:    0,
		TypesOnly:    false,
		Filter:       fmt.Sprintf("(&(objectClass=user)(telephoneNumber=%s))", number),
		Attributes:   []string{"mail"},
		Controls:     nil,
	}

	ans, err := ad.conn.Search(&serchReq)

	defer func() {
		if panicValue := recover(); panicValue != nil {
			fmt.Errorf("Error to get user %s", err)
		}
	}()

	if ans.Entries[0].DN != "" {
		ad.userIsIn = true
		ad.mail = ans.Entries[0].Attributes[0].Values[0]
	}
}

func (ad *ADConn) GetAuthResult() bool {
	ad.userSearchRequest(ad.number)
	return ad.userIsIn
}

func (ad *ADConn) GetMail() string {
	return ad.mail
}

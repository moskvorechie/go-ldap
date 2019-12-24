package ldap

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/pkg/errors"
)

var (
	Host   string
	Port   string
	User   string
	Pass   string
	BaseDN string
	Object string
	l      *ldap.Conn
)

func GetUser(username string, password string) (user ldap.Entry, err error) {

	if Host == "" || Port == "" || User == "" || Pass == "" || BaseDN == "" {
		err = errors.New("Not all LDAP var exist")
		return
	}

	// Connect
	l, err = ldap.Dial("tcp", fmt.Sprintf("%s:%s", Host, Port))
	if err != nil {
		err = errors.Wrap(err, "Error connect to LDAP")
		return
	}
	defer l.Close()

	// Bind admin
	err = l.Bind(User, Pass)
	if err != nil {
		err = errors.Wrap(err, "Error auth to LDAP")
		return
	}

	// Search
	searchRequest := ldap.NewSearchRequest(
		BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(Object, username),
		[]string{},
		nil)
	sr, err := l.Search(searchRequest)
	if err != nil {
		err = errors.Wrap(err, "Error search in LDAP")
		return
	}

	// Check count
	if len(sr.Entries) != 1 {
		err = errors.Wrap(err, "Error count search result LDAP")
		return
	}

	// Bind as the user to verify their password
	user = *sr.Entries[0]
	err = l.Bind(user.DN, password)
	if err != nil {
		err = errors.Wrap(err, "Error match user in LDAP")
		return
	}

	return
}

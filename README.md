# Module for auth by LDAP

Usage example
```golang

// Init module
ldap.Host = "10.0.0.0"
ldap.Port = "1234"
ldap.User = "MyUser"
ldap.Pass = "MyPass"
ldap.BaseDN = "dc=company,dc=local"
ldap.Object = "(&(objectClass=Person)(sAMAccountName=%s))"

// Check in LDAP
user, err := ldap.GetUser(login, password)
if err != nil {
    panic(err)
}
if user.DN == "" {
    logs.Info("Auth fail")
    return
}

// Check user have group access "MyGroup"
groupAccess := false
for _, v := range user.Attributes {
    if v.Name == "memberOf" && len(v.Values) > 0 {
        for _, vv := range v.Values {
            if strings.Contains(vv, "MyGroup") {
                groupAccess = true
                break
            }
        }
    }
    if groupAccess {
        break
    }
}
if !groupAccess {
    logs.Info("Auth fail")
    return
}

logs.Info("Auth success")
return
```


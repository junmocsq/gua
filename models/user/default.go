package user

import "github.com/junmocsq/gua/models/dbcache"

func init() {
	dbcache.NewDb().DB().AutoMigrate(&User{}, &StdAddress{}, &Address{})
}

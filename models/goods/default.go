package goods

import "github.com/junmocsq/gua/models/dbcache"

func init() {
	dbcache.NewDb().DB().AutoMigrate(&Brand{}, &Category{}, &Goods{}, &GoodsTags{}, &Tag{})
}

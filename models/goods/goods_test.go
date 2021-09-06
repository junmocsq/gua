package goods

import (
	"github.com/junmocsq/gua/models/dbcache"
	"testing"
)

func TestGoods_Add(t *testing.T) {
	dbcache.NewDb().DB().AutoMigrate(  &Goods{}, &GoodsTags{})
	g := NewGoods()
	goods := Goods{
		Name:       "计算机程序设计艺术",
		Points:     22,
		State:      1,
		Amount:     139.8,
		Quantity:   100,
		Spec:       "高德纳封神之作",
		BrandId:    1,
		Category:   1,
		MainImg:    "",
		ScrollImgs: "",
		IntroImgs:  "",
		Tags: []GoodsTags{{TagId: 1},{TagId: 3},{TagId: 5},{TagId: 7}},
	}
	g.Add(&goods)

	t.Log(g.One(1))
}

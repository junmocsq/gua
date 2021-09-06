package goods

import (
	"github.com/junmocsq/gua/models/dbcache"
	"github.com/sirupsen/logrus"
)

// Category 分类
type Category struct {
	Id        int    `gorm:"primaryKey;type:int;NOT NULL"`
	Pid       int    `gorm:"type:int;NOT NULL;comment:父级id"`
	Name      string `gorm:"size:50"`
	State     uint8  `gorm:"type:tinyint unsigned;NOT NULL;default 1;comment:1 正常 2 下架"`
	SortIndex int    `gorm:"type:int;NOT NULL;default:0"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

func (Category) TableName() string {
	return "g_category"
}

func (c *category) Tag() string {
	return "goods_category"
}

type category struct {
}

func NewCategory() *category {
	return &category{}
}

func (c *category) Add(ca *Category) int64 {
	if ca.Pid > 0 {
		parent := c.One(ca.Pid)
		if parent == nil || parent.State != 1 {
			return 0
		}
	}
	db := dbcache.NewDb()
	stmt := db.DryRun().Create(ca).Statement
	n, err := db.SetTag(c.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(ca)
	if err != nil {
		logrus.Errorf("category err:%s", err)
		n = 0
	}
	return n
}

func (c *category) One(Id int) *Category {
	db := dbcache.NewDb()
	var ca Category
	stmt := db.DryRun().Find(&ca, Id).Statement
	err := db.SetTag(c.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&ca)
	if err != nil {
		logrus.Errorf("category err:%s", err)
	}
	if ca.Id == 0 {
		return nil
	}
	return &ca
}

func (c *category) All() []*Category {
	db := dbcache.NewDb()
	var ca []*Category
	stmt := db.DryRun().Find(&ca).Statement
	err := db.SetTag(c.Tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Fetch(&ca)
	if err != nil {
		logrus.Errorf("category err:%s", err)
	}
	return ca
}

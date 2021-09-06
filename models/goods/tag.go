package goods

import "github.com/junmocsq/gua/models/dbcache"

type Tag struct {
	Id        int    `gorm:"primaryKey;type:int;NOT NULL"`
	Name      string `gorm:"size:50;NOT NULL;default:''"`
	State     uint8  `gorm:"type:tinyint unsigned;NOT NULL;default 1;comment:1 正常 2 下架"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

func (Tag) TableName() string {
	return "g_tag"
}

type tag struct {
}

func NewTag() *tag {
	return &tag{}
}
func (t *tag) tag() string {
	return "goods_tag"
}
func (t *tag) Add(tag string) int {
	db := dbcache.NewDb()
	var tt = Tag{
		Name:  tag,
		State: 1,
	}
	stmt := db.DryRun().Create(&tt).Statement
	_, err := db.SetTag(t.tag()).PrepareSql(stmt.SQL.String(), stmt.Vars...).Create(&tt)
	if err != nil {
		return 0
	}
	return tt.Id
}

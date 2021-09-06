package goods

import (
	"fmt"
	"github.com/junmocsq/gua/models/dbcache"
)

type Follow struct {
	Uid       int   `gorm:"primaryKey;type:int;NOT NULL"`
	Gid       int   `gorm:"primaryKey;type:int;NOT NULL"`
	CreatedAt int64 `gorm:"autoCreateTime;NOT NULL"`
}

func (Follow) TableName() string {
	return "g_follow"
}

type follow struct {
}

func NewFollow() *follow {
	return &follow{}
}

func (f *follow) tag(uid int) string {
	return fmt.Sprintf("goods_follow_%d", uid)
}

func (f *follow) Follow(uid, gid int) bool {
	if f.Check(uid, gid) {
		return true
	}
	fo := Follow{
		Uid: uid, Gid: gid,
	}
	db := dbcache.NewDb()
	res := db.DryRun().Create(&fo).Statement
	n, err := db.SetTag(f.tag(uid)).PrepareSql(res.SQL.String(), res.Vars...).EXEC()
	return err == nil && n > 0
}

// Unfollow 不存在关系时，删除也返回true
func (f *follow) Unfollow(uid, gid int) bool {
	db := dbcache.NewDb()
	res := db.DryRun().Where("uid=? AND gid=?", uid, gid).Delete(&Follow{}).Statement
	_, err := db.SetTag(f.tag(uid)).PrepareSql(res.SQL.String(), res.Vars...).EXEC()
	return err == nil
}

func (f *follow) Check(uid, gid int) bool {
	db := dbcache.NewDb()
	var fo = Follow{
		Uid: uid, Gid: gid,
	}
	res := db.DryRun().Take(&fo).Statement
	err := db.SetTag(f.tag(uid)).PrepareSql(res.SQL.String(), res.Vars...).Fetch(&fo)
	return err == nil && fo.CreatedAt > 0
}

func (f *follow) UserFollowings(uid, page, num int) []int {
	db := dbcache.NewDb()
	var result []int
	res := db.DryRun().Model(&Follow{}).Select("gid").Where("uid=?", uid).Limit(num).Offset(page*num - num).
		Order("created_at desc").Find(&result).Statement
	err := db.SetTag(f.tag(uid)).PrepareSql(res.SQL.String(), res.Vars...).Fetch(&result)
	if err != nil {
		return nil
	}
	return result
}

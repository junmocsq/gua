package goods

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCategory_Add(t *testing.T) {
	t.Skip()
	var arr = make(map[string]map[string][]string)
	arr["图书"] = map[string][]string{
		"文艺":   {"小说", "文学", "青春", "传记", "艺术"},
		"少儿":   {"0-2岁", "3-6岁", "7-10岁", "11-14岁"},
		"人文":   {"历史", "哲学", "心理学", "文化政治"},
		"经管励志": {"经济", "管理", "成功"},
		"科学":   {"计算机", "工业技术", "医学", "科学与自然"},
	}
	arr["服饰"] = map[string][]string{
		"女装": {"衬衫", "卫衣", "牛仔", "大衣", "外套"},
		"男装": {"衬衫", "卫衣", "牛仔", "大衣", "外套"},
		"配件": {"眼镜", "帽子", "围巾", "丝巾"},
		"鞋":  {"跑步鞋", "高跟鞋", "平底鞋"},
		"包":  {"背包", "挎包", "钱包"},
	}
	c := NewCategory()
	for k, v := range arr {
		ca := Category{
			Pid:   0,
			Name:  k,
			State: 1,
		}
		c.Add(&ca)
		t.Log(ca)
		for _k, _v := range v {
			ca1 := Category{
				Pid:   ca.Id,
				Name:  _k,
				State: 1,
			}
			c.Add(&ca1)
			for _, __v := range _v {
				ca2 := Category{
					Pid:   ca1.Id,
					Name:  __v,
					State: 1,
				}
				c.Add(&ca2)
			}
		}
	}
}

func TestCategory(t *testing.T) {
	c := NewCategory()
	c.One(1)
	Convey("category select", t, func() {
		Convey("one", func() {
			So(c.One(1), ShouldNotBeNil)
		})
		Convey("all", func() {
			So(len(c.All()), ShouldBeGreaterThan, 0)
		})
	})
}

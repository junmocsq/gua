package goods

import "github.com/junmocsq/gua/models/goods"

type category struct {
}

func NewCategory() *category {
	return &category{}
}
func (c *category) Map() map[int]*goods.Category {
	list := goods.NewCategory().All()
	categoryMap := make(map[int]*goods.Category)
	for _, v := range list {
		categoryMap[v.Id] = v
	}
	return categoryMap
}

func (c *category) ListMap(ids ...int) (result map[int][]struct {
	Id   int
	Name string
}) {
	all := c.Map()
	for _, id := range ids {
		var tmp []struct {
			Id   int
			Name string
		}
		tmpId := id
		for {
			temp, ok := all[tmpId]
			if ok {
				tmp = append(tmp, struct {
					Id   int
					Name string
				}{
					temp.Id, temp.Name,
				})
			} else {
				break
			}
			if temp.Pid == 0 {
				break
			}
			tmpId = temp.Pid
		}
		result[id] = tmp
	}
	return
}

package goods

// GoodsTags 商品标签关联表
type GoodsTags struct {
	Id      int `gorm:"primaryKey;type:int;NOT NULL"`
	TagId   int `gorm:"type:int;uniqueIndex:t_g_idx"`
	GoodsId int `gorm:"type:int;uniqueIndex:t_g_idx;index:goods_idx"`
}

func (GoodsTags) TableName() string {
	return "g_goods_tags"
}

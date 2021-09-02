package goods

// Category 分类
type Category struct {
	Id        int    `gorm:"primaryKey;type:int;NOT NULL"`
	Name      string `gorm:"size:50"`
	State     uint8  `gorm:"type:tinyint unsigned;NOT NULL;default 1;comment:1 正常 2 下架"`
	SortIndex int    `gorm:"type:int;NOT NULL;default:0"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

func (Category) TableName() string {
	return "g_category"
}

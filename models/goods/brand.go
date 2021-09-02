package goods

// Brand 品牌
type Brand struct {
	Id        int    `gorm:"primaryKey;type:int;NOT NULL"`
	Name      string `gorm:"size:50;NOT NULL;default:''"`
	Icon      string `gorm:"size:50;NOT NULL;default:'';comment:品牌商标"`
	Intro     string `gorm:"size:500;NOT NULL;default:'';comment:品牌介绍"`
	State     uint8  `gorm:"type:tinyint unsigned;NOT NULL;default 1;comment:1 正常 2 下架"`
	SortIndex int    `gorm:"type:int;default:0"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

func (Brand) TableName() string {
	return "g_brand"
}

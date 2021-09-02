package goods

type Tag struct {
	Id        int    `gorm:"primaryKey;type:int;NOT NULL"`
	Name      string `gorm:"size:50;NOT NULL;default:''"`
	State     uint8  `gorm:"type:tinyint unsigned;NOT NULL;default 1;comment:1 正常 2 下架"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}

func (Tag) TableName() string {
	return "g_tag"
}

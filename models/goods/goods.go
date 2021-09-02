package goods

// Goods 商品
type Goods struct {
	Id         int     `gorm:"primaryKey;type:int;NOT NULL"`
	Name       string  `gorm:"size:200;NOT NULL;default:'';comment:商品积分"`
	Points     int     `gorm:"type:int;NOT NULL;default:0;comment:商品积分"`
	State      uint8   `gorm:"type:tinyint;NOT NULL;default 1;comment:状态：1 正常 2 下架 3 屏蔽"`
	Amount     float32 `gorm:"type:decimal(12,2);NOT NULL;default:0;comment:售价"`
	Quantity   int     `gorm:"type:int;NOT NULL;default:0;comment:库存"`
	Spec       string  `gorm:"size:1000;NOT NULL;default:'';comment:规格"`
	BrandId    int     `gorm:"type:int;NOT NULL;default:0;comment:品牌"`
	Category   int     `gorm:"type:int;NOT NULL;default:0;comment:分类"`
	Tags       string  `gorm:"size:100;NOT NULL;default:'';comment:标签,冗余记录"`
	MainImg    string  `gorm:"size:100;NOT NULL;default:'';comment:商品主图"`
	ScrollImgs string  `gorm:"size:1000;NOT NULL;default:'';comment:商品滚动图，多图"`
	IntroImgs  string  `gorm:"size:1000;NOT NULL;default:'';comment:商品简介图，多图"`
	UpdatedAt  int64   `gorm:"autoUpdateTime;NOT NULL;default:0"`
	CreatedAt  int64   `gorm:"autoCreateTime"`
}

func (Goods) TableName() string {
	return "g_goods"
}

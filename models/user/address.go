package user

type Address struct {
	Id         int        `gorm:"type:int;NOT NULL"`
	Uid        int        `gorm:"index;type:int;NOT NULL;default:0"`
	Username   string     `gorm:"size:50;NOT NULL;default:'';comment:收件人"`
	Phone      string     `gorm:"size:50;NOT NULL;default:'';comment:收件人电话"`
	IsDefault  uint8      `gorm:"type:tinyint unsigned;NOT NULL;default:0;comment:是否设置为默认地址"`
	AddressId  int        `gorm:"type:int;NOT NULL;default:0;comment:标准地址id"`
	Detail     string     `gorm:"size:50;NOT NULL;default:'';comment:详细地址"`
	CreatedAt  int64      `gorm:"autoCreateTime;NOT NULL;default:0"`
	StdAddress StdAddress `gorm:"foreignKey:AddressId;references:Id"`
}

func (Address) TableName() string {
	return "u_address"
}

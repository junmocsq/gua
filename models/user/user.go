package user

type User struct {
	Id         int    `gorm:"primaryKey;type:int;NOT NULL"`
	Nickname   string `gorm:"uniqueIndex;size:50;NOT NULL;default:''"`
	Avatar     string `gorm:"size:50;NOT NULL;default:''"`
	Passwd     string `gorm:"size:40;NOT NULL;default:''"`
	Email      string `gorm:"size:50;NOT NULL;default:''"`
	Phone      int64  `gorm:"type:int;NOT NULL;default:0"`
	NationCode int    `gorm:"type:int;NOT NULL;default:86"`
	Salt       string `gorm:"type:char;size:6;NOT NULL;default:''"`
	Gender     uint8  `gorm:"type:tinyint unsigned;NOT NULL;default:0;comment:0 unknown 1 male 2 female"`
	WhatIsUp   string `gorm:"size:100;NOT NULL;default:''"`
	LoginTime  int64  `gorm:"NOT NULL;default:0"`
	UpdatedAt  int64  `gorm:"autoUpdateTime;NOT NULL;default:0"`
	CreatedAt  int64  `gorm:"autoCreateTime;NOT NULL;default:0"`
}

func (User) TableName() string {
	return "u_user"
}

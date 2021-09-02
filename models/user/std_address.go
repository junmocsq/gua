package user

type StdAddress struct {
	Id       int    `gorm:"type:int;NOT NULL"`
	Nation   string `gorm:"size:50;NOT NULL;default:'';comment:国家"`
	Province string `gorm:"size:50;NOT NULL;default:'';comment:省"`
	City     string `gorm:"size:50;NOT NULL;default:'';comment:市"`
	Area     string `gorm:"size:50;NOT NULL;default:'';comment:区"`
}

func (StdAddress) TableName() string {
	return "u_std_address"
}

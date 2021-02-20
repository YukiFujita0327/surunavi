package db

type UserInfo struct {
	BaseModel					`gorm:"embedded"`
	Password	string			`gorm:"column:PASSWORD"`
	Name		string			`gorm:"column:NAME"`
}

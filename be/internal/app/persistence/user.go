package persistence

type User struct {
	UserId   int64 `gorm:"primaryKey"`
	Username string
	Password string
}

func (User) TableName() string {
	return "user"
}

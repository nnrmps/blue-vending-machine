package persistence

type User struct {
	UserId   int64 `gorm:"primaryKey"`
	Username string
	Password string
}

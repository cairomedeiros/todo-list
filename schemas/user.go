package schemas

type User struct {
	Id       int    `gorm:"type:int;primary_key"`
	Name     string `gorm:"type:varchar(255)"`
	Email    string `gorm:"unique"`
	Password string `gorm:"type:varchar(255)"`
}

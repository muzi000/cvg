package module

type User struct {
	Id   uint   `db:"id" json:"id" form:"id"`
	Name string ` db:"name" json:"name" form:"name"`
	//Passwd string `gorm:"column:passwd" db:"passwd" json:"passwd" form:"passwd"`
	Role int `db:"role" json:"role" form:"role"`
}

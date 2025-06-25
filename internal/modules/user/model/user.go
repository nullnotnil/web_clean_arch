package model

type User struct {
	ID       uint   `gorm:"primarykey" json:"userid"`
	Fullname string `gorm:"column:fullname;" json:"full_name"`
	Email    string `gorm:"column:email;" json:""`
}

func (u *User) TableName() string { return "user" }

/* /////// */
type UserCreating struct {
	Fullname string `gorm:"column:fullname;" json:"full_name"`
	Email    string `gorm:"column:email;" json:""`
}

func (u *UserCreating) TableName() string { return "user" }

package user

import (
	"gitea.ysicing.net/cloud/pangu/internal/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"index;unique;"`
	Password string `json:"-"`
	Mail     string `gorm:"unique"`
	Nickname string `json:",omitempty"`
	Token    string `json:"-"`
}

func (User) TableName() string {
	return "user"
}

func init() {
	db.Migrate(User{})
}

func Create(o *User) error {
	return db.DB.Create(o).Error
}

func Update(o *User) error {
	return db.DB.Updates(o).Error
}

func Get(where string, args ...interface{}) (*User, error) {
	var u User
	err := db.DB.Model(User{}).Where(where, args...).Last(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func Gets(where string, args ...interface{}) ([]User, error) {
	var u []User
	err := db.DB.Model(User{}).Where(where, args...).Find(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func FindAll() (o []User) {
	if db.DB.Find(&o).Error != nil {
		return nil
	}
	return
}

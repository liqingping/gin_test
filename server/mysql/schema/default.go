package schema

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type User struct {
	ID	string
	UserID  uint
	Number  string
	gorm.Model
}

func (User) TableName() string {
	return "a_user"
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	uid := uuid.NewV4()
	_ = scope.SetColumn("ID", uid.String())
	return nil
}
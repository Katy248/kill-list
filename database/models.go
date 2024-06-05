package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	PasswordHash string
}
type List struct {
	gorm.Model
	Name        string
	OwnerID     uint
	Description string
}
type ListItem struct {
	gorm.Model
	ListID uint
	Name   string
}
type ListItemDataType struct {
	gorm.Model
	ListID uint
	Name   string
}
type ListItemData struct {
	gorm.Model
	ListItemID uint
	TypeID     uint
	Value      string
}

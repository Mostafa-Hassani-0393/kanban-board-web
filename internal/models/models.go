package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	Password string
	Boards   []Board
}

type Board struct {
	gorm.Model
	Name   string
	UserID uint
	List   []List
}

type List struct {
	gorm.Model
	Name    string
	BoardID uint
	Cards   []Card
}

type Card struct {
	gorm.Model
	Title  string
	Desc   string
	ListID uint
}

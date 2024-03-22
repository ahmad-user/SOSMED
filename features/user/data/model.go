package data

import (
	comment "ALTA_BE_SOSMED/features/comment/data"
	post "ALTA_BE_SOSMED/features/post/data"
)

type User struct {
	UserID    int
	Nama      string
	Email     string `gorm:"type:varchar(30);primaryKey"`
	Password  string
	Picture   string
	Tgl_lahir string
	Gender    bool
	Post      []post.Post       `gorm:"foreignKey:Pemilik;references:Email"`
	Comment   []comment.Comment `gorm:"foreignKey:Pemiliks;references:Email"`
}

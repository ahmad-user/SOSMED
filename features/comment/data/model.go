package data

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Comment  string `json:"comment" form:"comment" validate:"required"`
	Pemiliks string `gorm:"type:varchar(30)"`
	Post_id  int    `gorm:"type:int(11)"`
}

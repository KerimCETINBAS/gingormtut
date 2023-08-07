package types

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title    string  `json:"title"`
	Author   *Author `json:"author"`
	AuthorID *uint
}

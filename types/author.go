package types

import "gorm.io/gorm"

type Author struct {
	gorm.Model
	Name  string   `json:"name"`
	Blogs *[]*Blog `json:"blogs"`
}

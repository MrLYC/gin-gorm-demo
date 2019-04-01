package models

import "github.com/jinzhu/gorm"

// Pairs :
type Pairs struct {
	gorm.Model
	Key   string `gorm:"type:varchar(100);unique_index"`
	Value string `gorm:"size:255"`
}

package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model  `json:",omitempty"`
	Name        string
	Description string
	//DeletedBy uint
	//CreatedBy uint
	//UpdatedBy uint
	//RestoredBy uint
}

package gateways

import (
	"github.com/jinzhu/gorm"
)

type (
	LoginRepository struct {
		Conn *gorm.DB
	}
)


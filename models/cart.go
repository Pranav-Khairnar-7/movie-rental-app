package model

import uuid "github.com/satori/go.uuid"

type Cart struct {
	Base
	UserID    uuid.UUID   `json:"userID" gorm:"type:VARCHAR(36);"`
	CartItem  []*CartItem `json:"cartItem"`
	CartValue float64     `json:"cartValue" gorm:"type:DECIMAL(5,2)"`
}

func (Cart) TableName() string {
	return "cart"
}

type CartItem struct {
	Base
	CartID  uuid.UUID `json:"cartID" gorm:"type:VARCHAR(36);"`
	MovieID uuid.UUID `json:"movieID" gorm:"type:VARCHAR(36);"`
	Movie   *Movie
}

func (CartItem) TableName() string {
	return "cart_item"
}

package models

import "time"

type WalletTransaction struct {
	ID              int       `json:"id" gorm:"column:Id;autoIncrement;not null" validate:"required"`
	DateTransaction time.Time `json:"dateTransaction" gorm:"column:DateTransaction;not null;autoCreateTime"`
	TypeTransaction int8      `json:"typeTransaction" gorm:"column:TypeTransaction;not null" validate:"required"`
	FromId          int       `json:"fromId" gorm:"column:FromId;not null" validate:"required"`
	ToId            int       `json:"toId" gorm:"column:ToId;not null" validate:"required"`
	Amount          int       `json:"amount" gorm:"column:Amount;default:0;not null" validate:"required"`
	DateCreated     time.Time `json:"dateCreated" gorm:"column:DateCreated;default:CURRENT_TIMESTAMP"`
}

func (c *WalletTransaction) TableName() string {
	return "wallet_transactions"
}

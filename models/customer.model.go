package models

import "time"

type Customer struct {
	ID          int       `json:"id" gorm:"column:Id;autoIncrement"`
	UserName    string    `json:"userName" gorm:"column:UserName;not null" validate:"required"`
	Password    string    `json:"password" gorm:"column:Password;not null" validate:"required"`
	FullName    string    `json:"fullName" gorm:"column:FullName;not null" validate:"required"`
	Balance     int       `json:"balance" gorm:"column:Balance;default:0"`
	DateCreated time.Time `json:"dateCreated" gorm:"column:DateCreated;default:CURRENT_TIMESTAMP"`
}

func (c *Customer) TableName() string {
	return "customers"
}

type RequestCustomerBalance struct {
	UserName string `json:"userName" validate:"required"`
}

type ResponseCustomerBalance struct {
	Id       int
	UserName string
	Balance  int
}

type RequestCustomerTransfer struct {
	ToUserName string `json:"to_username" validate:"required"`
	Amount     int    `json:"amount" validate:"required"`
}

type RequestCustomerBalanceTopUp struct {
	Amount int `json:"amount" validate:"required"`
}

type ResponseCustomerTopTransaction struct {
	UserName        string
	TransactedValue int
}

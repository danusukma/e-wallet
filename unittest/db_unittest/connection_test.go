package db_unittest

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dialect := mysql.Open("root:admin1234@tcp(localhost:3306)/e_wallet?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestDeleteExecuteSQL(t *testing.T) {
	err := db.Exec("DELETE FROM customers WHERE UserName = ?", "User-A").Error
	//fmt.Println(err.Error())
	assert.Nil(t, err)

}

func TestInsertExecuteSQL(t *testing.T) {
	err := db.Exec("INSERT INTO customers (UserName,Password,FullName) VALUES (?, ?, ?)", "User-A", "Password-A", "FullName-A").Error
	assert.Nil(t, err)

}

package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	db, err := gorm.Open("mysql", "root:7QZKO020y1@tcp(localhost:3306)/conscience_backend?parseTime=true&charset=utf8mb4&loc=Hongkong")
	assert.NoError(t, err)
	defer db.Close()

	Init(db)
}

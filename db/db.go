package db

import (
	"github.com/glebarez/sqlite"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

var PasteTx *TxWrapper[Paste]
var DB *gorm.DB

func init() {
	db := lo.Must(gorm.Open(sqlite.Open("data.db"), &gorm.Config{}))
	PasteTx = NewTxWrapper[Paste](db)

	lo.Must0(db.AutoMigrate(&Paste{}))
}

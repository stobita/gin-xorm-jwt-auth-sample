package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	engine, _ = xorm.NewEngine("mysql", "go_test:go_test@/go_test")
}

package db

import (
	"fmt"

	"github.com/HarvestStars/gopool/protocol"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// mysql接口
var DataBase *gorm.DB

// Setup 启动mysql配置
func Setup(user string, pwd string, host string, db string) {
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, db)
	var err error
	DataBase, err = gorm.Open("mysql", url)
	if err != nil {
		panic("failed to connect database")
	}
	DataBase.AutoMigrate(&protocol.MinerInfo{}, &protocol.BlockMined{}, &protocol.LiquidInfo{}, &protocol.LiquidHeight{})
}

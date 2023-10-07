package db

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	UserName string = "root"
	Password string = "1234"
	Port     int    = 3306
	Addr     string = "127.0.0.1"
	Database string = "warehouse_db"
)

var once sync.Once
var unUseDB *MysqlDB

type MysqlDB struct {
	mysql *gorm.DB
}

func (mysql *MysqlDB) Init() {
	once.Do(func() {
		unUseDB = &MysqlDB{
			mysql: InitDB(),
		}
	})
}

func InitDB() *gorm.DB {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", UserName, Password, Addr, Port, Database)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		log.Fatal(err)
	}

	// enable, series check schema
	db.SingularTable(true)
	// enable, db log
	db.LogMode(true)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return db
}

func GetDB() *gorm.DB {
	return InitDB()
}

package db

import (
	"flag"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	UserName string = "root"
	Addr     string = "mysql"
	Database string = "warehouse_db"
)

var (
	MysqlPassword = flag.String("mysql_password", "1234", "database password")
	MysqlPort     = flag.Int("mysql_port", 3306, "db service port")
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
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", UserName, *MysqlPassword, Addr, *MysqlPort, Database)
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

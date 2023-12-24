package model

import (
	"fmt"
	"time"
	"warehouse/db"
)

type UserInfo struct {
	ID              string    `json:"id" gorm:"column:User_ID;size:5"`
	Mail            string    `json:"mail" gorm:"column:User_Mail"`
	Password        string    `json:"password" gorm:"column:User_Password"`
	Auth            string    `json:"auth" gorm:"column:User_Auth"`
	Name            string    `json:"name" gorm:"column:User_Name"`
	Phone           string    `json:"phone" gorm:"column:User_Phone"`
	TaxID           string    `json:"taxid" gorm:"column:Tax_ID"`
	Principal       string    `json:"principal" gorm:"column:Main_Principal"`
	Connector       string    `json:"connector" gorm:"column:Main_Connector"`
	MainPhone       string    `json:"main_phone" gorm:"column:Main_Phone"`
	Brand           string    `json:"brand" gorm:"column:User_Brand"`
	SecondConnector string    `json:"second_connector", gorm:"Second_Connector"`
	SecondPhone     string    `json:"second_phone", gorm:"Second_Phone"`
	UserAddr        string    `json:"user_addr", gorm:"User_Addr"`
	UserLevel       string    `json:"user_level", gorm:"User_Level"`
	UserNote        string    `json:"user_note", gorm:"User_Note"`
	CreatedAt       time.Time `json:"createTime" gorm:"column:User_CreateTime"`
	UpdatedAt       time.Time `json:"updateTime" gorm:"column:User_UpdateTime"`
}

type UserInfoMapping struct {
	Brand   string `json:"brand" gorm:"column:User_Brand"`
	BrandCH string `json:"brand_ch" gorm:"column:User_Brand_CH"`
}

func (user *UserInfo) TableName() string {
	return "ACCOUNT_MANAGE"
}

func (user *UserInfo) Register() error {
	mysql := db.GetDB()

	// sql := fmt.Sprintf(`INSERT INTO ACCOUNT_MANAGE(User_ID, User_Mail, User_Password, User_name, User_Auth, User_Phone, Main_Phone, Main_Principal, Main_Connector, User_Brand) VALUES("%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s")`,
	// 	user.ID, user.Mail, user.Password, user.Name, user.Auth, user.Phone,
	// 	user.MainPhone, user.Principal, user.Connector, user.Brand)
	// mysql.Exec(sql)
	return mysql.Create(&user).Error
}

func (user *UserInfo) Login(ip string) error {
	// user store in db
	mysql := db.GetDB()
	return mysql.Where("(User_Mail = ? or User_Phone = ?) AND User_Password = ?",
		user.Mail, user.Phone, user.Password).First(&user).Error
}

func (user *UserInfo) GetUser() {

	// user store in db
	mysql := db.GetDB()
	if err := mysql.Where("User_ID = ?",
		user.ID).Find(&user).Error; err != nil {

		fmt.Println("not get")
		return
	}
}

func (user *UserInfo) GetAllUser() (users []UserInfo) {
	// user store in db
	mysql := db.GetDB()
	if err := mysql.Select("*").Find(&users); err != nil {

		fmt.Println("not get")
		return
	}

	return
}

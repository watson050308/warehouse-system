package userController

import (
	"fmt"
	"log"
	"warehouse/db"
	"warehouse/model"
)

func UserRegister(user model.UserInfo) {
	mysql := db.GetDB()

	sql := fmt.Sprintf(`INSERT INTO ACCOUNT_MANAGE(User_ID, User_Mail, User_Password, User_name, 
		                                           User_Auth, User_Phone, Main_Phone, Main_Principal,
												   Main_Connector, User_Brand)
												   VALUES("%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s", "%s")`,
		user.ID, user.Mail, user.Password, user.Name, user.Auth, user.Phone,
		user.MainPhone, user.Principal, user.Connector, user.Brand)
	mysql.Exec(sql)
}

func UserLogin(user model.UserInfo, ip string) {
	if (user.Mail == "" && user.Phone == "") || user.Password == "" {
		log.Println("not get mail, phone or password")
		return
	}

	// user store in db
	mysql := db.GetDB()
	if err := mysql.Where("User_Mail = ? or User_Phone = ? AND User_Password = ?",
		user.Mail, user.Phone, user.Password).First(&user).Error; err != nil {

		log.Println("not get")
		return
	}
}

func GetUser(user *model.UserInfo, id string) {
	if id == "" {
		fmt.Println("not get id")
		return
	}

	// user store in db
	mysql := db.GetDB()
	if err := mysql.Where("User_ID = ?",
		id).Find(&user).Error; err != nil {

		log.Println("not get")
		return
	}
}

func GetAllUser(users *[]model.UserInfo) {
	// user store in db
	mysql := db.GetDB()
	if err := mysql.Select("*").Find(&users); err != nil {

		log.Println("not get")
		return
	}
}

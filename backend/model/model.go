package model

import "time"

type UserInfo struct {
	ID         string    `json:"id" gorm:"column:User_ID;size:5"`
	Mail       string    `json:"mail" gorm:"column:User_Mail"`
	Password   string    `json:"password" gorm:"column:User_Password"`
	Auth       string    `json:"auth" gorm:"column:User_Auth"`
	Name       string    `json:"name" gorm:"column:User_Name"`
	Phone      string    `json:"phone" gorm:"column:User_Phone"`
	TaxID      string    `json:"taxid" gorm:"column:Tax_ID"`
	Principal  string    `json:"principal" gorm:"column:Main_Principal"`
	Connector  string    `json:"connector" gorm:"column:Main_Connector"`
	MainPhone  string    `json:"main_phone" gorm:"column:Main_Phone"`
	Brand      string    `json:"brand" gorm:"column:User_Brand"`
	CreateTime time.Time `json:"createTime" gorm:"column:User_CreateTime"`
}

func (info *UserInfo) TableName() string {
	return "ACCOUNT_MANAGE"
}

type ProductInfo struct {
	ID           string `json:"id"`
	ProductName  string `json:"productName"`
	ProductCount int    `json:"productCount"`
	ProductUnit  int    `json:"productUint"`
	UnitType     string `json:"uintType"`
	ProductType  string `json:"productType"`
	MinimumCount int    `json:"minimumCount"`
	ProductState string `json:"productState"`
}

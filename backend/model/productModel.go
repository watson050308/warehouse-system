package model

import (
	"fmt"
	"warehouse/db"
)

type ProductInfo struct {
	ID           string  `json:"id" gorm:"column:Product_ID;size:4"`
	Name         string  `json:"name" gorm:"column:Product_Name"`
	Desc         string  `json:"desc" gorm:"column:Product_Desc"`
	Price        int     `json:"price" gorm:"column:Product_Price"`
	OnSalePrice  int     `json:"on_sale_price" gorm:"column:On_Sale_Price"`
	CostPrice    float32 `json:"cost_price" gorm:"column:Cost_Price"`
	Unit         string  `json:"uint" gorm:"column:Product_Unit"`
	State        string  `json:"state" gorm:"column:Product_State"`
	Count        int     `json:"count" gorm:"column:Product_Count"`
	SafetyCount  int     `json:"safety_count" gorm:"column:Safety_Count"`
	DeliveryTime string  `json:"delivery_time" gorm:"column:Delivery_Time"`
}

type ProductMapping struct {
	Type   string `json:"type" gorm:"column:Product_Type"`
	TypeCH string `json:"type_ch" gorm:"column:Product_Type_CH"`
}

func (product *ProductInfo) TableName() string {
	return "PRODUCT_MANAGE"
}

func (product *ProductInfo) GetAllItem() (products []ProductInfo) {
	mysql := db.GetDB()
	mysql.Table("PRODUCT_MANAGE").Select("*").Where("Product_Name LIKE ?",
		"%"+product.Name+"%").Scan(&products)
	return
}

func (product *ProductInfo) GetItem() {
	mysql := db.GetDB()
	if err := mysql.Where("Product_ID = ?",
		product.ID).Find(&product).Error; err != nil {
		fmt.Println("not get item")
		return
	}
}

func (product *ProductInfo) AddItem() error {
	mysql := db.GetDB()
	return mysql.Create(&product).Error
}

func (product *ProductInfo) DeleteItem(id string) error {
	mysql := db.GetDB()

	return mysql.Where("Product_ID = ?", id).Delete(&product).Error
}

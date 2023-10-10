CREATE DATABASE warehouse_db;

CREATE TABLE warehouse_db.PRODUCT_MANAGE
(
    Product_ID char(4) NOT NULL,
    Product_Name nvarchar(50) NOT NULL,
    Product_Desc nvarchar(50),
    Product_Price int NOT NULL,
    On_Sale_Price int,
    Cost_Price float NOT NULL,
    Product_Unit nchar(2) NOT NULL,
    Product_State nvarchar(10) NOT NULL,
    Product_Count int NOT NULL,
    Safety_Count int NOT NULL,
    Delivery_Time nvarchar(10)
)
;

CREATE TABLE warehouse_db.PRODUCT_MAPPING
(
    Product_Type char(1) NOT NULL,
    Product_Type_CH nchar(2) NOT NULL
)
;

CREATE TABLE warehouse_db.ACCOUNT_MANAGE
(
    User_ID char(5) NOT NULL,
    User_Mail char(100) NOT NULL,
    User_Password char(20) NOT NULL,
    User_Auth char(1),
    User_Name nvarchar(20) NOT NULL,
    User_Phone char(10) NOT NULL,
    Tax_ID char(8),
    Main_Principal nvarchar(10) NOT NULL,
    Main_Connector nvarchar(10) NOT NULL,
    Main_Phone char(10) NOT NULL,
    Second_Connector nvarchar(10),
    Second_Phone char(10),
    User_Addr nvarchar(50),
    User_Brand char(2) NOT NULL,
    User_Level char(1),
    User_CreateTime timestamp
)
;

CREATE TABLE warehouse_db.USER_BRAND_MAPPING
(
    User_Brand char(2) NOT NULL,
    User_Brand_CH char(10) NOT NULL
)
;

CREATE TABLE warehouse_db.ORDER_MANAGE
(
    Order_ID char(10) NOT NULL,
    User_ID char(5) NOT NULL,
    Product_ID char(5) NOT NULL,
    Product_Count int NOT NULL,
    Order_State char(1) NOT NULL,
    Order_Note nvarchar(50),
    Order_Date datetime
)
;

ALTER TABLE warehouse_db.PRODUCT_MANAGE
ADD PRIMARY KEY (Product_ID);

ALTER TABLE warehouse_db.PRODUCT_MAPPING
ADD PRIMARY KEY (Product_Type);

ALTER TABLE warehouse_db.ACCOUNT_MANAGE
ADD PRIMARY KEY (User_ID);

ALTER TABLE warehouse_db.ORDER_MANAGE
ADD PRIMARY KEY (Order_ID, User_ID, Product_ID);


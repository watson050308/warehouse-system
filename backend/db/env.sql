CREATE DATABASE warehouse_db;

CREATE TABLE warehouse_db.PRODUCT_MANAGE
(
    Cost_Price float NOT NULL,
    Delivery_Time varchar(10),
    On_Sale_Price int,
    Product_Count int NOT NULL,
    Product_Desc nvarchar(100),
    Product_ID char(4) NOT NULL,
    Product_Name nvarchar(50) NOT NULL,
    Product_Price int NOT NULL,
    Product_State char(1),
    Product_Unit char(1),
    Safety_Count int NOT NULL
)
;

CREATE TABLE warehouse_db.PRODUCT_MAPPING
(
    Product_Type char(1) NOT NULL,
    Product_Type_CH nchar(20) NOT NULL
)
;

CREATE TABLE warehouse_db.ACCOUNT_MANAGE
(
    Main_Connector nvarchar(10) NOT NULL,
    Main_Phone char(10) NOT NULL,
    Main_Principal nvarchar(10) NOT NULL,
    Second_Connector nvarchar(10),
    Second_Phone char(10),
    Tax_ID char(8),
    User_Addr nvarchar(50),
    User_Auth char(1) NOT NULL,
    User_Brand char(2) NOT NULL,
    User_CreateTime timestamp,
    User_ID char(5) NOT NULL,
    User_Level char(1),
    User_Mail char(100) NOT NULL,
    User_Name nvarchar(50) NOT NULL,
    User_Note nvarchar(50),
    User_Password char(20) NOT NULL,
    User_Phone char(10) NOT NULL,
    User_UpdateTime timestamp
)
;

CREATE TABLE warehouse_db.USER_BRAND_MAPPING
(
    User_Brand char(2) NOT NULL,
    User_Brand_CH nchar(10) NOT NULL
)
;

CREATE TABLE warehouse_db.ORDER_MANAGE
(
    Order_Date datetime,
    Order_ID char(11) NOT NULL,
    Order_Note nvarchar(100),
    Order_State char(1) NOT NULL,
    Product_Count int NOT NULL,
    Product_ID char(5) NOT NULL,
    User_ID char(5) NOT NULL
)
;

CREATE TABLE warehouse_db.RESPOND_MANAGE
(
    Respond_Date datetime,
    Respond_ID char(11) NOT NULL,
    Respond_State char(1) NOT NULL,
    Respond_Text nvarchar(100),
    User_ID char(5) NOT NULL
)
;

ALTER TABLE warehouse_db.PRODUCT_MANAGE
ADD PRIMARY KEY (Product_ID);

ALTER TABLE warehouse_db.PRODUCT_MAPPING
ADD PRIMARY KEY (Product_Type);

ALTER TABLE warehouse_db.ACCOUNT_MANAGE
ADD PRIMARY KEY (User_ID);

ALTER TABLE warehouse_db.USER_BRAND_MAPPING
ADD PRIMARY KEY (User_Brand);

ALTER TABLE warehouse_db.ORDER_MANAGE
ADD PRIMARY KEY (Order_ID, User_ID, Product_ID);

ALTER TABLE warehouse_db.RESPOND_MANAGE
ADD PRIMARY KEY (Respond_ID);


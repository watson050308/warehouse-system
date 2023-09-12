CREATE DATABASE warehouse_db;

CREATE TABLE warehouse_db.PRODUCT_MANAGE
(
    Product_ID char(5) NOT NULL,
    Product_Name nvarchar(20) NOT NULL,
    Product_Desc nvarchar(50),
    Product_Count int NOT NULL,
    Product_Price int NOT NULL,
    Product_Unit char(1),
    Product_Type char(1),
    Product_State char(1),
    Warehouse_Unit varchar(10),
    Delivery_Time nvarchar(10)
)
;

CREATE TABLE warehouse_db.MEMBER_MANAGE
(
    User_ID char(5) NOT NULL,
    User_Name nvarchar(10),
    User_Brand char(2),
    User_Level char(1)
)
;

CREATE TABLE warehouse_db.ACCOUNT_MANAGE
(
    User_ID char(5) NOT NULL,
    User_Account char(20) NOT NULL,
    User_Password char(20) NOT NULL,
    User_Auth char(1) NOT NULL,
    User_CreateTime timestamp
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

ALTER TABLE warehouse_db.MEMBER_MANAGE
ADD PRIMARY KEY (User_ID);

ALTER TABLE warehouse_db.ACCOUNT_MANAGE
ADD PRIMARY KEY (User_ID);

ALTER TABLE warehouse_db.ORDER_MANAGE
ADD PRIMARY KEY (Order_ID, User_ID, Product_ID);


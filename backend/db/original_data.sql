INSERT INTO ACCOUNT_MANAGE(User_ID,User_Mail,User_Password,User_Auth,User_Name,User_Phone,Tax_ID,Main_Principal,Main_Connector,Main_Phone,User_Addr,User_Brand,User_Note)
VALUES
("A0001","RO.guangfu13@gmail.com","abc27477566","1","松山光復店-直營店","0227473388","25622146","蔡林高","呂知穎","0912345678","光復南路13巷29號1樓","A","月結")
;

INSERT INTO USER_BRAND_MAPPING(User_Brand, User_Brand_CH)
VALUES
 ("A", "紅橘子店家")
,("B", "愛餡貓店家")
,("Z", "others")
;

//2024-1-14 sprint1:insert data in table: PRODUCT_MANAGE
INSERT INTO PRODUCT_MANAGE(Cost_Price,Delivery_Time,On_Sale_Price,Product_Count,Product_Desc,Product_ID,Product_Name,Product_Price,Product_State,Product_Unit,Safety_Count)
VALUES
 ("165","不確定","","18","東穎-1藍8入:line@","A0001","沙拉A5-8入/藍","235","正常供應中","盒","40")
,("35.417","不確定","","521","24 包=1箱/翔淇/每次叫480盒","A0002","紅橘子即溶檸檬冰紅茶(375g)-24入/箱","65","正常供應中","包","130")
;


//2024-1-14 sprint1:insert test data in tables
INSERT INTO PRODUCT_MANAGE(Cost_Price,Delivery_Time,On_Sale_Price,Product_Count,Product_Desc,Product_ID,Product_Name,Product_Price,Product_State,Product_Unit,Safety_Count)
VALUES
 ("165","不確定","","18","東穎-1藍8入:line@","A0001","沙拉A5-8入/藍","235","正常供應中","盒","40")
,("35.417","不確定","","521","24 包=1箱/翔淇/每次叫480盒","A0002","紅橘子即溶檸檬冰紅茶(375g)-24入/箱","65","正常供應中","包","130")
,("161.9","不確定","","23","R0131喜多/12~24","A0003","顆粒柳橙","240","正常供應中","瓶","12")
,("220","1個月","","33","冷藏產品測試","B0001","Product_Name test 1","88","缺貨測試中","袋","100")
,("88","2週","111","44","冷凍產品測試","C0001","Product_Name 測試2號","222","缺貨測試中","箱","200")
,("59","3天","","55","包材產品測試+字很多的測試 包含空格 包含abcd123","D0001","Product_Name 測試3號","99","測試中","條","300")
,("77.7","4個工作日","","66","其他產品測試","E0001","Product_Name 測試4號","200","永久測試下架","束","400")
,("33.333","5日","100","77","愛餡貓包材產品測試","H0001","Product_Name test測試5號","99","正常供應測試中","包","500")
;

INSERT INTO PRODUCT_MAPPING(Product_Type,Product_Type_CH)
VALUES
 ("A","常溫")
,("B","冷藏")
,("C","冷凍")
,("D","包材")
,("E","其他")
,("F","愛餡貓小物")
,("G","愛餡貓常溫")
,("H","愛餡貓包材")
,("I","愛餡貓冷藏")
,("J","愛餡貓冷凍")
;

INSERT INTO ACCOUNT_MANAGE(User_ID,User_Mail,User_Password,User_Auth,User_Name,User_Phone,Tax_ID,Main_Principal,Main_Connector,Main_Phone,User_Addr,User_Brand,User_Note)
VALUES
 ("A0001","RO.guangfu13@gmail.com","abc27477566","1","松山光復店-直營店","0227473388","25622146","蔡林高","呂知穎","0912345678","光復南路13巷29號1樓","A","月結")
,("B0001","test002@gmail.com","abctest002","1","愛餡貓店家-測試店","0227473002","00200202","新垣結衣","星野源","0911111111","光復南路test01巷test01號","B","日結")
,("C0001","test003@gmail.com","abctest003","1","廠商1-測試店","0227473003","00300303","多拉A夢","大雄","0922222222","光復南路test02巷test02號","C","日結下貨收現")
,("D0001","test004@gmail.com","abctest004","1","廠商2-測試店","0227473004","00400404","蠟筆小新","小白","0933333333","光復南路test04巷test04號","D","月結匯款")
,("Z0001","test005@gmail.com","abctest005","1","others-測試店","0227473005","00500505","卡比","瓦豆魯迪","0944444444","光復南路test05巷test05號","Z","月結日結")
;

INSERT INTO USER_BRAND_MAPPING(User_Brand,User_Brand_CH)
VALUES
 ("A","紅橘子店家")
,("B","愛餡貓店家")
,("C","廠商1")
,("D","廠商2")
,("Z","others")
;

INSERT INTO ORDER_MANAGE(Order_Date,Order_ID,Order_Note,Order_State,Product_Count,Product_ID,User_ID)
VALUES
 ("2024-01-12 00:00:00","20240112001","希望快點到貨！","0","Product_Count","Product_ID","A0001")
,("2023-12-31 00:00:00","20231231001","試試看這樣的字會不會太多呢？？？？？？？？？？？？？？？？？","1","Product_Count","Product_ID","B0001")
,("2023-11-28 00:00:00","20231128001","","2","Product_Count","Product_ID","C0001")
,("2024-01-05 00:00:00","20240105001","Order_Note中英測試","3","Product_Count","Product_ID","D0001")
,("2023-05-26 00:00:00","20230526001","","5","Product_Count","Product_ID","Z0001")
;

INSERT INTO RESPOND_MANAGE(Respond_Date,Respond_ID,Respond_State,Respond_Text,User_ID)
VALUES
 ("2023-05-26 00:00:00","20230526001","0","送貨來少了一條起司","A0001")
,("2024-01-05 00:00:00","20240105001","0","今天司機大哥找錯錢了，多找了100塊！","B0001")
,("2024-01-12 00:00:00","20240112001","1","請問缺貨的脆皮雞柳什麼時候會有貨？","C0001")
,("2023-12-31 00:00:00","20231231001","1","過年期間有哪幾天放假不送貨？","D0001")
,("2023-11-28 00:00:00","20231128001","0","testttttttt","Z0001")
;

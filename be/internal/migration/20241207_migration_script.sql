--create table user
CREATE TABLE "user" (
                        user_id bigserial PRIMARY KEY,
                        username text NOT NULL UNIQUE,
                        password text NOT NULL
);

--insert default value table users
INSERT INTO "user" (username, password)
VALUES ('admin', '28f0116ef42bf718324946f13d787a1d41274a08335d52ee833d5b577f02a32a');



--create reserved_money table
CREATE TABLE reserved_money (
                                coins1 int8 default 0,
                                coins5 int8 default 0,
                                coins10 int8 default 0,
                                bank20 int8 default 0,
                                bank50 int8 default 0,
                                bank100 int8 default 0,
                                bank500 int8 default 0,
                                bank1000 int8 default 0
);

--insert default value to reserved_money table
insert into reserved_money (coins1,coins5,coins10,bank20,bank50,bank100,bank500,bank1000)
values (100,100,100,100,100,100,100,100);



--create table product
CREATE TABLE product (
                         product_id text PRIMARY KEY,
                         name text NOT NULL,
                         image text,
                         stock int8 default 0,
                         price int8 default 0
);

--insert default value table product
INSERT INTO product (product_id , name, image, stock, price)
VALUES ('61ad2981-9423-4071-97dd-4fa84fda5fa3','จิ้งจอกน้อยมาแล้วจ้าาาา','https://png.pngtree.com/png-clipart/20240111/original/pngtree-cute-little-fox-cartoon-hand-drawn-elements-png-image_14082007.png',10,1999),
       ('f95b325d-f0ed-4b2c-8467-78fb4dc0030c','ไก่สดมาแล้วจ้าาา','https://media.istockphoto.com/id/951241438/th/%E0%B9%80%E0%B8%A7%E0%B8%84%E0%B9%80%E0%B8%95%E0%B8%AD%E0%B8%A3%E0%B9%8C/%E0%B9%84%E0%B8%81%E0%B9%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%8C%E0%B8%95%E0%B8%B9%E0%B8%99%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B9%81%E0%B8%82%E0%B9%87%E0%B8%87%E0%B9%81%E0%B8%81%E0%B8%A3%E0%B9%88%E0%B8%87.jpg?s=1024x1024&w=is&k=20&c=sYOOWSRKmMILt2IYLOAvCpDow8b7Nzgiz28cIMHbOzQ%3D',100,9999),
       ('49030f6a-14a7-48db-9e00-39efd0f5e4f7','กบกระโดดอ๊บอ๊บ','https://media.istockphoto.com/id/1456949429/th/%E0%B9%80%E0%B8%A7%E0%B8%84%E0%B9%80%E0%B8%95%E0%B8%AD%E0%B8%A3%E0%B9%8C/%E0%B8%97%E0%B8%B2%E0%B8%A3%E0%B8%81%E0%B8%81%E0%B8%9A%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%8C%E0%B8%95%E0%B8%B9%E0%B8%99%E0%B8%99%E0%B9%88%E0%B8%B2%E0%B8%A3%E0%B8%B1%E0%B8%81%E0%B8%99%E0%B8%B1%E0%B9%88%E0%B8%87.jpg?s=612x612&w=0&k=20&c=S3qsZo3PkydT5hLF9EJsYtjpclVrQnpcgRhOM7L4E58%3D',0,899)
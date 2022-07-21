```CREATE TABLE items (
    id int NOT NULL AUTO_INCREMENT,
    title varchar(255) NOT NULL,
    seller varchar(255) ,
    price int NOT NULL,
    stock int NOT NULL,
    updatedAt datetime,
    createdAt datetime,
    sold_quantity int,
    status varchar(50),
    PRIMARY KEY(id)
)```
CREATE TABLE items (
    id int NOT NULL AUTO_INCREMENT,
    categoryId int NOT NULL,
    name varchar(255) NOT NULL,
    image varchar(255) NOT NULL,
    price int NOT NULL,
    stock int NOT NULL,
    updatedAt datetime,
    createdAt datetime,
    discountQty int,
    discountType varchar(40),
    discountResult int,
    discountExpiredAt varchar(255),
    PRIMARY KEY(id)
)

ALTER TABLE products
modify column createdAt varchar(255);

curl --location --request PUT 'http://localhost:3030/products/1' \
--data-raw '{
    "categoryId": 1
}'

CREATE TABLE payments (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    type varchar(255) NOT NULL,
    logo varchar(255) ,
    updatedAt varchar(255),
    createdAt varchar(255),
    PRIMARY KEY(id)
)   

CREATE TABLE plan (
    planId int NOT NULL AUTO_INCREMENT,
    planName varchar(255) NOT NULL,
    amountOptions int,
    tenureOptions int NOT NULL,
    benefitPercentage int NOT NULL,
    benefitType varchar(255) NOT NULL,
    promotionPercentage int, 
    promotionUsers int,
    promotionStartDate datetime,
    promotionEndDate datetime,
    updatedAt datetime,
    createdAt datetime,
    PRIMARY KEY(planId)
)

CREATE TABLE customerGoals (
    goalId int NOT NULL AUTO_INCREMENT,
    planId int NOT NULL, 
    userId int NOT NULL, 
    selectedAmount int NOT NULL,
    selectedTenure int NOT NULL,
    startedDate varchar(255),
    depositedAmount int NOT NULL,
    benefitPercentage int NOT NULL,
    benefitType varchar(255) NOT NULL,
    updatedAt varchar(255),
    createdAt varchar(255),
    PRIMARY KEY(goalId)
)
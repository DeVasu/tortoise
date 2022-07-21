
## To create Table `plan`
```
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
    updatedAt varchar(255),
    createdAt varchar(255),
    PRIMARY KEY(planId)
)
```
## To create Table `customerGoals`
```
CREATE TABLE customerGoals (
    goalId int NOT NULL AUTO_INCREMENT,
    planId int NOT NULL, 
    userId int NOT NULL, 
    selectedAmount int NOT NULL,
    selectedTenure int NOT NULL,
    startedDate varchar(255) NOT NULL,
    depositedAmount int NOT NULL,
    benefitPercentage int NOT NULL,
    benefitType varchar(255) NOT NULL,
    updatedAt varchar(255),
    createdAt varchar(255),
    PRIMARY KEY(goalId)
)
```
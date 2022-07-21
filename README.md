# Tortoise - Plan API

# Postman Documentation
* https://documenter.getpostman.com/view/20763059/UzR1KhXx

## Points To Remember
* Database with table name `plan` and `customerGoals` should already be existing on you mysql server. If not, create them using queries in SQL_QUERIES.md file.
* After all tables are setup in your database, Update information at `app/application.go` file for database connection
* You can also set the port on which this application runs on, from `app/application.go` file (default --> 3030)

## Current Mappings

* GET `/ping` --> To test availability of the server

* GET `/admin/plan` --> To list all the available plans 
    * ### Response Body 
    ```
    [
        {
            "planId": 1,
            "planName": "buy2",
            "amountOptions": 10,
            "tenureOptions": 3,
            "benefitPercentage": 4,
            "benefitType": "cashback",
            "promotion": {
                "newPercentage": 0,
                "usersLeft": 0,
                "startDate": "",
                "endDate": ""
            },
            "updatedAt": "2006-01-02T15:04:05Z",
            "createdAt": "2006-01-02T15:04:05Z",
            "promotionValid": false
        },
        {
            "planId": 2,
            "planName": "buy3",
            "amountOptions": 10,
            "tenureOptions": 3,
            "benefitPercentage": 4,
            "benefitType": "cashback",
            "promotion": {
                "newPercentage": 12,
                "usersLeft": 0,
                "startDate": "2022-07-20T13:50:12Z",
                "endDate": "2022-08-20T13:50:12Z"
            },
            "updatedAt": "2022-07-20T13:49:45Z",
            "createdAt": "2022-07-20T13:49:45Z",
            "promotionValid": true
        }
    ]
    ```

* POST `/admin/plan` --> To create a new Plan
  * ### Request Body 
  ```
  {
    "planName": "wrapperes",
    "amountOptions": 10,
    "tenureOptions": 3,
    "benefitPercentage": 5,
    "benefitType": "extraVoucher"
  }
  ```
  * ### Response Body
  ```
  {
    "planId": 8,
    "planName": "wrapperes",
    "amountOptions": 10,
    "tenureOptions": 3,
    "benefitPercentage": 5,
    "benefitType": "extraVoucher",
    "promotion": null,
    "updatedAt": "2022-07-21T15:39:33Z",
    "createdAt": "2022-07-21T15:39:33Z",
    "promotionValid": false
  }
  ```

* PUT `/admin/plan/:planId/addPromo` --> Add promotion to an existing plan
  * ### Request Body 
    ``` 
    {
        "newPercentage": 34,
        "usersLeft": 2
    }
    ```
  * ### Response Body
  ```
    {
        "planId": 3,
        "planName": "buy3",
        "amountOptions": 10,
        "tenureOptions": 3,
        "benefitPercentage": 4,
        "benefitType": "extraVoucher",
        "promotion": {
            "newPercentage": 34,
            "usersLeft": 2,
            "startDate": "2022-07-20",
            "endDate": ""
        },
        "updatedAt": "2022-07-20T13:50:12Z",
        "createdAt": "2022-07-20T13:50:12Z",
        "promotionValid": false
    }
  ```
* DELETE `/admin/plan/:planId/deletePromo` --> Delete an Existing Promotion
    * ### Response Body
    ```
    {
        "message": "successfully deleted plan",
        "status": 200,
        "error": "none"
    }
    ```
* GET `/admin/goals` --> Lists Customer Goals
  * ### Response Body 
    ```
    [
        {
            "goalId": 1,
            "planId": 1,
            "userId": 10,
            "selectedAmount": 10,
            "selectedTenure": 3,
            "startedDate": "2022-07-21",
            "depositedAmount": 10,
            "benefitPercentage": 4,
            "benefitType": "cashback",
            "updatedAt": "2022-07-21T12:43:39Z",
            "createdAt": "2022-07-21T12:43:39Z"
        },
        {
            "goalId": 2,
            "planId": 1,
            "userId": 10,
            "selectedAmount": 10,
            "selectedTenure": 3,
            "startedDate": "2022-07-21",
            "depositedAmount": 10,
            "benefitPercentage": 4,
            "benefitType": "cashback",
            "updatedAt": "2022-07-21T12:44:50Z",
            "createdAt": "2022-07-21T12:44:50Z"
        }
    ]
    ```

* POST `/plan/:planId/enroll` --> To enroll in a plan
  * ### Request Body 
  ```
  {
      "userId": 20,
  }
  ```
  * ### Response Body
  ```
    {
        "goalId": 20,
        "planId": 3,
        "userId": 2,
        "selectedAmount": 10,
        "selectedTenure": 3,
        "startedDate": "2022-07-21",
        "depositedAmount": 10,
        "benefitPercentage": 4,
        "benefitType": "extraVoucher",
        "updatedAt": "2022-07-21T15:51:21Z",
        "createdAt": "2022-07-21T15:51:21Z"
    }
  ```



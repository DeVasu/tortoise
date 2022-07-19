package products

import (
	// "net/http"
	// "strconv"

	// "github.com/federicoleon/bookstore_oauth-go/oauth"
	// "github.com/DeVasu/tortoise/domain/users"

	// "github.com/DeVasu/tortoise/utils/errors"

	"fmt"
	"net/http"
	"strconv"

	"github.com/DeVasu/tortoise/domain/products"
	"github.com/DeVasu/tortoise/utils"
	"github.com/DeVasu/tortoise/utils/rest_errors"
	"github.com/gin-gonic/gin"
)



func getUserId(userIdParam string) (int64, rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

func Delete(c *gin.Context) {

	productId, idErr := getUserId(c.Param("productId"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var temp products.Product

	temp.Id = productId
	err := temp.Delete()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	res := &utils.Response{
		Success: true,
		Message: "Success",
	}
	
	c.JSON(http.StatusOK, res)

}

func Update(c *gin.Context) {

	productId, idErr := getUserId(c.Param("productId"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var temp products.Product
	if err := c.ShouldBindJSON(&temp); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}

	temp.Id = productId

	err := temp.Update()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	res := &utils.Response{
		Success: true,
		Message: "Success",
	}
	
	c.JSON(http.StatusOK, res)

}

func GetById(c *gin.Context) {
	
	productId, idErr := getUserId(c.Param("productId"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	result := &products.Product{
		Id: productId,
	}
	err := result.GetById()
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// //i want to list cashiers
func List(c *gin.Context) {

	result := &products.Product{}
	

	listOf, err := result.List()
	fmt.Println(listOf[:2])


	if err != nil {
		c.JSON(http.StatusBadRequest, "\"{\"err\":\"wrong\"}")
		return
	}

	c.JSON(http.StatusOK, listOf)

}

func Create(c *gin.Context) {

	res := &products.Product{}
	
	if err := c.ShouldBindJSON(&res); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	
	err := res.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, res)

}
package goals

import (
	// "net/http"
	// "strconv"

	// "github.com/federicoleon/bookstore_oauth-go/oauth"
	// "github.com/DeVasu/tortoise/domain/users"

	// "github.com/DeVasu/tortoise/utils/errors"

	// "fmt"
	"net/http"
	"strconv"

	// "strconv"

	"github.com/DeVasu/tortoise/domain/plan"
	// "github.com/DeVasu/tortoise/utils"
	rest_errors "github.com/DeVasu/tortoise/utils/errors"
	"github.com/gin-gonic/gin"
)



func Create(c *gin.Context) {

	res := &plan.Plan{}
	
	if err := c.ShouldBindJSON(&res); err != nil {
		resErr := rest_errors.NewBadRequestError(err.Error())
		c.JSON(resErr.Status, resErr)
		return
	}
	
	err := res.Create()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, res)

}



func getPlanId(planIdParam string) (int64, *rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(planIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}


// func sortPlans([]plan.Plan) []plan.Plan
// func Delete(c *gin.Context) {

// 	productId, idErr := getUserId(c.Param("productId"))
// 	if idErr != nil {
// 		c.JSON(idErr.Status(), idErr)
// 		return
// 	}

// 	var temp products.Product

// 	temp.Id = productId
// 	err := temp.Delete()
// 	if err != nil {
// 		c.JSON(err.Status(), err)
// 		return
// 	}
// 	res := &utils.Response{
// 		Success: true,
// 		Message: "Success",
// 	}
	
// 	c.JSON(http.StatusOK, res)

// }

// func Update(c *gin.Context) {

// 	productId, idErr := getUserId(c.Param("productId"))
// 	if idErr != nil {
// 		c.JSON(idErr.Status(), idErr)
// 		return
// 	}

// 	var temp products.Product
// 	if err := c.ShouldBindJSON(&temp); err != nil {
// 		c.JSON(http.StatusBadGateway, err)
// 		return
// 	}

// 	temp.Id = productId

// 	err := temp.Update()
// 	if err != nil {
// 		c.JSON(err.Status(), err)
// 		return
// 	}

// 	res := &utils.Response{
// 		Success: true,
// 		Message: "Success",
// 	}
	
// 	c.JSON(http.StatusOK, res)

// }

// func GetById(c *gin.Context) {
	
// 	productId, idErr := getUserId(c.Param("productId"))
// 	if idErr != nil {
// 		c.JSON(idErr.Status(), idErr)
// 		return
// 	}
// 	result := &products.Product{
// 		Id: productId,
// 	}
// 	err := result.GetById()
// 	if err != nil {
// 		c.JSON(err.Status(), err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, result)
// }




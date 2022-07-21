package goals

import (
	"net/http"
	"strconv"

	goals "github.com/DeVasu/tortoise/domain/customerGoals"
	rest_errors "github.com/DeVasu/tortoise/utils/errors"
	"github.com/gin-gonic/gin"
)



func Create(c *gin.Context) {


	planId, idErr := getPlanId(c.Param("planId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var tempUser goals.User
	if err := c.ShouldBindJSON(&tempUser); err != nil {
		resErr := rest_errors.NewBadRequestError(err.Error())
		c.JSON(resErr.Status, resErr)
		return
	}

	res := &goals.Goal{
		UserId: tempUser.UserId,
		PlanId: planId,
	}
	
	err := res.Create()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, res)

}

func List(c *gin.Context) {

	result := &goals.Goal{}
	

	listOf, err := result.List()
	// fmt.Println(listOf[:2])

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, listOf)

}

func getPlanId(planIdParam string) (int64, *rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(planIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}
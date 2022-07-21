package plan

import (
	"net/http"
	"strconv"

	"github.com/DeVasu/tortoise/domain/plan"
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

//i want to list cashiers
func List(c *gin.Context) {

	result := &plan.Plan{}
	

	listOf, err := result.List()
	// fmt.Println(listOf[:2])

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, listOf)

}

func AddPromotion(c *gin.Context) {

	planId, idErr := getPlanId(c.Param("planId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var promo plan.Promotion
	if err := c.ShouldBindJSON(&promo); err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}

	plan := plan.Plan{
		PlanId: planId,
	}

	err := plan.AddPromotion(promo)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	// res := &utils.Response{
	// 	Success: true,
	// 	Message: "Success",
	// }
	
	c.JSON(http.StatusOK, plan)

}

func DeletePromotion(c *gin.Context) {

	planId, idErr := getPlanId(c.Param("planId"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	plan := plan.Plan{
		PlanId: planId,
	}

	err := plan.DeletePromotion()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	successMsg := rest_errors.NewRestErr("successfully deleted plan", http.StatusOK, "none")
	c.JSON(successMsg.Status, successMsg)

}

func getPlanId(planIdParam string) (int64, *rest_errors.RestErr) {
	userId, userErr := strconv.ParseInt(planIdParam, 10, 64)
	if userErr != nil {
		return 0, rest_errors.NewBadRequestError("user id should be a number")
	}
	return userId, nil
}

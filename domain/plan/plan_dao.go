package plan

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/DeVasu/tortoise/datasources/mysql/tortoise_db"
	rest_errors "github.com/DeVasu/tortoise/utils/errors"
)

const (
	queryInsertPlan          = 	"INSERT INTO plan(planName, amountOptions, tenureOptions, benefitPercentage, benefitType, createdAt, updatedAt) VALUES(?, ?, ?, ?, ?, ?, ?);"
	queryListPlans	   	     = 	"SELECT * from plan;"
	queryById 			     = 	"SELECT * from plan where planId=?;"
	queryAddPromotion 	     = 	"UPDATE plan SET promotionPercentage=?, promotionUsers=?, promotionStartDate=?, promotionEndDate=? WHERE planId = ?;"
	queryDeletePromotion     = 	"UPDATE plan SET promotionPercentage=NULL, promotionUsers=NULL, promotionStartDate=NULL, promotionEndDate=NULL WHERE planId = ?;"
	queryDecrementPromotion  = 	"UPDATE plan SET promotionUsers = promotionUsers - 1  WHERE planId = ? AND promotionUsers != 0"
)

func(plan *Plan) Create() *rest_errors.RestErr {

	//validation
	validationErr := plan.validate()
	if validationErr != nil {
		return validationErr
	}
	plan.CreatedAt = time.Now().Format("2006-01-02T15:04:05Z")
	plan.UpdatedAt = plan.CreatedAt

	// starting to insert
	stmt, err := tortoise_db.Client.Prepare(queryInsertPlan)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to get plans")
	}
	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(plan.PlanName, plan.AmountOptions, plan.TenureOptions, plan.BenefitPercentage, plan.BenefitType, plan.CreatedAt, plan.UpdatedAt)
	if saveErr != nil {
		return rest_errors.NewInternalServerError("error when tying to save plan")
	}

	//get Last insert Id
	planId, err := insertResult.LastInsertId()
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to save plan")
	}
	plan.PlanId = planId
	return nil
}

func (p *Plan) List() ([]Plan, *rest_errors.RestErr) {
	stmt, err := tortoise_db.Client.Prepare(queryListPlans)
	if err != nil {
		fmt.Println(err.Error())
		return nil, rest_errors.NewInternalServerError("error when tying to get cashier ( Database Error)")
	}
	defer stmt.Close()
	rows, err := stmt.Query() //update with limit and skip
	if err != nil {
		fmt.Println(err.Error())
		return nil, rest_errors.NewInternalServerError("error when tying to get cashier ( Database Error)")
	}
	defer rows.Close()

	results := make([]Plan, 0)
	for rows.Next() {
		var temp Plan
		var promo Promotion
		var newPercentage, usersLeft  sql.NullInt64
		var startDate, endDate sql.NullString
		if err := rows.Scan(
			&temp.PlanId,
			&temp.PlanName,
			&temp.AmountOptions,
			&temp.TenureOptions,
			&temp.BenefitPercentage,
			&temp.BenefitType,
			&newPercentage,
			&usersLeft,
			&startDate,
			&endDate,
			&temp.UpdatedAt,
			&temp.CreatedAt,
			); err != nil {
			fmt.Println(err.Error())
			return nil, rest_errors.NewInternalServerError("error when tying to get plans")
		}
		promo.EndDate = endDate.String
		promo.StartDate = startDate.String
		promo.NewPercentage = newPercentage.Int64
		promo.UsersLeft = usersLeft.Int64

		temp.Promotion = &promo
		temp.PromotionValid = temp.isPromotionValid()
		results = append(results, temp)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no cashiers matching status %s", "ok"))
	}
	return results, nil	
}

func(plan *Plan) GetById() *rest_errors.RestErr {
	stmt, err := tortoise_db.Client.Prepare(queryById)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to get cashier")
	}
	defer stmt.Close()

	result := stmt.QueryRow(plan.PlanId)

	var promo Promotion
	var newPercentage, usersLeft sql.NullInt64
	var startDate, endDate sql.NullString

	if err := result.Scan(
		&plan.PlanId,
		&plan.PlanName,
		&plan.AmountOptions,
		&plan.TenureOptions,
		&plan.BenefitPercentage,
		&plan.BenefitType,
		&newPercentage,
		&usersLeft,
		&startDate,
		&endDate,
		&plan.UpdatedAt,
		&plan.CreatedAt,
		); err != nil {
		return rest_errors.NewInternalServerError("error when tying to gett plan")
	}
	promo.EndDate = endDate.String
	promo.StartDate = startDate.String
	promo.NewPercentage = newPercentage.Int64
	promo.UsersLeft = usersLeft.Int64

	plan.Promotion = &promo
	plan.PromotionValid = plan.isPromotionValid()

	return nil
}

func(plan *Plan) AddPromotion(promo Promotion) *rest_errors.RestErr {



	plan.GetById()

	if plan.Promotion.NewPercentage != 0 {
		return rest_errors.NewBadRequestError("promtion already exists for this plan, please delete existing one")
	}

	valErr := promo.isValid()
	if valErr != nil {
		return valErr
	}

	stmt, err := tortoise_db.Client.Prepare(queryAddPromotion)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to get plan")
	}
	defer stmt.Close()

	plan.Promotion = &promo
	_, err = stmt.Exec(
		plan.Promotion.NewPercentage,
		plan.Promotion.UsersLeft,
		plan.Promotion.StartDate,
		plan.Promotion.EndDate,
		plan.PlanId,
	)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to update user ")
	}
	plan.PromotionValid = plan.isPromotionValid()
	return nil
}	

func(plan *Plan) DeletePromotion() *rest_errors.RestErr {

	stmt, err := tortoise_db.Client.Prepare(queryDeletePromotion)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to get plan")
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		plan.PlanId,
	)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to delete promotion ")
	}
	return nil
}	
func(plan *Plan) DecrementPromotion() *rest_errors.RestErr {

	stmt, err := tortoise_db.Client.Prepare(queryDecrementPromotion)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to get plan")
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		plan.PlanId,
	)
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to delete promotion")
	}
	return nil
}	
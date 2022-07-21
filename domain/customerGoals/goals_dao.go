package goals

import (
	// "database/sql"
	"fmt"
	"time"

	"github.com/DeVasu/tortoise/datasources/mysql/tortoise_db"
	"github.com/DeVasu/tortoise/domain/plan"
	rest_errors "github.com/DeVasu/tortoise/utils/errors"
)

const (
	queryInsertGoal        = 	"INSERT INTO customerGoals(planId, userId, selectedAmount, selectedTenure, startedDate, depositedAmount, benefitPercentage, benefitType, createdAt, updatedAt) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	queryListGoals	   	     = 	"SELECT * from customerGoals;"
)

func(goal *Goal) Create() *rest_errors.RestErr {

	//validation
	if validationErr := goal.validate(); validationErr != nil {
		return validationErr
	}
	if fillErr := goal.fillPlanInfo(); fillErr != nil {
		return fillErr
	}


	goal.StartedDate = time.Now().Format("2006-01-02")
	goal.CreatedAt = time.Now().Format("2006-01-02T15:04:05Z")
	goal.UpdatedAt = goal.CreatedAt

	// starting to insert
	stmt, err := tortoise_db.Client.Prepare(queryInsertGoal)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to get goals")
	}
	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(goal.PlanId, goal.UserId, goal.SelectedAmount, goal.SelectedTenure, goal.StartedDate, goal.DepositedAmount, goal.BenefitPercentage, goal.BenefitType, goal.UpdatedAt, goal.CreatedAt)
	if saveErr != nil {
		return rest_errors.NewInternalServerError("error when tying to save goal")
	}

	//get Last insert Id
	goalId, err := insertResult.LastInsertId()
	if err != nil {
		return rest_errors.NewInternalServerError("error when tying to save goal")
	}

	tempPlan := &plan.Plan{
		PlanId: goal.PlanId,
	}
	decrementErr := tempPlan.DecrementPromotion(); 
	if decrementErr != nil {
		return decrementErr
	}

	goal.GoalId = goalId
	return nil
}

func (goal *Goal) List() ([]Goal, *rest_errors.RestErr) {
	stmt, err := tortoise_db.Client.Prepare(queryListGoals)
	if err != nil {
		fmt.Println(err.Error())
		return nil, rest_errors.NewInternalServerError("error when tying to get goals ( Database Error)")
	}
	defer stmt.Close()
	rows, err := stmt.Query() 
	if err != nil {
		fmt.Println(err.Error())
		return nil, rest_errors.NewInternalServerError("error when tying to get goals ( Database Error)")
	}
	defer rows.Close()

	results := make([]Goal, 0)
	for rows.Next() {
		var temp Goal
		// var promo Promotion
		// var newPercentage, usersLeft  sql.NullInt64
		// var startDate, endDate sql.NullString
		if err := rows.Scan(
			&temp.GoalId,
			&temp.PlanId,
			&temp.UserId,
			&temp.SelectedAmount,
			&temp.SelectedTenure,
			&temp.StartedDate,
			&temp.DepositedAmount,
			&temp.BenefitPercentage,
			&temp.BenefitType,
			&temp.UpdatedAt,
			&temp.CreatedAt,
			); err != nil {
			fmt.Println(err.Error())
			return nil, rest_errors.NewInternalServerError("error when tying to get Goals")
		}
		// promo.EndDate = endDate.String
		// promo.StartDate = startDate.String
		// promo.NewPercentage = newPercentage.Int64
		// promo.UsersLeft = usersLeft.Int64

		// temp.Promotion = &promo
		// temp.PromotionValid = temp.isPromotionValid()
		results = append(results, temp)
	}
	if len(results) == 0 {
		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no cashiers matching status %s", "ok"))
	}
	return results, nil	
}
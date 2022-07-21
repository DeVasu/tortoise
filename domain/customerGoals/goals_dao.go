package goals

import (
	"time"

	"github.com/DeVasu/tortoise/datasources/mysql/cashiers_db"
	rest_errors "github.com/DeVasu/tortoise/utils/errors"
)

const (
	queryInsertGoal        = 	"INSERT INTO customerGoals(planId, userId, selectedAmount, selectedTenure, startedDate, depositedAmount, benefitPercentage, benefitType, createdAt, updatedAt) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"
	queryListProducts 	   = 	"SELECT * from plan;"
	queryById 			   = 	"SELECT * from plan where planId=?;"
	queryAddPromotion 	   = 	"UPDATE plan SET promotionPercentage=?, promotionUsers=?, promotionStartDate=?, promotionEndDate=? WHERE planId = ?;"
	queryDeletePromotion   = 	"UPDATE plan SET promotionPercentage=NULL, promotionUsers=NULL, promotionStartDate=NULL, promotionEndDate=NULL WHERE planId = ?;"
	queryDeleteProduct = "DELETE FROM products WHERE id=?;"
	queryUpdateProduct = "UPDATE products SET categoryId=?, name=?, image=?, price=?, stock=? WHERE id = ?;"
	
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
	stmt, err := cashiers_db.Client.Prepare(queryInsertGoal)
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
	goal.GoalId = goalId
	return nil
}

// func (p *Plan) List() ([]Plan, *rest_errors.RestErr) {
// 	stmt, err := cashiers_db.Client.Prepare(queryListProducts)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, rest_errors.NewInternalServerError("error when tying to get cashier ( Database Error)")
// 	}
// 	defer stmt.Close()
// 	rows, err := stmt.Query() //update with limit and skip
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, rest_errors.NewInternalServerError("error when tying to get cashier ( Database Error)")
// 	}
// 	defer rows.Close()

// 	results := make([]Plan, 0)
// 	for rows.Next() {
// 		var temp Plan
// 		var promo Promotion
// 		var newPercentage, usersLeft  sql.NullInt64
// 		var startDate, endDate sql.NullString
// 		if err := rows.Scan(
// 			&temp.PlanId,
// 			&temp.PlanName,
// 			&temp.AmountOptions,
// 			&temp.TenureOptions,
// 			&temp.BenefitPercentage,
// 			&temp.BenefitType,
// 			&newPercentage,
// 			&usersLeft,
// 			&startDate,
// 			&endDate,
// 			&temp.UpdatedAt,
// 			&temp.CreatedAt,
// 			); err != nil {
// 			fmt.Println(err.Error())
// 			return nil, rest_errors.NewInternalServerError("error when tying to get plans")
// 		}
// 		promo.EndDate = endDate.String
// 		promo.StartDate = startDate.String
// 		promo.NewPercentage = newPercentage.Int64
// 		promo.UsersLeft = usersLeft.Int64

// 		temp.Promotion = &promo
// 		temp.PromotionValid = temp.isPromotionValid()
// 		results = append(results, temp)
// 	}
// 	if len(results) == 0 {
// 		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no cashiers matching status %s", "ok"))
// 	}
// 	return results, nil	
// }

// func(plan *Plan) GetById() *rest_errors.RestErr {
// 	stmt, err := cashiers_db.Client.Prepare(queryById)
// 	if err != nil {
// 		return rest_errors.NewInternalServerError("error when tying to get cashier")
// 	}
// 	defer stmt.Close()

// 	result := stmt.QueryRow(plan.PlanId)

// 	var promo Promotion
// 	var newPercentage, usersLeft sql.NullInt64
// 	var startDate, endDate sql.NullString

// 	if err := result.Scan(
// 		&plan.PlanId,
// 		&plan.PlanName,
// 		&plan.AmountOptions,
// 		&plan.TenureOptions,
// 		&plan.BenefitPercentage,
// 		&plan.BenefitType,
// 		&newPercentage,
// 		&usersLeft,
// 		&startDate,
// 		&endDate,
// 		&plan.UpdatedAt,
// 		&plan.CreatedAt,
// 		); err != nil {
// 		return rest_errors.NewInternalServerError("error when tying to gett plan")
// 	}
// 	promo.EndDate = endDate.String
// 	promo.StartDate = startDate.String
// 	promo.NewPercentage = newPercentage.Int64
// 	promo.UsersLeft = usersLeft.Int64

// 	plan.Promotion = &promo

// 	return nil
// }

// func(plan *Plan) AddPromotion(promo Promotion) *rest_errors.RestErr {

// 	valErr := promo.isValid()
// 	if valErr != nil {
// 		return valErr
// 	}

// 	plan.GetById()

// 	if plan.Promotion.NewPercentage != 0 {
// 		return rest_errors.NewBadRequestError("promtion already exists for this plan, please delete existing one")
// 	}

// 	stmt, err := cashiers_db.Client.Prepare(queryAddPromotion)
// 	if err != nil {
// 		return rest_errors.NewInternalServerError("error when tying to get plan")
// 	}
// 	defer stmt.Close()

// 	plan.Promotion = &promo
// 	_, err = stmt.Exec(
// 		plan.Promotion.NewPercentage,
// 		plan.Promotion.UsersLeft,
// 		plan.Promotion.StartDate,
// 		plan.Promotion.EndDate,
// 		plan.PlanId,
// 	)
// 	if err != nil {
// 		return rest_errors.NewInternalServerError("error when tying to update user ")
// 	}
// 	return nil
// }	

// func(plan *Plan) DeletePromotion() *rest_errors.RestErr {

// 	stmt, err := cashiers_db.Client.Prepare(queryDeletePromotion)
// 	if err != nil {
// 		return rest_errors.NewInternalServerError("error when tying to get plan")
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(
// 		plan.PlanId,
// 	)
// 	if err != nil {
// 		return rest_errors.NewInternalServerError("error when tying to delete promotion ")
// 	}
// 	return nil
// }	


	// if p.CategoryId != 0 {
	// 	temp.CategoryId = p.CategoryId
	// }
	// if len(p.Name) != 0 {
	// 	temp.Name = p.Name
	// }
	// if len(p.Image) != 0 {
	// 	temp.Image = p.Image
	// }
	// if p.Price != 0 {
	// 	temp.Price = p.Price
	// }
	// if p.Stock != 0 {
	// 	temp.Stock = p.Stock
	// }

// func(p *Product) Delete() *rest_errors.RestErr {
// 	stmt, err := cashiers_db.Client.Prepare(queryDeleteProduct)
// 	if err != nil {
// 		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(p.Id)
// 	if err != nil {
// 		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
// 	}
// 	return nil
// }

// func(p *Product) Update() *rest_errors.RestErr {

// 	temp := &Product{
// 		Id : p.Id,
// 	}
// 	temp.GetById()

// 	if p.CategoryId != 0 {
// 		temp.CategoryId = p.CategoryId
// 	}
// 	if len(p.Name) != 0 {
// 		temp.Name = p.Name
// 	}
// 	if len(p.Image) != 0 {
// 		temp.Image = p.Image
// 	}
// 	if p.Price != 0 {
// 		temp.Price = p.Price
// 	}
// 	if p.Stock != 0 {
// 		temp.Stock = p.Stock
// 	}


// 	stmt, err := cashiers_db.Client.Prepare(queryUpdateProduct)
// 	if err != nil {
// 		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec(
// 		temp.CategoryId,
// 		temp.Name,
// 		temp.Image,
// 		temp.Price,
// 		temp.Stock,
// 		temp.Id,
// 	)
// 	if err != nil {
// 		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
// 	}
// 	return nil
// }

// func(p *Product) GetById() *rest_errors.RestErr {
// 	stmt, err := cashiers_db.Client.Prepare(queryById)
// 	if err != nil {
// 		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
// 	}
// 	defer stmt.Close()

// 	result := stmt.QueryRow(p.Id)

// 	if err := result.Scan(&p.Id,
// 		&p.CategoryId,
// 		&p.Name,
// 		&p.Image,
// 		&p.Price,
// 		&p.Stock,
// 		&p.UpdatedAt,
// 		&p.CreatedAt,
// 		&p.Discount.Qty, 
// 		&p.Discount.Type,
// 		&p.Discount.Result,
// 		&p.Discount.ExpiredAt,
// 		); err != nil {
// 		return rest_errors.NewInternalServerError("error when tying to gett cashier", errors.New("database error"))
// 	}

// 	return nil
// }




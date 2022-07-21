package goals

import (
	"github.com/DeVasu/tortoise/domain/plan"
	rest_errors "github.com/DeVasu/tortoise/utils/errors"
)

type Goal struct {
	GoalId				int64		`json:"goalId,omitempty"`
	PlanId 				int64  		`json:"planId,omitempty"`
	UserId 				int64 		`json:"userId,omitempty"`
	SelectedAmount 		int64 		`json:"selectedAmoung"`
	SelectedTenure 		int64 		`json:"selectedTenure"`
	StartedDate 		string 		`json:"startedDate,omitempty"`
	DepositedAmount 	int64 		`json:"depositedAmount,omitempty"`
	BenefitPercentage 	int64 		`json:"benefitPercentage"`
	BenefitType 		string 		`json:"benefitType"`  //cashback/ Extravoucher
	UpdatedAt 			string 		`json:"updatedAt,omitempty"`
	CreatedAt 			string 		`json:"createdAt,omitempty"`
}

func(goal *Goal) validate() *rest_errors.RestErr {

	if goal.UserId == 0 {
		return rest_errors.NewBadRequestError("userId cannot be 0 or empty")
	}
	if goal.PlanId == 0 {
		return rest_errors.NewBadRequestError("planId  cannot be 0 or empty")
	}
	return nil
}

func(goal *Goal) fillPlanInfo() *rest_errors.RestErr {

	selectedPlan := &plan.Plan{
		PlanId: goal.PlanId,
	}

	if getPlanErr := selectedPlan.GetById(); getPlanErr != nil {
		return getPlanErr
	}

	goal.SelectedAmount = selectedPlan.AmountOptions
	goal.SelectedTenure = selectedPlan.TenureOptions
	goal.DepositedAmount = selectedPlan.AmountOptions
	if selectedPlan.PromotionValid {
		goal.BenefitPercentage = selectedPlan.Promotion.NewPercentage
	} else {
		goal.BenefitPercentage = selectedPlan.BenefitPercentage
	}
	goal.BenefitType = selectedPlan.BenefitType

	return nil
}

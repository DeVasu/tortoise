package plan

import (
	"time"

	rest_errors "github.com/DeVasu/tortoise/utils/errors"
)

type Plan struct {
	PlanId 				int64  	`json:"planId,omitempty"`
	PlanName 			string 	`json:"planName,omitempty"`
	AmountOptions 		int64 	`json:"amountOptions"`
	TenureOptions 		int64 	`json:"tenureOptions"`
	BenefitPercentage 	int64 	`json:"benefitPercentage"`
	BenefitType 		string 	`json:"benefitType"`  //cashback/ Extravoucher
	Promotion			*Promotion `json:"promotion"`
	UpdatedAt 			string 	`json:"updatedAt,omitempty"`
	CreatedAt 			string 	`json:"createdAt,omitempty"`
	PromotionValid  	bool	`json:"promotionValid"`
}

type Promotion struct {
	NewPercentage	int64	`json:"newPercentage"`
	UsersLeft		int64	`json:"usersLeft"`
	StartDate		string  `json:"startDate"`
	EndDate			string  `json:"endDate"`
}

func(promo Promotion) isValid() *rest_errors.RestErr {

	if promo.NewPercentage == 0 {
		return rest_errors.NewBadRequestError("promotionPercentage can't be empty or 0")
	}
	if promo.UsersLeft == 0 && promo.StartDate == "" {
		return rest_errors.NewBadRequestError("either usersLeft or startDate/endDate has to be set")
	}
	if promo.UsersLeft==0 {
		if(promo.StartDate > promo.EndDate) {
			return rest_errors.NewBadRequestError("startDate is bigger than endDate")
		}
	}
	return nil

}

func(p Plan) isPromotionValid() bool {
	if p.Promotion == nil {
		return false
	}
	if p.Promotion.StartDate == "" {
		if p.Promotion.UsersLeft > 0 {
			return true
		} else {
			return false
		}
	}
	currentTime := time.Now().Format("2006-01-02")
	
	if currentTime >= p.Promotion.StartDate && currentTime <= p.Promotion.EndDate {
		return true
	}
	return false
} 

func(p Plan) validate() *rest_errors.RestErr {
	if p.PlanName == "" {
		return rest_errors.NewBadRequestError("Name cannot be empty")
	}
	return nil
}